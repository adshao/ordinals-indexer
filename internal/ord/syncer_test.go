package ord

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/ord/page"
	"github.com/adshao/ordinals-indexer/internal/ord/parser"
)

type MockPageParser struct {
	mock.Mock
}

func (m *MockPageParser) Parse(p page.Page) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0), args.Error(1)
}

type MockPage struct {
	mock.Mock
}

func (m *MockPage) URL() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockPage) Parse(r []byte) (interface{}, bool, error) {
	args := m.Called(r)
	return args.Get(0), args.Bool(1), args.Error(2)
}

func TestSyncerParseInscriptions(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"caller", log.DefaultCaller,
	)
	c := &conf.Ord{
		Worker: &conf.Ord_Worker{
			Concurrency: 10,
		},
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	syncer, _, _ := NewSyncer(c, nil, nil, nil, logger)
	concurrency := 2
	syncer.inscriptionUidChan = make(chan string, concurrency)
	syncer.resultChan = make(chan *result, concurrency)
	syncer.processChan = make(chan uids)
	syncer.processFinishedChan = make(chan error)

	mockPageParser := &MockPageParser{}
	nextID := int64(4984502)
	mockPageParser.On("Parse", mock.Anything).Once().Return(&page.Inscriptions{
		UIDs: []string{
			"347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0",
			"9bb83fa001542416bdf1eaeed41699f619110e9b68fb25b5cd2628dfb328c063i0",
			"e9948dd04f3b63810e52c77f431daf6179cb06219724493cf8e5349b6c3cb562i0",
			"e616c29b8bd9c6da52866882a04213278854e86e39c7704032ab2b9cdabdc860i0",
			"c06b2dcb0b92b9fa9b509c4128745dc28a4ee8917e8d8df00ed5b063dc28c55ei0",
			"6d487246d3f279e053ebf7b92bf1ed949ee63935a6b8160a05d6cae7af4b625ei0",
		},
		NextID: &nextID,
	}, nil)

	mockPageParser.On("Parse", mock.Anything).Return(&page.Inscriptions{
		UIDs: []string{
			"09af268da3a45bb20f49296904f73ec70e1ead6676ba65c97036dd118d7fdcf0i0",
			"a6bf3307d613fe515b28333aa54a0e844bf28e5f6beeddd1dc0d026b276094f0i0",
			"2199307ba2e25cff294d4da4d1c06727d0e0bca05c48b2497b66f271df57f3efi0",
			"a88bdeee9dc55ef00d599e96176a59a416b939890956601087150022ae7bb0efi0",
			"6279bba4ee7fe35f20d6e2a3df989d0528aeb31a774f76d6d0001556272a66eei0",
			"959e6e747a6597d811b5910837cd6e3ac8bc58709ec80460c05be6c7600df9ebi0",
			"f5d970d0a009bb140db9437cbff5157d11c95bba49e2287c123d53bbee21ecebi0",
		},
		NextID: nil,
	}, nil)
	syncer.pageParser = mockPageParser

	go func(syncer *Syncer) {
		i := 0
		uids := make([]string, 0)
		uidNum := 0
		for {
			select {
			case <-syncer.stopC:
				t.Logf("stopping")
				return
			case uids = <-syncer.processChan:
				t.Logf("processing %d uids, i=%d", len(uids), i)
			case <-syncer.inscriptionUidChan:
				uidNum++
			default:
				if len(uids) == uidNum && uidNum != 0 {
					if i == 0 {
						r := require.New(t)
						r.Equal(6, len(uids))
						r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", uids[0])
						r.Equal("9bb83fa001542416bdf1eaeed41699f619110e9b68fb25b5cd2628dfb328c063i0", uids[1])
						r.Equal("6d487246d3f279e053ebf7b92bf1ed949ee63935a6b8160a05d6cae7af4b625ei0", uids[5])
					}
					if i == 1 {
						r := require.New(t)
						r.Equal(7, len(uids))
						r.Equal("09af268da3a45bb20f49296904f73ec70e1ead6676ba65c97036dd118d7fdcf0i0", uids[0])
						r.Equal("a6bf3307d613fe515b28333aa54a0e844bf28e5f6beeddd1dc0d026b276094f0i0", uids[1])
						r.Equal("f5d970d0a009bb140db9437cbff5157d11c95bba49e2287c123d53bbee21ecebi0", uids[6])
					}
					i++
					uidNum = 0
					syncer.processFinishedChan <- nil
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(syncer)

	err := syncer.parseInscriptions(4984402)
	r := require.New(t)
	r.NoError(err)
	close(syncer.stopC)
}

type brc721SigTestSuite struct {
	suite.Suite
	c            *conf.Ord
	collectionUc *biz.CollectionUsecase
	tokenUc      *biz.TokenUsecase
	d            *data.Data
	cleanup      func()
	syncer       *Syncer
	logger       log.Logger
	info         *page.Inscription
}

func TestBRC721Suite(t *testing.T) {
	suite.Run(t, new(brc721SigTestSuite))
}

func (s *brc721SigTestSuite) SetupSuite() {
	s.c = &conf.Ord{
		Worker: &conf.Ord_Worker{
			Concurrency: 10,
		},
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	s.d = &data.Data{}
	logger := log.GetLogger()
	s.logger = logger
	collectionRepo := data.NewCollectionRepo(s.d, logger)
	tokenRepo := data.NewTokenRepo(s.d, logger)
	s.collectionUc = biz.NewCollectionUsecase(collectionRepo, logger)
	s.tokenUc = biz.NewTokenUsecase(tokenRepo, logger)
}

func (s *brc721SigTestSuite) SetupTest() {
	d, cleanup := data.NewTData(s.T())
	s.cleanup = cleanup
	*s.d = *d
	s.syncer, _, _ = NewSyncer(s.c, s.d, s.collectionUc, s.tokenUc, s.logger)
	info := &page.Inscription{
		ID:            4984402,
		UID:           "347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0",
		Address:       "bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf",
		OutputValue:   10000,
		ContentLength: 2167,
		ContentType:   "application/json",
		Timestamp:     time.Unix(1624296000, 0),
		GenesisHeight: 788904,
		GenesisFee:    143010,
		GenesisTx:     "347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564",
		Location:      "347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0:0",
		Output:        "347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0",
		Offset:        0,
	}
	contentData := &parser.BRC721Deploy{
		P:    "brc-721",
		Op:   "deploy",
		Tick: "ordinals",
		Max:  "1000",
	}
	baseURI := "https://ordinals.io/"
	contentData.BaseURI = &baseURI
	info.Content = &page.Content{
		Data: contentData,
		Type: parser.NameBRC721Deploy,
	}
	s.info = info
}

func (s *brc721SigTestSuite) TearDownTest() {
	s.cleanup()
}

func (s *brc721SigTestSuite) TestNewDeploy() {
	err := s.syncer.processBRC721Deploy(s.info)
	r := s.Require()
	r.NoError(err)
	collections, err := s.collectionUc.GetCollectionByInscriptionID(context.Background(), 4984402)
	r.NoError(err)
	r.Len(collections, 1)
	collection := collections[0]
	r.NotNil(collection)
	r.True(collection.ID > 0)
	r.Equal("brc-721", collection.P)
	r.Equal("ordinals", collection.Tick)
	r.Equal(uint64(1000), collection.Max)
	r.Equal(uint64(0), collection.Supply)
	r.Equal("https://ordinals.io/", collection.BaseURI)
	r.Equal("", collection.Name)
	r.Equal("", collection.Description)
	r.Equal("", collection.Image)
	r.Len(collection.Attributes, 0)
	r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564", collection.TxHash)
	r.Equal(uint64(788904), collection.BlockHeight)
	r.Equal(int64(1624296000), collection.BlockTime.Unix())
	r.Equal("bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf", collection.Address)
	r.Equal(int64(4984402), collection.InscriptionID)
	r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", collection.InscriptionUID)
	r.Equal("", collection.Sig.PubKey)
	r.Len(collection.Sig.Fields, 0)
}

func (s *brc721SigTestSuite) newCollection() *biz.Collection {
	return &biz.Collection{
		P:              "brc-721",
		Tick:           "ordinals",
		Max:            1000,
		Supply:         0,
		BaseURI:        "https://ordinals.io/",
		Name:           "",
		Description:    "",
		Image:          "",
		Attributes:     nil,
		TxHash:         "347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564",
		BlockHeight:    788904,
		BlockTime:      time.Unix(1624296000, 0),
		Address:        "bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf",
		InscriptionID:  4984402,
		InscriptionUID: "347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0",
	}
}

func (s *brc721SigTestSuite) TestDeployWithExistingCollection() {
	collection := s.newCollection()
	newCollection, err := s.collectionUc.CreateCollection(context.Background(), collection)
	r := s.Require()
	r.NoError(err)
	r.NotNil(newCollection)
	r.True(newCollection.ID > 0)
	err = s.syncer.processBRC721Deploy(s.info)
	r.NoError(err)
	collections, err := s.collectionUc.ListCollections(context.Background(), &biz.CollectionListOption{})
	r.NoError(err)
	r.Len(collections, 1)
	r.Equal(newCollection.ID, collections[0].ID)
}
