package page

import (
	"testing"

	"github.com/adshao/go-brc721/sig"
	"github.com/stretchr/testify/require"

	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/ord/parser"
)

func TestContentPageBRC721Deploy(t *testing.T) {
	defer resetHTTPGet()

	c := &conf.Ord{
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	pp := &pageParser{
		httpGet: mockHTTPGet,
		c:       c,
	}
	brc721DeployContent := []byte(`{"p": "brc-721", "op": "deploy", "tick": "sato", "max": "10000", "meta": {"name": "Satoshi", "description": "The ChatGPT 09/May/2023 Financial institutions on the precipice as three banks collapse in 2023.", "image": "data:image/svg+xml;base64,PHN2ZyB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgaGVpZ2h0PSI2NCIgd2lkdGg9IjY0IiB2ZXJzaW9uPSIxLjEiIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyI+CjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAuMDA2MzA4NzYsLTAuMDAzMDE5ODQpIj4KPHBhdGggZmlsbD0iI2Y3OTMxYSIgZD0ibTYzLjAzMywzOS43NDRjLTQuMjc0LDE3LjE0My0yMS42MzcsMjcuNTc2LTM4Ljc4MiwyMy4zMDEtMTcuMTM4LTQuMjc0LTI3LjU3MS0yMS42MzgtMjMuMjk1LTM4Ljc4LDQuMjcyLTE3LjE0NSwyMS42MzUtMjcuNTc5LDM4Ljc3NS0yMy4zMDUsMTcuMTQ0LDQuMjc0LDI3LjU3NiwyMS42NCwyMy4zMDIsMzguNzg0eiIvPgo8cGF0aCBmaWxsPSIjRkZGIiBkPSJtNDYuMTAzLDI3LjQ0NGMwLjYzNy00LjI1OC0yLjYwNS02LjU0Ny03LjAzOC04LjA3NGwxLjQzOC01Ljc2OC0zLjUxMS0wLjg3NS0xLjQsNS42MTZjLTAuOTIzLTAuMjMtMS44NzEtMC40NDctMi44MTMtMC42NjJsMS40MS01LjY1My0zLjUwOS0wLjg3NS0xLjQzOSw1Ljc2NmMtMC43NjQtMC4xNzQtMS41MTQtMC4zNDYtMi4yNDItMC41MjdsMC4wMDQtMC4wMTgtNC44NDItMS4yMDktMC45MzQsMy43NXMyLjYwNSwwLjU5NywyLjU1LDAuNjM0YzEuNDIyLDAuMzU1LDEuNjc5LDEuMjk2LDEuNjM2LDIuMDQybC0xLjYzOCw2LjU3MWMwLjA5OCwwLjAyNSwwLjIyNSwwLjA2MSwwLjM2NSwwLjExNy0wLjExNy0wLjAyOS0wLjI0Mi0wLjA2MS0wLjM3MS0wLjA5MmwtMi4yOTYsOS4yMDVjLTAuMTc0LDAuNDMyLTAuNjE1LDEuMDgtMS42MDksMC44MzQsMC4wMzUsMC4wNTEtMi41NTItMC42MzctMi41NTItMC42MzdsLTEuNzQzLDQuMDE5LDQuNTY5LDEuMTM5YzAuODUsMC4yMTMsMS42ODMsMC40MzYsMi41MDMsMC42NDZsLTEuNDUzLDUuODM0LDMuNTA3LDAuODc1LDEuNDM5LTUuNzcyYzAuOTU4LDAuMjYsMS44ODgsMC41LDIuNzk4LDAuNzI2bC0xLjQzNCw1Ljc0NSwzLjUxMSwwLjg3NSwxLjQ1My01LjgyM2M1Ljk4NywxLjEzMywxMC40ODksMC42NzYsMTIuMzg0LTQuNzM5LDEuNTI3LTQuMzYtMC4wNzYtNi44NzUtMy4yMjYtOC41MTUsMi4yOTQtMC41MjksNC4wMjItMi4wMzgsNC40ODMtNS4xNTV6bS04LjAyMiwxMS4yNDljLTEuMDg1LDQuMzYtOC40MjYsMi4wMDMtMTAuODA2LDEuNDEybDEuOTI4LTcuNzI5YzIuMzgsMC41OTQsMTAuMDEyLDEuNzcsOC44NzgsNi4zMTd6bTEuMDg2LTExLjMxMmMtMC45OSwzLjk2Ni03LjEsMS45NTEtOS4wODIsMS40NTdsMS43NDgtNy4wMWMxLjk4MiwwLjQ5NCw4LjM2NSwxLjQxNiw3LjMzNCw1LjU1M3oiLz4KPC9nPgo8L3N2Zz4="}}`)
	mockHTTPResult("http://localhost:8080/content/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", brc721DeployContent)

	page := NewContentPage("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0")
	data, err := pp.Parse(page)

	r := require.New(t)
	r.Nil(err)
	content, ok := data.(*Content)
	r.True(ok)
	r.Equal("brc-721-deploy", content.Type)
	o, ok := content.Data.(*parser.BRC721Deploy)
	r.True(ok)
	r.Equal("brc-721", o.P)
	r.Equal("deploy", o.Op)
	r.Equal("sato", o.Tick)
	r.Equal("10000", o.Max)
	r.Nil(o.BaseURI)
	r.NotNil(o.Meta)
	r.Equal("Satoshi", o.Meta.Name)
	r.Equal("The ChatGPT 09/May/2023 Financial institutions on the precipice as three banks collapse in 2023.", o.Meta.Description)
	r.Equal("data:image/svg+xml;base64,PHN2ZyB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgaGVpZ2h0PSI2NCIgd2lkdGg9IjY0IiB2ZXJzaW9uPSIxLjEiIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyI+CjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAuMDA2MzA4NzYsLTAuMDAzMDE5ODQpIj4KPHBhdGggZmlsbD0iI2Y3OTMxYSIgZD0ibTYzLjAzMywzOS43NDRjLTQuMjc0LDE3LjE0My0yMS42MzcsMjcuNTc2LTM4Ljc4MiwyMy4zMDEtMTcuMTM4LTQuMjc0LTI3LjU3MS0yMS42MzgtMjMuMjk1LTM4Ljc4LDQuMjcyLTE3LjE0NSwyMS42MzUtMjcuNTc5LDM4Ljc3NS0yMy4zMDUsMTcuMTQ0LDQuMjc0LDI3LjU3NiwyMS42NCwyMy4zMDIsMzguNzg0eiIvPgo8cGF0aCBmaWxsPSIjRkZGIiBkPSJtNDYuMTAzLDI3LjQ0NGMwLjYzNy00LjI1OC0yLjYwNS02LjU0Ny03LjAzOC04LjA3NGwxLjQzOC01Ljc2OC0zLjUxMS0wLjg3NS0xLjQsNS42MTZjLTAuOTIzLTAuMjMtMS44NzEtMC40NDctMi44MTMtMC42NjJsMS40MS01LjY1My0zLjUwOS0wLjg3NS0xLjQzOSw1Ljc2NmMtMC43NjQtMC4xNzQtMS41MTQtMC4zNDYtMi4yNDItMC41MjdsMC4wMDQtMC4wMTgtNC44NDItMS4yMDktMC45MzQsMy43NXMyLjYwNSwwLjU5NywyLjU1LDAuNjM0YzEuNDIyLDAuMzU1LDEuNjc5LDEuMjk2LDEuNjM2LDIuMDQybC0xLjYzOCw2LjU3MWMwLjA5OCwwLjAyNSwwLjIyNSwwLjA2MSwwLjM2NSwwLjExNy0wLjExNy0wLjAyOS0wLjI0Mi0wLjA2MS0wLjM3MS0wLjA5MmwtMi4yOTYsOS4yMDVjLTAuMTc0LDAuNDMyLTAuNjE1LDEuMDgtMS42MDksMC44MzQsMC4wMzUsMC4wNTEtMi41NTItMC42MzctMi41NTItMC42MzdsLTEuNzQzLDQuMDE5LDQuNTY5LDEuMTM5YzAuODUsMC4yMTMsMS42ODMsMC40MzYsMi41MDMsMC42NDZsLTEuNDUzLDUuODM0LDMuNTA3LDAuODc1LDEuNDM5LTUuNzcyYzAuOTU4LDAuMjYsMS44ODgsMC41LDIuNzk4LDAuNzI2bC0xLjQzNCw1Ljc0NSwzLjUxMSwwLjg3NSwxLjQ1My01LjgyM2M1Ljk4NywxLjEzMywxMC40ODksMC42NzYsMTIuMzg0LTQuNzM5LDEuNTI3LTQuMzYtMC4wNzYtNi44NzUtMy4yMjYtOC41MTUsMi4yOTQtMC41MjksNC4wMjItMi4wMzgsNC40ODMtNS4xNTV6bS04LjAyMiwxMS4yNDljLTEuMDg1LDQuMzYtOC40MjYsMi4wMDMtMTAuODA2LDEuNDEybDEuOTI4LTcuNzI5YzIuMzgsMC41OTQsMTAuMDEyLDEuNzcsOC44NzgsNi4zMTd6bTEuMDg2LTExLjMxMmMtMC45OSwzLjk2Ni03LjEsMS45NTEtOS4wODIsMS40NTdsMS43NDgtNy4wMWMxLjk4MiwwLjQ5NCw4LjM2NSwxLjQxNiw3LjMzNCw1LjU1M3oiLz4KPC9nPgo8L3N2Zz4=", o.Meta.Image)
}

func TestContentPageBRC721DeployWithSig(t *testing.T) {
	defer resetHTTPGet()

	c := &conf.Ord{
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	pp := &pageParser{
		httpGet: mockHTTPGet,
		c:       c,
	}
	brc721DeployContent := []byte(`{"p": "brc-721", "op": "deploy", "tick": "sato", "max": "10000", "buri": "https://abc/", "sig": {"pk": "0379f79637ec1cc5375c4e269e9d70eda426b5ecba5d4088234a89e8943dc4aa9f", "fields": ["rec", "uid"]}}`)
	mockHTTPResult("http://localhost:8080/content/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", brc721DeployContent)

	page := NewContentPage("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0")
	data, err := pp.Parse(page)

	r := require.New(t)
	r.Nil(err)
	content, ok := data.(*Content)
	r.True(ok)
	r.Equal("brc-721-deploy", content.Type)
	o, ok := content.Data.(*parser.BRC721Deploy)
	r.True(ok)
	r.Equal("brc-721", o.P)
	r.Equal("deploy", o.Op)
	r.Equal("sato", o.Tick)
	r.Equal("10000", o.Max)
	r.Equal("https://abc/", *o.BaseURI)
	r.Nil(o.Meta)
	r.NotNil(o.Sig)
	r.Equal("0379f79637ec1cc5375c4e269e9d70eda426b5ecba5d4088234a89e8943dc4aa9f", o.Sig.PubKey)
	r.Len(o.Sig.Fields, 2)
	r.Equal(sig.SigFieldReceiver, o.Sig.Fields[0])
	r.Equal(sig.SigFieldUid, o.Sig.Fields[1])
}

func TestContentPageBRC721DeployWithSigInvalidFields(t *testing.T) {
	defer resetHTTPGet()

	c := &conf.Ord{
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	pp := &pageParser{
		httpGet: mockHTTPGet,
		c:       c,
	}
	brc721DeployContent := []byte(`{"p": "brc-721", "op": "deploy", "tick": "sato", "max": "10000", "buri": "https://abc/", "sig": {"pk": "0379f79637ec1cc5375c4e269e9d70eda426b5ecba5d4088234a89e8943dc4aa9f", "fields": ["rec123", "uid"]}}`)
	mockHTTPResult("http://localhost:8080/content/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", brc721DeployContent)

	page := NewContentPage("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0")
	data, err := pp.Parse(page)
	r := require.New(t)
	r.Nil(err)
	content, ok := data.(*Content)
	r.True(ok)
	r.Equal("brc-721-deploy", content.Type)
	o, ok := content.Data.(*parser.BRC721Deploy)
	r.True(ok)
	r.Equal("brc-721", o.P)
	r.Equal("deploy", o.Op)
	r.Equal("sato", o.Tick)
	r.Equal("10000", o.Max)
	r.Equal("https://abc/", *o.BaseURI)
	r.Nil(o.Meta)
	r.NotNil(o.Sig)
	r.Equal("0379f79637ec1cc5375c4e269e9d70eda426b5ecba5d4088234a89e8943dc4aa9f", o.Sig.PubKey)
	r.Len(o.Sig.Fields, 2)
	r.Equal(sig.SigField("rec123"), o.Sig.Fields[0])
	r.Equal(sig.SigFieldUid, o.Sig.Fields[1])
}

func TestContentPageBRC721Mint(t *testing.T) {
	defer resetHTTPGet()

	c := &conf.Ord{
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	pp := &pageParser{
		httpGet: mockHTTPGet,
		c:       c,
	}
	brc721DeployContent := []byte(`{"p": "brc-721", "op": "mint", "tick": "sato"}`)
	mockHTTPResult("http://localhost:8080/content/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0", brc721DeployContent)

	page := NewContentPage("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0")
	data, err := pp.Parse(page)

	r := require.New(t)
	r.Nil(err)
	content, ok := data.(*Content)
	r.True(ok)
	r.Equal("brc-721-mint", content.Type)
	o, ok := content.Data.(*parser.BRC721Mint)
	r.True(ok)
	r.Equal("brc-721", o.P)
	r.Equal("mint", o.Op)
	r.Equal("sato", o.Tick)
}

func TestContentPageBRC721Update(t *testing.T) {
	defer resetHTTPGet()

	c := &conf.Ord{
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	pp := &pageParser{
		httpGet: mockHTTPGet,
		c:       c,
	}
	brc721DeployContent := []byte(`{
		"p": "brc-721",
		"op": "update",
		"tick": "ordinals",
		"upd": false,
		"buri": "https://ipfs.io/abc/"}`)
	mockHTTPResult("http://localhost:8080/content/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0", brc721DeployContent)

	page := NewContentPage("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0")
	data, err := pp.Parse(page)

	r := require.New(t)
	r.Nil(err)
	content, ok := data.(*Content)
	r.True(ok)
	r.Equal("brc-721-update", content.Type)
	o, ok := content.Data.(*parser.BRC721Update)
	r.True(ok)
	r.Equal("brc-721", o.P)
	r.Equal("update", o.Op)
	r.Equal("ordinals", o.Tick)
	r.Equal("https://ipfs.io/abc/", *o.BaseURI)
	r.Nil(o.Meta)
}
