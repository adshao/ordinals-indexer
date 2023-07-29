package ord

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/ord/page"
	"github.com/adshao/ordinals-indexer/internal/ord/parser"

	"github.com/adshao/go-brc721/sig"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var lastInscriptionIdFile = int64(0)

var ProviderSet = wire.NewSet(NewSyncer)

type result struct {
	info *page.Inscription
	err  error
}

type uids []string

type Syncer struct {
	c                     *conf.Ord
	data                  *data.Data
	collectionUc          *biz.CollectionUsecase
	tokenUc               *biz.TokenUsecase
	pageParser            page.PageParser
	logger                *log.Helper
	inscriptionUidChan    chan string
	resultChan            chan *result
	processChan           chan uids
	processFinishedChan   chan error
	eventChan             chan Event
	stopC                 chan struct{}
	lastInscriptionIdChan chan int64
}

func NewSyncer(c *conf.Ord, data *data.Data, collectionUc *biz.CollectionUsecase, tokenUc *biz.TokenUsecase, logger log.Logger) (*Syncer, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the syncer resources")
	}
	syncer := &Syncer{
		c:            c,
		data:         data,
		collectionUc: collectionUc,
		tokenUc:      tokenUc,
		pageParser:   page.NewPageParser(c),
		logger:       log.NewHelper(logger),
	}
	concurrency := c.Worker.Concurrency
	syncer.inscriptionUidChan = make(chan string, concurrency)
	syncer.resultChan = make(chan *result, concurrency)
	syncer.processChan = make(chan uids)
	syncer.processFinishedChan = make(chan error)
	syncer.eventChan = make(chan Event, concurrency)
	syncer.stopC = make(chan struct{})
	syncer.lastInscriptionIdChan = make(chan int64)
	return syncer, cleanup, nil
}

func (s *Syncer) Run() error {
	// TODO: we need to detect reorg and delete invalid data before we upsert new data
	concurrency := s.c.Worker.Concurrency
	workers := make([]*Worker, concurrency)
	wg := &sync.WaitGroup{}
	wg.Add(int(concurrency))
	for i := 0; i < int(concurrency); i++ {
		workers[i] = &Worker{
			wid:        i,
			baseURL:    s.c.Server.Addr,
			pageParser: page.NewPageParser(s.c),
			data:       s.data,
			uidChan:    s.inscriptionUidChan,
			resultChan: s.resultChan,
			stopC:      s.stopC,
			logger:     s.logger,
		}
		go func(worker *Worker) {
			defer wg.Done()
			worker.Start()
		}(workers[i])
	}
	go func() {
		s.receveResult()
	}()

	go func() {
		for {
			select {
			case <-s.stopC:
				s.logger.Infof("stopping inscriptions processor")
				return
			default:
				lastInscriptionId, _ := s.getLastInscriptionId()
				s.lastInscriptionIdChan <- lastInscriptionId
				err := s.parseInscriptions(lastInscriptionId)
				if err != nil {
					s.logger.Errorf("failed to parse inscriptions: %v", err)
				}
				time.Sleep(60 * time.Second)
			}
		}
	}()

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt)
	signal := <-terminateSignals
	s.logger.Infof("received signal %s, stopping workers...", signal)
	close(s.inscriptionUidChan)
	close(s.stopC)
	wg.Wait()
	s.logger.Info("all workers have been stopped")
	return nil
}

func (s *Syncer) receveResult() {
	insUids := make(uids, 0)
	results := make(map[string]*result)
	resultCount := 0
	var lastInscriptionId int64
	var stopped bool
	for {
		select {
		case lastInscriptionId = <-s.lastInscriptionIdChan:
			s.logger.Debugf("received lastInscriptionId: %d", lastInscriptionId)
		case result := <-s.resultChan:
			results[result.info.UID] = result
			s.logger.Debugf("received result for inscription %d", result.info.ID)
			resultCount++
		case insUids = <-s.processChan:
			s.logger.Debugf("receiving %d inscriptions", len(insUids))
		case <-s.stopC:
			stopped = true
		default:
			if resultCount == len(insUids) && len(insUids) > 0 {
				if stopped {
					s.logger.Infof("result processor stopped")
					return
				}
				var err error
				processedResultCount := 0
				resultCount = 0
				resultsInOrder := make([]*result, len(insUids))
				for i := 0; i < len(insUids); i++ {
					s.logger.Debugf("insUids[%d]: %s", i, insUids[i])
					resultsInOrder[i] = results[insUids[len(insUids)-i-1]]
				}
				if len(resultsInOrder) != len(insUids) {
					err = fmt.Errorf("resultsInOrder length %d != insUids length %d", len(resultsInOrder), len(insUids))
				} else {
					// make sure to process results in ascending order of inscriptionId
					lastId := int64(0)
					for _, result := range resultsInOrder {
						if result.info.ID < lastId {
							err = fmt.Errorf("results are not in order, lastId: %d, currentId: %d", lastId, result.info.ID)
							break
						}
						lastId = result.info.ID
					}
					if err == nil {
						processedResultCount, err = s.processResults(resultsInOrder, lastInscriptionId)
					}
				}
				if err != nil {
					s.logger.Errorf("failed to process results: %v", err)
					s.processFinishedChan <- err
				} else {
					s.logger.Infof("processed %d results", processedResultCount)
					s.processFinishedChan <- nil
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Syncer) saveLastInscriptionIdToFile(lastInscriptionId int64) error {
	if lastInscriptionIdFile >= lastInscriptionId {
		return nil
	}
	return ioutil.WriteFile(".last_inscription_id", []byte(strconv.FormatInt(lastInscriptionId, 10)), 0644)
}

func (s *Syncer) processResults(resultsInOrder []*result, lastInscriptionId int64) (int, error) {
	count := 0
	var lastSuccessInscriptionId int64
	for _, result := range resultsInOrder {
		if result.info.ID < lastInscriptionId {
			s.logger.Debugf("inscription %d is less than lastInscriptionId %d, ignore", result.info.ID, lastInscriptionId)
			continue
		}
		if result.err != nil {
			return count, result.err
		}
		err := s.processResult(result)
		if err != nil {
			return count, err
		}
		s.logger.Infof("processed inscription %d", result.info.ID)
		lastSuccessInscriptionId = result.info.ID
		count++
	}
	if lastSuccessInscriptionId > 0 && lastSuccessInscriptionId > lastInscriptionId {
		s.saveLastInscriptionIdToFile(lastSuccessInscriptionId)
	}
	return count, nil
}

func (s *Syncer) processResult(result *result) error {
	info := result.info
	if info.Content == nil {
		return fmt.Errorf("content of inscription %d is nil", info.ID)
	}
	switch info.Content.Type {
	case parser.NameBRC721Deploy:
		err := s.processBRC721Deploy(info)
		if err != nil {
			return err
		}
	case parser.NameBRC721Mint:
		err := s.processBRC721Mint(info)
		if err != nil {
			return err
		}
	case parser.NameBRC721Update:
		// TODO: enable this after we have a way to identify the owner of the collection
		// err := s.processBRC721Update(inscriptionId, info)
		// if err != nil {
		//     return err
		// }
	default:
	}
	return nil
}

func (s *Syncer) processBRC721Deploy(info *page.Inscription) error {
	o := info.Content.Data.(*parser.BRC721Deploy)
	// check if the collection already exists
	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
	if err != nil {
		return err
	}
	if collection != nil {
		if collection.InscriptionID > info.ID {
			// TODO: need to check if the collection is valid
			s.logger.Warnf("collection %s already exists, but inscriptionId %d is greater than %d, ignore inscription %d", o.Tick, collection.InscriptionID, info.ID, info.ID)
		} else {
			s.logger.Infof("collection %s already exists, ignore inscription %d", o.Tick, info.ID)
		}
		return nil
	}
	// check deploy sig
	if o.Sig != nil {
		err = o.Sig.Validate()
		if err != nil {
			s.logger.Warnf("invalid deploy sig for collection %s, ignore inscription %d", o.Tick, info.ID)
			return nil
		}
	}
	// create the collection
	collection = &biz.Collection{
		P:      biz.ProtocolTypeBRC721,
		Tick:   o.Tick,
		Supply: 0,
	}
	max, err := strconv.ParseUint(o.Max, 10, 64)
	if err != nil {
		s.logger.Warnf("invalid max %s, ignore inscription %d", o.Max, info.ID)
		return nil
	}
	collection.Max = max
	if o.BaseURI != nil {
		collection.BaseURI = *o.BaseURI
	}
	if o.Meta != nil {
		collection.Name = o.Meta.Name
		collection.Description = o.Meta.Description
		collection.Image = o.Meta.Image
		collection.Attributes = o.Meta.Attributes
	}
	collection.TxHash = info.GenesisTx
	collection.BlockHeight = info.GenesisHeight
	collection.BlockTime = info.Timestamp
	collection.Address = info.Address
	collection.InscriptionID = info.ID
	collection.InscriptionUID = info.UID
	if o.Sig != nil {
		collection.Sig = *o.Sig
	}
	collection, err = s.collectionUc.CreateCollection(context.Background(), collection)
	if err != nil {
		return err
	}
	s.logger.Infof("created collection %s for inscription %d", o.Tick, info.ID)
	return nil
}

func (s *Syncer) processBRC721Mint(info *page.Inscription) error {
	inscriptionId := info.ID
	o := info.Content.Data.(*parser.BRC721Mint)
	// check if the collection exists
	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
	if err != nil {
		return err
	}
	if collection == nil {
		s.logger.Infof("collection %s not found, ignore mint inscription %d", o.Tick, inscriptionId)
		return nil
	}
	if collection.InscriptionID >= inscriptionId {
		s.logger.Warnf("collection %s inscriptionId %d is greater than %d, ignore mint inscription %d", o.Tick, collection.InscriptionID, inscriptionId, inscriptionId)
		return nil
	}
	s.logger.Debugf("collection: %+v", collection)
	// check if supply is full
	if collection.Supply >= collection.Max {
		s.logger.Infof("collection %s supply is full, ignore mint inscription %d", o.Tick, inscriptionId)
		return nil
	}
	t, err := s.tokenUc.FindByInscriptionID(context.Background(), inscriptionId)
	if err != nil {
		return err
	}
	if len(t) > 0 {
		s.logger.Infof("token with inscription %d already processed, ignore mint inscription", inscriptionId)
		return nil
	}
	valid, mintSig, err := s.checkBRC721MintSig(info, collection, o)
	if err != nil {
		return err
	}
	if !valid {
		return nil
	}
	// create token
	token := &biz.Token{
		Tick:           o.Tick,
		P:              biz.ProtocolTypeBRC721,
		TokenID:        collection.Supply + 1,
		TxHash:         info.GenesisTx,
		BlockHeight:    info.GenesisHeight,
		BlockTime:      info.Timestamp,
		Address:        info.Address,
		InscriptionID:  inscriptionId,
		InscriptionUID: info.UID,
		CollectionID:   collection.ID,
	}
	if mintSig != nil {
		token.Sig = *mintSig
	}
	token, err = s.tokenUc.CreateToken(context.Background(), token)
	if err != nil {
		s.logger.Errorf("failed to create token: %T: %v", err, err)
		return err
	}
	s.logger.Infof("created token %d for inscription %d", token.TokenID, inscriptionId)

	collection.Supply++
	collection, err = s.collectionUc.UpdateCollection(context.Background(), collection)
	if err != nil {
		return err
	}
	s.logger.Infof("updated collection %s supply to %d", o.Tick, collection.Supply)
	return nil
}

func (s *Syncer) checkBRC721MintSig(info *page.Inscription, collection *biz.Collection, o *parser.BRC721Mint) (bool, *sig.MintSig, error) {
	inscriptionId := info.ID
	var mintSig *sig.MintSig
	// check mint sig
	if collection.Sig.PubKey == "" || len(collection.Sig.Fields) == 0 {
		return true, nil, nil
	}
	// verify mint sig
	if o.Sig == nil {
		s.logger.Warnf("missing mint sig for collection %s, ignore mint inscription %d", o.Tick, inscriptionId)
		return false, nil, nil
	}
	if o.Sig.Signature == "" {
		s.logger.Warnf("missing mint sig.s for collection %s, ignore mint inscription %d", o.Tick, inscriptionId)
		return false, nil, nil
	}
	mintSig = &sig.MintSig{
		Signature: o.Sig.Signature,
	}
	for _, field := range collection.Sig.Fields {
		switch field {
		case sig.SigFieldReceiver:
			mintSig.Receiver = info.Address
		case sig.SigFieldUid:
			if o.Sig.Uid == "" {
				s.logger.Warnf("missing mint sig.uid for collection %s, ignore mint inscription %d", o.Tick, inscriptionId)
				return false, nil, nil
			}
			utoken, err := s.tokenUc.FindByTickSigUID(context.Background(), biz.ProtocolTypeBRC721, o.Tick, o.Sig.Uid)
			if err != nil {
				s.logger.Errorf("failed to find token by tick %s and sig.uid %s: %v", o.Tick, o.Sig.Uid, err)
				return false, nil, err
			}
			if utoken != nil {
				s.logger.Warnf("mint sig.uid %s already exists for collection %s, ignore mint inscription %d", o.Sig.Uid, o.Tick, inscriptionId)
				return false, nil, nil
			}
			mintSig.Uid = o.Sig.Uid
		case sig.SigFieldExpiredTime:
			if o.Sig.ExpiredTime == 0 {
				s.logger.Warnf("missing mint sig.expt for collection %s, ignore mint inscription %d", o.Tick, inscriptionId)
				return false, nil, nil
			}
			// parse '2023-07-27 12:51:49 UTC' to unix timestamp in seconds
			if info.Timestamp.IsZero() {
				s.logger.Errorf("Impossible! Invalid timestamp %s for collection %s, ignore mint inscription %d", info.Timestamp, o.Tick, inscriptionId)
				return false, nil, fmt.Errorf("invalid timestamp %s for inscription %d", info.Timestamp, inscriptionId)
			}
			if uint64(info.Timestamp.Unix()) > o.Sig.ExpiredTime {
				s.logger.Warnf("mint sig.expt %d is expired for collection %s, ignore mint inscription %d", o.Sig.ExpiredTime, o.Tick, inscriptionId)
				return false, nil, nil
			}
			mintSig.ExpiredTime = o.Sig.ExpiredTime
		case sig.SigFieldExpiredHeight:
			if o.Sig.ExpiredHeight == 0 {
				s.logger.Warnf("missing mint sig.exph for collection %s, ignore mint inscription %d", o.Tick, inscriptionId)
				return false, nil, nil
			}
			if info.GenesisHeight > o.Sig.ExpiredHeight {
				s.logger.Warnf("mint sig.exph %d is expired for collection %s, ignore mint inscription %d", o.Sig.ExpiredHeight, o.Tick, inscriptionId)
				return false, nil, nil
			}
			mintSig.ExpiredHeight = o.Sig.ExpiredHeight
		}
	}
	pubKey, err := sig.ParsePubKey(collection.Sig.PubKey)
	if err != nil {
		s.logger.Errorf("Impossible! Invalid public key %s for collection %s, ignore mint inscription %d", collection.Sig.PubKey, o.Tick, inscriptionId)
		return false, nil, err
	}
	valid, err := mintSig.Verify(pubKey)
	if err != nil {
		s.logger.Warnf("failed to verify mint sig for collection %s, ignore mint inscription %d: %v", o.Tick, inscriptionId, err)
		return false, nil, nil
	}
	if !valid {
		s.logger.Warnf("invalid mint sig for collection %s, ignore mint inscription %d", o.Tick, inscriptionId)
		return false, nil, nil
	}
	return true, mintSig, nil
}

func (s *Syncer) processBRC721Update(info *page.Inscription) error {
	inscriptionId := info.ID
	o := info.Content.Data.(*parser.BRC721Update)
	// check if the collection exists
	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
	if err != nil {
		return err
	}
	if collection == nil {
		s.logger.Infof("collection %s not found, ignore inscription %d", o.Tick, inscriptionId)
		return nil
	}
	// update collection
	if o.BaseURI != nil {
		collection.BaseURI = *o.BaseURI
	}
	_, err = s.collectionUc.UpdateCollection(context.Background(), collection)
	if err != nil {
		return err
	}
	s.logger.Infof("updated collection %s", o.Tick)
	return nil
}

func (s *Syncer) getLastInscriptionId() (int64, error) {
	lastInscriptionId, err := s.readLastInscriptionIdFromFile()
	if err == nil {
		s.logger.Infof("get lastInscriptionId from file: %d", lastInscriptionId)
		return lastInscriptionId, nil
	}

	// lastInscriptionId, err := s.data.Rdb.GetUint64(ctx, lastInscriptionIdKey)
	// if err != nil {
	//     s.logger.Warnf("failed to get lastInscriptionId from redis: %v", err)
	//     return 0, err
	// }
	// if lastInscriptionId != 0 {
	//     s.logger.Infof("get lastInscriptionId from redis: %d", lastInscriptionId)
	//     return lastInscriptionId, nil
	// }
	lastInscriptionId = s.c.Server.InscriptionIdStart
	s.logger.Infof("get lastInscriptionId from config: %d", lastInscriptionId)
	return lastInscriptionId, nil
}

func (s *Syncer) readLastInscriptionIdFromFile() (int64, error) {
	b, err := ioutil.ReadFile(".last_inscription_id")
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return 0, err
	}
	lastInscriptionIdFile = id
	return id, nil
}

func (s *Syncer) parseInscriptions(inscriptionId int64) error {
	inscriptionsPage := page.NewInscriptionsPage(inscriptionId)
	s.logger.Infof("parsing inscriptions page %s", inscriptionsPage.URL())
	data, err := s.pageParser.Parse(inscriptionsPage)
	if err != nil {
		return err
	}
	inscriptions, ok := data.(*page.Inscriptions)
	if !ok {
		return fmt.Errorf("invalid data type: %T for URL %s", data, inscriptionsPage.URL())
	}

	insUids := inscriptions.UIDs
	s.processChan <- insUids
	for _, insUid := range insUids {
		s.inscriptionUidChan <- insUid
	}
	// wait for the process to finish
	err = <-s.processFinishedChan
	if err != nil {
		return err
	}

	// check if there is a next page
	if inscriptions.NextID != nil {
		return s.parseInscriptions(*inscriptions.NextID)
	}
	return nil
}
