package page

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/adshao/ordinals-indexer/internal/conf"
)

func TestInscriptionPage(t *testing.T) {
	defer resetHTTPGet()

	c := &conf.Ord{
		Server: &conf.Ord_Server{
			Addr: "http://localhost:8080",
		},
	}
	parser := &pageParser{
		httpGet: mockHTTPGet,
		c:       c,
	}
	homePageBody := []byte(`
	<!doctype html>
	<html lang=en>
	  <head>
		<meta charset=utf-8>
		<meta name=format-detection content='telephone=no'>
		<meta name=viewport content='width=device-width,initial-scale=1.0'>
		<meta property=og:title content='Inscription 4984402'>
		<meta property=og:image content='https://ip-172-31-9-253/static/favicon.png'>
		<meta property=twitter:card content=summary>
		<title>Inscription 4984402</title>
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
	<h1>Inscription 4984402</h1>
	<div class=inscription>
	<a class=prev href=/inscription/9bb83fa001542416bdf1eaeed41699f619110e9b68fb25b5cd2628dfb328c063i0>❮</a>
	<iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0></iframe>
	<a class=next href=/inscription/e3678715396719368e039fa56a09aa77eb30a2ea525f5489779626e355a31b65i0>❯</a>
	</div>
	<dl>
	  <dt>id</dt>
	  <dd class=monospace>347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0</dd>
	  <dt>address</dt>
	  <dd class=monospace>bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf</dd>
	  <dt>output value</dt>
	  <dd>10000</dd>
	  <dt>preview</dt>
	  <dd><a href=/preview/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0>link</a></dd>
	  <dt>content</dt>
	  <dd><a href=/content/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0>link</a></dd>
	  <dt>content length</dt>
	  <dd>2167 bytes</dd>
	  <dt>content type</dt>
	  <dd>application/json</dd>
	  <dt>timestamp</dt>
	  <dd><time>2023-05-09 07:13:59 UTC</time></dd>
	  <dt>genesis height</dt>
	  <dd><a href=/block/788904>788904</a></dd>
	  <dt>genesis fee</dt>
	  <dd>143010</dd>
	  <dt>genesis transaction</dt>
	  <dd><a class=monospace href=/tx/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564>347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564</a></dd>
	  <dt>location</dt>
	  <dd class=monospace>347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0:0</dd>
	  <dt>output</dt>
	  <dd><a class=monospace href=/output/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0>347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0</a></dd>
	  <dt>offset</dt>
	  <dd>0</dd>
	</dl>
	
	  </main>
	  </body>
	</html>`)
	mockHTTPResult("http://localhost:8080/inscription/347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", homePageBody)

	inscriptionPage := NewInscriptionPage("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0")
	data, err := parser.Parse(inscriptionPage)

	r := require.New(t)
	r.Nil(err)
	inscription, ok := data.(*Inscription)
	r.True(ok)
	r.Equal(int64(4984402), inscription.ID)
	r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564i0", inscription.UID)
	r.Equal("bc1phwdjdq59tqlszsd4gljqqsgvrygpasre4dj4ant98wvc30lqgqzsxxgkvf", inscription.Address)
	r.Equal(uint64(10000), inscription.OutputValue)
	r.Equal(uint64(2167), inscription.ContentLength)
	r.Equal("application/json", inscription.ContentType)
	r.Equal("2023-05-09 07:13:59 +0000 UTC", inscription.Timestamp.String())
	r.Equal(uint64(788904), inscription.GenesisHeight)
	r.Equal(uint64(143010), inscription.GenesisFee)
	r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564", inscription.GenesisTx)
	r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0:0", inscription.Location)
	r.Equal("347052bed5c7929f5e5186188ee3abd571f9c1f619d6ac6238b96437b7b72564:0", inscription.Output)
	r.Equal(uint64(0), inscription.Offset)
}
