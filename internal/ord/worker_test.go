package ord

import (
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/adshao/ordinals-indexer/internal/ord/page"
	"github.com/adshao/ordinals-indexer/internal/ord/parser"
)

func TestWorkerParseBRC721DeployInscription(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"caller", log.DefaultCaller,
	)

	worker := &Worker{
		wid:        1,
		baseURL:    "http://localhost:8080",
		data:       nil,
		uidChan:    make(chan string, 10),
		resultChan: make(chan *result, 10),
		stopC:      make(chan struct{}),
		logger:     log.NewHelper(logger),
	}

	mockPageParser := &MockPageParser{}
	worker.pageParser = mockPageParser

	mockPageParser.On("Parse", mock.Anything).Once().Return(&page.Inscription{
		ID:            9553787,
		UID:           "3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0",
		Address:       "bc1putjs4fvkp3uaq6nhph7h2e7pmpwduq6zrxkt5kyyxe4rn47yrwzqup8lfu",
		OutputValue:   10000,
		ContentLength: 3440,
		ContentType:   "image/webp",
		Timestamp:     time.Date(2023, 5, 28, 3, 28, 17, 0, time.UTC),
		GenesisHeight: 791720,
		GenesisFee:    21000,
		GenesisTx:     "3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab",
		Location:      "3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0:0",
		Output:        "3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0",
		Offset:        0,
	}, nil)

	mockPageParser.On("Parse", mock.Anything).Return(&page.Content{
		Data: &parser.BRC721Deploy{
			P:    "brc-721",
			Op:   "deploy",
			Tick: "sato",
			Max:  "10000",
			Meta: &parser.BRC721Meta{
				Name:        "Satoshi",
				Description: "The ChatGPT 09/May/2023 Financial institutions on the precipice as three banks collapse in 2023.",
				Image:       "data:image/svg+xml;base64,PHN2ZyB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgaGVpZ2h0PSI2NCIgd2lkdGg9IjY0IiB2ZXJzaW9uPSIxLjEiIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyI+CjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAuMDA2MzA4NzYsLTAuMDAzMDE5ODQpIj4KPHBhdGggZmlsbD0iI2Y3OTMxYSIgZD0ibTYzLjAzMywzOS43NDRjLTQuMjc0LDE3LjE0My0yMS42MzcsMjcuNTc2LTM4Ljc4MiwyMy4zMDEtMTcuMTM4LTQuMjc0LTI3LjU3MS0yMS42MzgtMjMuMjk1LTM4Ljc4LDQuMjcyLTE3LjE0NSwyMS42MzUtMjcuNTc5LDM4Ljc3NS0yMy4zMDUsMTcuMTQ0LDQuMjc0LDI3LjU3NiwyMS42NCwyMy4zMDIsMzguNzg0eiIvPgo8cGF0aCBmaWxsPSIjRkZGIiBkPSJtNDYuMTAzLDI3LjQ0NGMwLjYzNy00LjI1OC0yLjYwNS02LjU0Ny03LjAzOC04LjA3NGwxLjQzOC01Ljc2OC0zLjUxMS0wLjg3NS0xLjQsNS42MTZjLTAuOTIzLTAuMjMtMS44NzEtMC40NDctMi44MTMtMC42NjJsMS40MS01LjY1My0zLjUwOS0wLjg3NS0xLjQzOSw1Ljc2NmMtMC43NjQtMC4xNzQtMS41MTQtMC4zNDYtMi4yNDItMC41MjdsMC4wMDQtMC4wMTgtNC44NDItMS4yMDktMC45MzQsMy43NXMyLjYwNSwwLjU5NywyLjU1LDAuNjM0YzEuNDIyLDAuMzU1LDEuNjc5LDEuMjk2LDEuNjM2LDIuMDQybC0xLjYzOCw2LjU3MWMwLjA5OCwwLjAyNSwwLjIyNSwwLjA2MSwwLjM2NSwwLjExNy0wLjExNy0wLjAyOS0wLjI0Mi0wLjA2MS0wLjM3MS0wLjA5MmwtMi4yOTYsOS4yMDVjLTAuMTc0LDAuNDMyLTAuNjE1LDEuMDgtMS42MDksMC44MzQsMC4wMzUsMC4wNTEtMi41NTItMC42MzctMi41NTItMC42MzdsLTEuNzQzLDQuMDE5LDQuNTY5LDEuMTM5YzAuODUsMC4yMTMsMS42ODMsMC40MzYsMi41MDMsMC42NDZsLTEuNDUzLDUuODM0LDMuNTA3LDAuODc1LDEuNDM5LTUuNzcyYzAuOTU4LDAuMjYsMS44ODgsMC41LDIuNzk4LDAuNzI2bC0xLjQzNCw1Ljc0NSwzLjUxMSwwLjg3NSwxLjQ1My01LjgyM2M1Ljk4NywxLjEzMywxMC40ODksMC42NzYsMTIuMzg0LTQuNzM5LDEuNTI3LTQuMzYtMC4wNzYtNi44NzUtMy4yMjYtOC41MTUsMi4yOTQtMC41MjksNC4wMjItMi4wMzgsNC40ODMtNS4xNTV6bS04LjAyMiwxMS4yNDljLTEuMDg1LDQuMzYtOC40MjYsMi4wMDMtMTAuODA2LDEuNDEybDEuOTI4LTcuNzI5YzIuMzgsMC41OTQsMTAuMDEyLDEuNzcsOC44NzgsNi4zMTd6bTEuMDg2LTExLjMxMmMtMC45OSwzLjk2Ni03LjEsMS45NTEtOS4wODIsMS40NTdsMS43NDgtNy4wMWMxLjk4MiwwLjQ5NCw4LjM2NSwxLjQxNiw3LjMzNCw1LjU1M3oiLz4KPC9nPgo8L3N2Zz4=",
			},
		},
		Type: parser.NameBRC721Deploy,
	}, nil)

	info, err := worker.parseInscriptionInfo("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0")
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(9553787), info.ID, "inscription_id")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0", info.UID, "uid")
	r.Equal("bc1putjs4fvkp3uaq6nhph7h2e7pmpwduq6zrxkt5kyyxe4rn47yrwzqup8lfu", info.Address, "address")
	r.Equal(uint64(10000), info.OutputValue, "output_value")
	r.Equal(time.Date(2023, 5, 28, 3, 28, 17, 0, time.UTC), info.Timestamp, "timestamp")
	r.Equal(uint64(791720), info.GenesisHeight, "genesis_height")
	r.Equal(uint64(21000), info.GenesisFee, "genesis_fee")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab", info.GenesisTx, "genesis_transaction")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0:0", info.Location, "location")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0", info.Output, "output")
	r.Equal(uint64(0), info.Offset, "offset")

	r.NotNil(info.Content, "content")
	r.Equal("brc-721-deploy", info.Content.Type, "content_parser")
	o, ok := info.Content.Data.(*parser.BRC721Deploy)
	r.True(ok, "content")
	r.Equal("brc-721", o.P, "content.P")
	r.Equal("deploy", o.Op, "content.Op")
	r.Equal("sato", o.Tick, "content.Tick")
	r.Equal("10000", o.Max, "content.Max")
	r.Equal("Satoshi", o.Meta.Name, "content.Meta.Name")
	r.Equal("The ChatGPT 09/May/2023 Financial institutions on the precipice as three banks collapse in 2023.", o.Meta.Description, "content.Meta.Description")
	r.Equal("data:image/svg+xml;base64,PHN2ZyB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgaGVpZ2h0PSI2NCIgd2lkdGg9IjY0IiB2ZXJzaW9uPSIxLjEiIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyI+CjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAuMDA2MzA4NzYsLTAuMDAzMDE5ODQpIj4KPHBhdGggZmlsbD0iI2Y3OTMxYSIgZD0ibTYzLjAzMywzOS43NDRjLTQuMjc0LDE3LjE0My0yMS42MzcsMjcuNTc2LTM4Ljc4MiwyMy4zMDEtMTcuMTM4LTQuMjc0LTI3LjU3MS0yMS42MzgtMjMuMjk1LTM4Ljc4LDQuMjcyLTE3LjE0NSwyMS42MzUtMjcuNTc5LDM4Ljc3NS0yMy4zMDUsMTcuMTQ0LDQuMjc0LDI3LjU3NiwyMS42NCwyMy4zMDIsMzguNzg0eiIvPgo8cGF0aCBmaWxsPSIjRkZGIiBkPSJtNDYuMTAzLDI3LjQ0NGMwLjYzNy00LjI1OC0yLjYwNS02LjU0Ny03LjAzOC04LjA3NGwxLjQzOC01Ljc2OC0zLjUxMS0wLjg3NS0xLjQsNS42MTZjLTAuOTIzLTAuMjMtMS44NzEtMC40NDctMi44MTMtMC42NjJsMS40MS01LjY1My0zLjUwOS0wLjg3NS0xLjQzOSw1Ljc2NmMtMC43NjQtMC4xNzQtMS41MTQtMC4zNDYtMi4yNDItMC41MjdsMC4wMDQtMC4wMTgtNC44NDItMS4yMDktMC45MzQsMy43NXMyLjYwNSwwLjU5NywyLjU1LDAuNjM0YzEuNDIyLDAuMzU1LDEuNjc5LDEuMjk2LDEuNjM2LDIuMDQybC0xLjYzOCw2LjU3MWMwLjA5OCwwLjAyNSwwLjIyNSwwLjA2MSwwLjM2NSwwLjExNy0wLjExNy0wLjAyOS0wLjI0Mi0wLjA2MS0wLjM3MS0wLjA5MmwtMi4yOTYsOS4yMDVjLTAuMTc0LDAuNDMyLTAuNjE1LDEuMDgtMS42MDksMC44MzQsMC4wMzUsMC4wNTEtMi41NTItMC42MzctMi41NTItMC42MzdsLTEuNzQzLDQuMDE5LDQuNTY5LDEuMTM5YzAuODUsMC4yMTMsMS42ODMsMC40MzYsMi41MDMsMC42NDZsLTEuNDUzLDUuODM0LDMuNTA3LDAuODc1LDEuNDM5LTUuNzcyYzAuOTU4LDAuMjYsMS44ODgsMC41LDIuNzk4LDAuNzI2bC0xLjQzNCw1Ljc0NSwzLjUxMSwwLjg3NSwxLjQ1My01LjgyM2M1Ljk4NywxLjEzMywxMC40ODksMC42NzYsMTIuMzg0LTQuNzM5LDEuNTI3LTQuMzYtMC4wNzYtNi44NzUtMy4yMjYtOC41MTUsMi4yOTQtMC41MjksNC4wMjItMi4wMzgsNC40ODMtNS4xNTV6bS04LjAyMiwxMS4yNDljLTEuMDg1LDQuMzYtOC40MjYsMi4wMDMtMTAuODA2LDEuNDEybDEuOTI4LTcuNzI5YzIuMzgsMC41OTQsMTAuMDEyLDEuNzcsOC44NzgsNi4zMTd6bTEuMDg2LTExLjMxMmMtMC45OSwzLjk2Ni03LjEsMS45NTEtOS4wODIsMS40NTdsMS43NDgtNy4wMWMxLjk4MiwwLjQ5NCw4LjM2NSwxLjQxNiw3LjMzNCw1LjU1M3oiLz4KPC9nPgo8L3N2Zz4=", o.Meta.Image, "content.Meta.Image")
	r.Nil(o.BaseURI, "content.BaseURI")
}

func TestWorkerParseBRC721MintInscription(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"caller", log.DefaultCaller,
	)

	worker := &Worker{
		wid:        1,
		baseURL:    "http://localhost:8080",
		data:       nil,
		uidChan:    make(chan string, 10),
		resultChan: make(chan *result, 10),
		stopC:      make(chan struct{}),
		logger:     log.NewHelper(logger),
	}

	mockPageParser := &MockPageParser{}
	worker.pageParser = mockPageParser
	mockPageParser.On("Parse", mock.Anything).Once().Return(&page.Inscription{
		ID:            4986756,
		UID:           "8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0",
		Address:       "bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf",
		OutputValue:   10000,
		ContentLength: 46,
		ContentType:   "application/json",
		Timestamp:     time.Date(2023, 5, 9, 7, 22, 28, 0, time.UTC),
		GenesisHeight: 788906,
		GenesisFee:    30870,
		GenesisTx:     "8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564",
		Location:      "8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0:0",
		Output:        "8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0",
		Offset:        0,
	}, nil)

	mockPageParser.On("Parse", mock.Anything).Return(&page.Content{
		Data: &parser.BRC721Mint{
			P:    "brc-721",
			Op:   "mint",
			Tick: "sato",
		},
		Type: parser.NameBRC721Mint,
	}, nil)

	info, err := worker.parseInscriptionInfo("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0")
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(4986756), info.ID, "inscription_id")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0", info.UID, "uid")
	r.Equal("bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf", info.Address, "address")
	r.Equal(uint64(10000), info.OutputValue, "output_value")
	r.Equal(time.Date(2023, 5, 9, 7, 22, 28, 0, time.UTC), info.Timestamp, "timestamp")
	r.Equal(uint64(788906), info.GenesisHeight, "genesis_height")
	r.Equal(uint64(30870), info.GenesisFee, "genesis_fee")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564", info.GenesisTx, "genesis_transaction")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0:0", info.Location, "location")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0", info.Output, "output")
	r.Equal(uint64(0), info.Offset, "offset")

	r.NotNil(info.Content, "content")
	r.Equal("brc-721-mint", info.Content.Type, "content_parser")
	o, ok := info.Content.Data.(*parser.BRC721Mint)
	r.True(ok, "content")
	r.Equal("brc-721", o.P, "content.P")
	r.Equal("mint", o.Op, "content.Op")
	r.Equal("sato", o.Tick, "content.Tick")
}

func TestWorkerParseBRC721DeployWithMeta(t *testing.T) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"caller", log.DefaultCaller,
	)

	worker := &Worker{
		wid:        1,
		baseURL:    "http://localhost:8080",
		data:       nil,
		uidChan:    make(chan string, 10),
		resultChan: make(chan *result, 10),
		stopC:      make(chan struct{}),
		logger:     log.NewHelper(logger),
	}

	mockPageParser := &MockPageParser{}
	worker.pageParser = mockPageParser
	mockPageParser.On("Parse", mock.Anything).Once().Return(&page.Inscription{
		ID:            9553787,
		UID:           "423992a9468b935e2e04234c0c5232f3d8f7acd0e1d867e9797cee3941bdc702i0",
		Address:       "bc1p0rwqj2dwtsurejsazqx9fp62fk0esytrlwe7atftadenv369v0ksnkngg3",
		OutputValue:   546,
		ContentLength: 362,
		ContentType:   "text/plain;charset=utf-8",
		Timestamp:     time.Date(2023, 5, 10, 4, 12, 43, 0, time.UTC),
		GenesisHeight: 789018,
		GenesisFee:    59736,
		GenesisTx:     "e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913",
		Location:      "e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0:0",
		Output:        "e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0",
		Offset:        0,
	}, nil)
	mockPageParser.On("Parse", mock.Anything).Return(&page.Content{
		Data: &parser.BRC721Deploy{
			P:    "brc-721",
			Op:   "deploy",
			Tick: "BNFT",
			Max:  "200913",
			Meta: &parser.BRC721Meta{
				Name:        "Ordinals",
				Description: "Taking the birth date of Bitcoin as the Genesis Yuan, brc721 will open a new round of blockchain legends.",
				Image:       "https://www.bitcoin.com/images/uploads/get-started-what-is-bitcoin-lg@2x.png",
			},
		},
		Type: parser.NameBRC721Deploy,
	}, nil)

	info, err := worker.parseInscriptionInfo("423992a9468b935e2e04234c0c5232f3d8f7acd0e1d867e9797cee3941bdc702i0")
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(9553787), info.ID, "inscription_id")
	r.Equal("423992a9468b935e2e04234c0c5232f3d8f7acd0e1d867e9797cee3941bdc702i0", info.UID, "uid")
	r.Equal("bc1p0rwqj2dwtsurejsazqx9fp62fk0esytrlwe7atftadenv369v0ksnkngg3", info.Address, "address")
	r.Equal(uint64(546), info.OutputValue, "output_value")
	r.Equal(time.Date(2023, 5, 10, 4, 12, 43, 0, time.UTC), info.Timestamp, "timestamp")
	r.Equal(uint64(789018), info.GenesisHeight, "genesis_height")
	r.Equal(uint64(59736), info.GenesisFee, "genesis_fee")
	r.Equal("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913", info.GenesisTx, "genesis_transaction")
	r.Equal("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0:0", info.Location, "location")
	r.Equal("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0", info.Output, "output")
	r.Equal(uint64(0), info.Offset, "offset")

	r.NotNil(info.Content, "content")
	r.Equal("brc-721-deploy", info.Content.Type, "content_parser")
	o, ok := info.Content.Data.(*parser.BRC721Deploy)
	r.True(ok, "content")
	r.Equal("brc-721", o.P, "content.P")
	r.Equal("deploy", o.Op, "content.Op")
	r.Equal("BNFT", o.Tick, "content.Tick")
	r.Equal("200913", o.Max, "content.Max")
	r.Equal("Ordinals", o.Meta.Name, "content.Meta.Name")
	r.Equal("Taking the birth date of Bitcoin as the Genesis Yuan, brc721 will open a new round of blockchain legends.", o.Meta.Description, "content.Meta.Description")
	r.Nil(o.BaseURI, "content.BaseURI")
}
