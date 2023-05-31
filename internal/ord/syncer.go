package ord

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
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
	"github.com/adshao/ordinals-indexer/internal/ord/parser"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var lastInscriptionIdFile = int64(0)

var ProviderSet = wire.NewSet(NewSyncer)

type result struct {
	inscriptionUid string
	inscriptionId  int64
	info           map[string]interface{}
	err            error
}

type uids []string

type Syncer struct {
	c                     *conf.Ord
	data                  *data.Data
	collectionUc          *biz.CollectionUsecase
	tokenUc               *biz.TokenUsecase
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
			lastInscriptionId, _ := s.getLastInscriptionId()
			s.lastInscriptionIdChan <- lastInscriptionId
			inscriptionURL, _ := url.JoinPath(s.c.Server.Addr, "inscriptions", fmt.Sprintf("%d", lastInscriptionId))
			s.logger.Infof("start crawling from %s", inscriptionURL)
			_, err := s.parseInscriptions(inscriptionURL)
			if err != nil {
				s.logger.Errorf("failed to parse inscriptions: %v", err)
			}
			time.Sleep(60 * time.Second)
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
	for {
		select {
		case lastInscriptionId = <-s.lastInscriptionIdChan:
			s.logger.Debugf("received lastInscriptionId: %d", lastInscriptionId)
		case result := <-s.resultChan:
			results[result.inscriptionUid] = result
			s.logger.Debugf("received result for inscription %d", result.inscriptionId)
			resultCount++
		case insUids = <-s.processChan:
			s.logger.Debugf("receiving %d inscriptions", len(insUids))
		case <-s.stopC:
			s.logger.Infof("stopping result processor")
			return
		default:
			if resultCount == len(insUids) && len(insUids) > 0 {
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
						if result.inscriptionId < lastId {
							err = fmt.Errorf("results are not in order, lastId: %d, currentId: %d", lastId, result.inscriptionId)
							break
						}
						lastId = result.inscriptionId
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
		if result.inscriptionId < lastInscriptionId {
			s.logger.Debugf("inscription %d is less than lastInscriptionId %d, ignore", result.inscriptionId, lastInscriptionId)
			continue
		}
		if result.err != nil {
			return count, result.err
		}
		err := s.processResult(result)
		if err != nil {
			return count, err
		}
		s.logger.Infof("processed inscription %d", result.inscriptionId)
		lastSuccessInscriptionId = result.inscriptionId
		count++
	}
	if lastSuccessInscriptionId > 0 && lastSuccessInscriptionId > lastInscriptionId {
		s.saveLastInscriptionIdToFile(lastSuccessInscriptionId)
	}
	return count, nil
}

func (s *Syncer) processResult(result *result) error {
	inscriptionId := result.inscriptionId
	info := result.info
	switch info["content_parser"].(string) {
	case parser.NameBRC721Deploy:
		err := s.processBRC721Deploy(inscriptionId, info)
		if err != nil {
			return err
		}
	case parser.NameBRC721Mint:
		err := s.processBRC721Mint(inscriptionId, info)
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

func (s *Syncer) processBRC721Deploy(inscriptionId int64, info map[string]interface{}) error {
	o := info["content"].(*parser.BRC721Deploy)
	// check if the collection already exists
	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
	if err != nil {
		return err
	}
	if collection != nil {
		if collection.InscriptionID > inscriptionId {
			s.logger.Warnf("collection %s already exists, but inscriptionId %d is less than %d, ignore inscription %d", o.Tick, collection.InscriptionID, inscriptionId, inscriptionId)
		} else {
			s.logger.Infof("collection %s already exists, ignore inscription %d", o.Tick, inscriptionId)
		}
		return nil
	}
	// create the collection
	collection = &biz.Collection{
		P:      biz.ProtocolTypeBRC721,
		Tick:   o.Tick,
		Supply: 0,
	}
	max, err := strconv.ParseUint(o.Max, 10, 64)
	if err != nil {
		return err
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
	collection.TxHash = info["genesis_transaction"].(string)
	collection.BlockHeight = info["genesis_height"].(uint64)
	collection.BlockTime = info["timestamp"].(time.Time)
	collection.Address = info["address"].(string)
	collection.InscriptionID = inscriptionId
	collection.InscriptionUID = info["uid"].(string)
	collection, err = s.collectionUc.CreateCollection(context.Background(), collection)
	if err != nil {
		return err
	}
	s.logger.Infof("created collection %s for inscription %d", o.Tick, inscriptionId)
	return nil
}

func (s *Syncer) processBRC721Mint(inscriptionId int64, info map[string]interface{}) error {
	o := info["content"].(*parser.BRC721Mint)
	// check if the collection exists
	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
	if err != nil {
		return err
	}
	if collection == nil {
		s.logger.Infof("collection %s not found, ignore inscription %d", o.Tick, inscriptionId)
		return nil
	}
	if collection.InscriptionID >= inscriptionId {
		s.logger.Warnf("collection %s inscriptionId %d is greater than %d, ignore inscription %d", o.Tick, collection.InscriptionID, inscriptionId, inscriptionId)
		return nil
	}
	s.logger.Debugf("collection: %+v", collection)
	// check if supply is full
	if collection.Supply >= collection.Max {
		s.logger.Infof("collection %s supply is full, ignore inscription %d", o.Tick, inscriptionId)
		return nil
	}
	t, err := s.tokenUc.FindByInscriptionID(context.Background(), inscriptionId)
	if err != nil {
		return err
	}
	if len(t) > 0 {
		s.logger.Infof("token with inscription %d already processed, ignore", inscriptionId)
		return nil
	}
	// create token
	token := &biz.Token{
		Tick:           o.Tick,
		P:              biz.ProtocolTypeBRC721,
		TokenID:        collection.Supply + 1,
		TxHash:         info["genesis_transaction"].(string),
		BlockHeight:    info["genesis_height"].(uint64),
		BlockTime:      info["timestamp"].(time.Time),
		Address:        info["address"].(string),
		InscriptionID:  inscriptionId,
		InscriptionUID: info["uid"].(string),
		CollectionID:   collection.ID,
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

func (s *Syncer) processBRC721Update(inscriptionId int64, info map[string]interface{}) error {
	o := info["content"].(*parser.BRC721Update)
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

func (s *Syncer) parseInscriptions(inscriptionURL string) (string, error) {
	if inscriptionURL == "" {
		return "", nil
	}
	s.logger.Debugf("fetching %s", inscriptionURL)
	resp, err := httpGet(inscriptionURL)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var insUids uids
	links := doc.Find("div.thumbnails a")
	links.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		uid := strings.Replace(href, "/inscription/", "", -1)
		if uid == "" {
			return
		}
		// record the inscription uids, so that we can process them in order
		insUids = append(insUids, uid)
	})

	s.processChan <- insUids
	for _, insUid := range insUids {
		s.inscriptionUidChan <- insUid
	}
	// wait for the process to finish
	err = <-s.processFinishedChan
	if err != nil {
		return "", err
	}

	prevLink := doc.Find("a.next")
	if prevLink.Length() > 0 {
		href, _ := prevLink.Attr("href")
		inscriptionURL, _ = url.JoinPath(s.c.Server.Addr, href)
		return s.parseInscriptions(inscriptionURL)
	}
	return "", nil
}
