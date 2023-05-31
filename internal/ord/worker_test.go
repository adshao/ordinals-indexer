package ord

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"

	"github.com/adshao/ordinals-indexer/internal/ord/parser"
)

func TestWorkerParseBRC721DeployInscription(t *testing.T) {
	defer resetHttpGet()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", "test_id",
		"service.name", "test_name",
		"service.version", "test_version",
		"trace.id", "test_trace_id",
		"span.id", "test_span_id",
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
	mockHttpGet := new(MockedHttpGet)
	httpGet = mockHttpGet.Get
	mockInscriptionPage := `
	<!doctype html>
	<html lang=en>
	  <head>
		<meta charset=utf-8>
		<meta name=format-detection content='telephone=no'>
		<meta name=viewport content='width=device-width,initial-scale=1.0'>
		<meta property=og:title content='Inscription 9553787'>
		<meta property=og:image content='https://ip-172-31-9-253/static/favicon.png'>
		<meta property=twitter:card content=summary>
		<title>Inscription 9553787</title>
		<link rel=alternate href=/feed.xml type=application/rss+xml title='Inscription RSS Feed'>
		<link rel=stylesheet href=/static/index.css>
		<link rel=stylesheet href=/static/modern-normalize.css>
		<script src=/static/index.js defer></script>
	  </head>
	  <body>
	  <header>
		<nav>
		  <a href=/>Ordinals<sup>alpha</sup></a>
		  <a href=https://docs.ordinals.com/>Handbook</a>
		  <a href=https://github.com/casey/ord>Wallet</a>
		  <a href=/clock>Clock</a>
		  <form action=/search method=get>
			<input type=text autocapitalize=off autocomplete=off autocorrect=off name=query spellcheck=false>
			<input type=submit value='&#9906'>
		  </form>
		</nav>
	  </header>
	  <main>
	<h1>Inscription 9553787</h1>
	<div class=inscription>
	<a class=prev href=/inscription/763acce82848621b47fbd58991f56da1f7a405ae028e7516a61fc6eefff237aai0>❮</a>
	<iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0></iframe>
	<a class=next href=/inscription/03c2f430a9f4513b9509e578f04808ca2618cf5a244e15c57c8da7e52ddfdeeci0>❯</a>
	</div>
	<dl>
	  <dt>id</dt>
	  <dd class=monospace>3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0</dd>
	  <dt>address</dt>
	  <dd class=monospace>bc1putjs4fvkp3uaq6nhph7h2e7pmpwduq6zrxkt5kyyxe4rn47yrwzqup8lfu</dd>
	  <dt>output value</dt>
	  <dd>10000</dd>
	  <dt>preview</dt>
	  <dd><a href=/preview/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0>link</a></dd>
	  <dt>content</dt>
	  <dd><a href=/content/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0>link</a></dd>
	  <dt>content length</dt>
	  <dd>3440 bytes</dd>
	  <dt>content type</dt>
	  <dd>image/webp</dd>
	  <dt>timestamp</dt>
	  <dd><time>2023-05-28 03:28:17 UTC</time></dd>
	  <dt>genesis height</dt>
	  <dd><a href=/block/791720>791720</a></dd>
	  <dt>genesis fee</dt>
	  <dd>21000</dd>
	  <dt>genesis transaction</dt>
	  <dd><a class=monospace href=/tx/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab>3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab</a></dd>
	  <dt>location</dt>
	  <dd class=monospace>3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0:0</dd>
	  <dt>output</dt>
	  <dd><a class=monospace href=/output/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0>3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0</a></dd>
	  <dt>offset</dt>
	  <dd>0</dd>
	</dl>
	
	  </main>
	  </body>
	</html>	`
	mockHttpGet.On("Get", "http://localhost:8080/inscription/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(mockInscriptionPage)),
	}, nil)
	mockInscriptionContent := `
	{
		"p": "brc-721",
		"op": "deploy",
		"tick": "sato",
		"max": "10000",
		"meta": {
			"name": "Satoshi",
			"description": "The ChatGPT 09/May/2023 Financial institutions on the precipice as three banks collapse in 2023.",
			"image": "data:image/svg+xml;base64,PHN2ZyB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgaGVpZ2h0PSI2NCIgd2lkdGg9IjY0IiB2ZXJzaW9uPSIxLjEiIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyI+CjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAuMDA2MzA4NzYsLTAuMDAzMDE5ODQpIj4KPHBhdGggZmlsbD0iI2Y3OTMxYSIgZD0ibTYzLjAzMywzOS43NDRjLTQuMjc0LDE3LjE0My0yMS42MzcsMjcuNTc2LTM4Ljc4MiwyMy4zMDEtMTcuMTM4LTQuMjc0LTI3LjU3MS0yMS42MzgtMjMuMjk1LTM4Ljc4LDQuMjcyLTE3LjE0NSwyMS42MzUtMjcuNTc5LDM4Ljc3NS0yMy4zMDUsMTcuMTQ0LDQuMjc0LDI3LjU3NiwyMS42NCwyMy4zMDIsMzguNzg0eiIvPgo8cGF0aCBmaWxsPSIjRkZGIiBkPSJtNDYuMTAzLDI3LjQ0NGMwLjYzNy00LjI1OC0yLjYwNS02LjU0Ny03LjAzOC04LjA3NGwxLjQzOC01Ljc2OC0zLjUxMS0wLjg3NS0xLjQsNS42MTZjLTAuOTIzLTAuMjMtMS44NzEtMC40NDctMi44MTMtMC42NjJsMS40MS01LjY1My0zLjUwOS0wLjg3NS0xLjQzOSw1Ljc2NmMtMC43NjQtMC4xNzQtMS41MTQtMC4zNDYtMi4yNDItMC41MjdsMC4wMDQtMC4wMTgtNC44NDItMS4yMDktMC45MzQsMy43NXMyLjYwNSwwLjU5NywyLjU1LDAuNjM0YzEuNDIyLDAuMzU1LDEuNjc5LDEuMjk2LDEuNjM2LDIuMDQybC0xLjYzOCw2LjU3MWMwLjA5OCwwLjAyNSwwLjIyNSwwLjA2MSwwLjM2NSwwLjExNy0wLjExNy0wLjAyOS0wLjI0Mi0wLjA2MS0wLjM3MS0wLjA5MmwtMi4yOTYsOS4yMDVjLTAuMTc0LDAuNDMyLTAuNjE1LDEuMDgtMS42MDksMC44MzQsMC4wMzUsMC4wNTEtMi41NTItMC42MzctMi41NTItMC42MzdsLTEuNzQzLDQuMDE5LDQuNTY5LDEuMTM5YzAuODUsMC4yMTMsMS42ODMsMC40MzYsMi41MDMsMC42NDZsLTEuNDUzLDUuODM0LDMuNTA3LDAuODc1LDEuNDM5LTUuNzcyYzAuOTU4LDAuMjYsMS44ODgsMC41LDIuNzk4LDAuNzI2bC0xLjQzNCw1Ljc0NSwzLjUxMSwwLjg3NSwxLjQ1My01LjgyM2M1Ljk4NywxLjEzMywxMC40ODksMC42NzYsMTIuMzg0LTQuNzM5LDEuNTI3LTQuMzYtMC4wNzYtNi44NzUtMy4yMjYtOC41MTUsMi4yOTQtMC41MjksNC4wMjItMi4wMzgsNC40ODMtNS4xNTV6bS04LjAyMiwxMS4yNDljLTEuMDg1LDQuMzYtOC40MjYsMi4wMDMtMTAuODA2LDEuNDEybDEuOTI4LTcuNzI5YzIuMzgsMC41OTQsMTAuMDEyLDEuNzcsOC44NzgsNi4zMTd6bTEuMDg2LTExLjMxMmMtMC45OSwzLjk2Ni03LjEsMS45NTEtOS4wODIsMS40NTdsMS43NDgtNy4wMWMxLjk4MiwwLjQ5NCw4LjM2NSwxLjQxNiw3LjMzNCw1LjU1M3oiLz4KPC9nPgo8L3N2Zz4="
		}
	}`
	mockHttpGet.On("Get", "http://localhost:8080/content/3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(mockInscriptionContent)),
	}, nil)

	info, err := worker.parseInscriptionInfo("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96abi0")
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(9553787), info["inscription_id"].(int64), "inscription_id")
	r.Equal("bc1putjs4fvkp3uaq6nhph7h2e7pmpwduq6zrxkt5kyyxe4rn47yrwzqup8lfu", info["address"].(string), "address")
	r.Equal(uint64(10000), info["output_value"].(uint64), "output_value")
	r.Equal(time.Date(2023, 5, 28, 3, 28, 17, 0, time.UTC), info["timestamp"].(time.Time), "timestamp")
	r.Equal(uint64(791720), info["genesis_height"].(uint64), "genesis_height")
	r.Equal(uint64(21000), info["genesis_fee"].(uint64), "genesis_fee")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab", info["genesis_transaction"].(string), "genesis_transaction")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0:0", info["location"].(string), "location")
	r.Equal("3501f4fa1f754e5e7c58a153efbcec92a93b2ff9721a215bec6cdb9dd48d96ab:0", info["output"].(string), "output")
	r.Equal(uint64(0), info["offset"].(uint64), "offset")

	r.Equal("brc-721-deploy", info["content_parser"], "content_parser")
	t.Logf("%T", info["content"])
	o, ok := info["content"].(*parser.BRC721Deploy)
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
	defer resetHttpGet()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", "test_id",
		"service.name", "test_name",
		"service.version", "test_version",
		"trace.id", "test_trace_id",
		"span.id", "test_span_id",
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
	mockHttpGet := new(MockedHttpGet)
	httpGet = mockHttpGet.Get
	mockInscriptionPage := `
<!doctype html>
<html lang=en>
  <head>
    <meta charset=utf-8>
    <meta name=format-detection content='telephone=no'>
    <meta name=viewport content='width=device-width,initial-scale=1.0'>
    <meta property=og:title content='Inscription 4986756'>
    <meta property=og:image content='https://ip-172-31-9-253/static/favicon.png'>
    <meta property=twitter:card content=summary>
    <title>Inscription 4986756</title>
    <link rel=alternate href=/feed.xml type=application/rss+xml title='Inscription RSS Feed'>
    <link rel=stylesheet href=/static/index.css>
    <link rel=stylesheet href=/static/modern-normalize.css>
    <script src=/static/index.js defer></script>
  </head>
  <body>
  <header>
    <nav>
      <a href=/>Ordinals<sup>alpha</sup></a>
      <a href=https://docs.ordinals.com/>Handbook</a>
      <a href=https://github.com/casey/ord>Wallet</a>
      <a href=/clock>Clock</a>
      <form action=/search method=get>
        <input type=text autocapitalize=off autocomplete=off autocorrect=off name=query spellcheck=false>
        <input type=submit value='&#9906'>
      </form>
    </nav>
  </header>
  <main>
<h1>Inscription 4986756</h1>
<div class=inscription>
<a class=prev href=/inscription/97b0af07bfc4b60e907d4df6278d7051e6f3ba54acbe87736a8a3d4c079e934fi0>❮</a>
<iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0></iframe>
<a class=next href=/inscription/62633833e6a3991ccc0f9e5559b5c6ecd36125e2fe22bd50232775169b09fe69i0>❯</a>
</div>
<dl>
  <dt>id</dt>
  <dd class=monospace>8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0</dd>
  <dt>address</dt>
  <dd class=monospace>bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf</dd>
  <dt>output value</dt>
  <dd>10000</dd>
  <dt>preview</dt>
  <dd><a href=/preview/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0>link</a></dd>
  <dt>content</dt>
  <dd><a href=/content/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0>link</a></dd>
  <dt>content length</dt>
  <dd>46 bytes</dd>
  <dt>content type</dt>
  <dd>application/json</dd>
  <dt>timestamp</dt>
  <dd><time>2023-05-09 07:22:28 UTC</time></dd>
  <dt>genesis height</dt>
  <dd><a href=/block/788906>788906</a></dd>
  <dt>genesis fee</dt>
  <dd>30870</dd>
  <dt>genesis transaction</dt>
  <dd><a class=monospace href=/tx/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564>8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564</a></dd>
  <dt>location</dt>
  <dd class=monospace>8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0:0</dd>
  <dt>output</dt>
  <dd><a class=monospace href=/output/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0>8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0</a></dd>
  <dt>offset</dt>
  <dd>0</dd>
</dl>

  </main>
  </body>
</html>
`
	mockHttpGet.On("Get", "http://localhost:8080/inscription/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(mockInscriptionPage)),
	}, nil)
	mockInscriptionContent := `{"p": "brc-721", "op": "mint", "tick": "sato"}`
	mockHttpGet.On("Get", "http://localhost:8080/content/8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(mockInscriptionContent)),
	}, nil)

	info, err := worker.parseInscriptionInfo("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564i0")
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(4986756), info["inscription_id"].(int64), "inscription_id")
	r.Equal("bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf", info["address"].(string), "address")
	r.Equal(uint64(10000), info["output_value"].(uint64), "output_value")
	r.Equal(time.Date(2023, 5, 9, 7, 22, 28, 0, time.UTC), info["timestamp"].(time.Time), "timestamp")
	r.Equal(uint64(788906), info["genesis_height"].(uint64), "genesis_height")
	r.Equal(uint64(30870), info["genesis_fee"].(uint64), "genesis_fee")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564", info["genesis_transaction"].(string), "genesis_transaction")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0:0", info["location"].(string), "location")
	r.Equal("8417b71ef08dd2c824b8c5712f558228fcb67032a6589185d12b67768c319564:0", info["output"].(string), "output")
	r.Equal(uint64(0), info["offset"].(uint64), "offset")

	r.Equal("brc-721-mint", info["content_parser"], "content_parser")
	t.Logf("%T", info["content"])
	o, ok := info["content"].(*parser.BRC721Mint)
	r.True(ok, "content")
	r.Equal("brc-721", o.P, "content.P")
	r.Equal("mint", o.Op, "content.Op")
	r.Equal("sato", o.Tick, "content.Tick")
}

func TestWorkerParseBRC721DeployWithMeta(t *testing.T) {
	defer resetHttpGet()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", "test_id",
		"service.name", "test_name",
		"service.version", "test_version",
		"trace.id", "test_trace_id",
		"span.id", "test_span_id",
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
	mockHttpGet := new(MockedHttpGet)
	httpGet = mockHttpGet.Get
	mockInscriptionPage := `
	
<!doctype html>
<html lang=en>
  <head>
    <meta charset=utf-8>
    <meta name=format-detection content='telephone=no'>
    <meta name=viewport content='width=device-width,initial-scale=1.0'>
    <meta property=og:title content='Inscription 5183645'>
    <meta property=og:image content='https://ordinals.net/static/favicon.png'>
    <meta property=twitter:card content=summary>
    <title>Inscription 5183645</title>
    <link rel=alternate href=/feed.xml type=application/rss+xml title='Inscription RSS Feed'>
    <link rel=stylesheet href=/static/index.css>
    <link rel=stylesheet href=/static/modern-normalize.css>
    <script src=/static/index.js defer></script>
  </head>
  <body>
  <header>
    <nav>
      <a href=/>Ordinals<sup>alpha</sup></a>
      <a href=https://docs.ordinals.com/>Handbook</a>
      <a href=https://github.com/casey/ord>Wallet</a>
      <a href=/clock>Clock</a>
      <a href=/rare.txt>rare.txt</a>
      <form action=/search method=get>
        <input type=text autocapitalize=off autocomplete=off autocorrect=off name=query spellcheck=false>
        <input type=submit value='&#9906'>
      </form>
    </nav>
  </header>
  <main>
<h1>Inscription 5183645</h1>
<div class=inscription>
<a class=prev href=/inscription/423992a9468b935e2e04234c0c5232f3d8f7acd0e1d867e9797cee3941bdc702i0>❮</a>
<iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0></iframe>
<a class=next href=/inscription/1540f3e1654e42fad2c79bd556aa9c4d57af0edad4a3b13fd63f91b76d89751ci0>❯</a>
</div>
<dl>
  <dt>id</dt>
  <dd class=monospace>e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0</dd>
  <dt>address</dt>
  <dd class=monospace>bc1p0rwqj2dwtsurejsazqx9fp62fk0esytrlwe7atftadenv369v0ksnkngg3</dd>
  <dt>output value</dt>
  <dd>546</dd>
  <dt>sat</dt>
  <dd><a href=/sat/1925232705830827>1925232705830827</a></dd>
  <dt>preview</dt>
  <dd><a href=/preview/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0>link</a></dd>
  <dt>content</dt>
  <dd><a href=/content/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0>link</a></dd>
  <dt>content length</dt>
  <dd>362 bytes</dd>
  <dt>content type</dt>
  <dd>text/plain;charset=utf-8</dd>
  <dt>timestamp</dt>
  <dd><time>2023-05-10 04:12:43 UTC</time></dd>
  <dt>genesis height</dt>
  <dd><a href=/block/789018>789018</a></dd>
  <dt>genesis fee</dt>
  <dd>59736</dd>
  <dt>genesis transaction</dt>
  <dd><a class=monospace href=/tx/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913>e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913</a></dd>
  <dt>location</dt>
  <dd class=monospace>e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0:0</dd>
  <dt>output</dt>
  <dd><a class=monospace href=/output/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0>e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0</a></dd>
  <dt>offset</dt>
  <dd>0</dd>
</dl>

  </main>
  </body>
</html>
`
	mockHttpGet.On("Get", "http://localhost:8080/inscription/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(mockInscriptionPage)),
	}, nil)
	mockInscriptionContent := `
	{
		"p": "brc-721",
		"op": "deploy",
		"tick": "BNFT",
		"max": "200913",
		"meta": {
			"name": "Ordinals",
			"description": "Taking the birth date of Bitcoin as the Genesis Yuan, brc721 will open a new round of blockchain legends.", 
			"image": "https://www.bitcoin.com/images/uploads/get-started-what-is-bitcoin-lg@2x.png"
		}
	}`
	mockHttpGet.On("Get", "http://localhost:8080/content/e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(mockInscriptionContent)),
	}, nil)

	info, err := worker.parseInscriptionInfo("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913i0")
	r := require.New(t)
	r.NoError(err)
	r.Equal(int64(5183645), info["inscription_id"].(int64), "inscription_id")
	r.Equal("bc1p0rwqj2dwtsurejsazqx9fp62fk0esytrlwe7atftadenv369v0ksnkngg3", info["address"].(string), "address")
	r.Equal(uint64(546), info["output_value"].(uint64), "output_value")
	r.Equal(time.Date(2023, 5, 10, 4, 12, 43, 0, time.UTC), info["timestamp"].(time.Time), "timestamp")
	r.Equal(uint64(789018), info["genesis_height"].(uint64), "genesis_height")
	r.Equal(uint64(59736), info["genesis_fee"].(uint64), "genesis_fee")
	r.Equal("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913", info["genesis_transaction"].(string), "genesis_transaction")
	r.Equal("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0:0", info["location"].(string), "location")
	r.Equal("e0c705195fc8c14379994501b63e6b6a371e6b3d8d2fac653540e9e01415d913:0", info["output"].(string), "output")
	r.Equal(uint64(0), info["offset"].(uint64), "offset")

	r.Equal("brc-721-deploy", info["content_parser"], "content_parser")
	t.Logf("%T", info["content"])
	o, ok := info["content"].(*parser.BRC721Deploy)
	r.True(ok, "content")
	r.Equal("brc-721", o.P, "content.P")
	r.Equal("deploy", o.Op, "content.Op")
	r.Equal("BNFT", o.Tick, "content.Tick")
	r.Equal("200913", o.Max, "content.Max")
	r.Equal("Ordinals", o.Meta.Name, "content.Meta.Name")
	r.Equal("Taking the birth date of Bitcoin as the Genesis Yuan, brc721 will open a new round of blockchain legends.", o.Meta.Description, "content.Meta.Description")
	r.Nil(o.BaseURI, "content.BaseURI")
}
