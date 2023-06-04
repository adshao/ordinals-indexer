package page

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/adshao/ordinals-indexer/internal/conf"
)

func TestInscriptionsHomePage(t *testing.T) {
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
		<meta property=og:title content='Inscriptions'>
		<meta property=og:image content='https://ip-172-31-9-253/static/favicon.png'>
		<meta property=twitter:card content=summary>
		<title>Inscriptions</title>
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
	<h1>Inscriptions</h1>
	<div class=thumbnails>
	  <a href=/inscription/20b16bba26122ef01fbe800a32a4950712d36abe0a5d8e0cc50264117aaf4d3ei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/20b16bba26122ef01fbe800a32a4950712d36abe0a5d8e0cc50264117aaf4d3ei0></iframe></a>
	  <a href=/inscription/cfc486c622248cc4bb8b167b9dbdd0988b421358017d0b65aa9e0805ea38ec3di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/cfc486c622248cc4bb8b167b9dbdd0988b421358017d0b65aa9e0805ea38ec3di0></iframe></a>
	  <a href=/inscription/fabdf02be4a4521c5ba668128404f7198226ad3a23e4bbce55d0a12d2b5f1c3di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/fabdf02be4a4521c5ba668128404f7198226ad3a23e4bbce55d0a12d2b5f1c3di0></iframe></a>
	  <a href=/inscription/8fe841280d748d2c7eaab19c42f6507200d42e86418a34e4a093e457872abd3bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8fe841280d748d2c7eaab19c42f6507200d42e86418a34e4a093e457872abd3bi0></iframe></a>
	  <a href=/inscription/d71e34d0a02ccdf89efdcd0b7b44c4f22685ab88bdd0d8f3f1b5e4f3f96aba3bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d71e34d0a02ccdf89efdcd0b7b44c4f22685ab88bdd0d8f3f1b5e4f3f96aba3bi0></iframe></a>
	  <a href=/inscription/710af4ddb896a02eb2b55e65967a89411eb710cf2aeb0de1941b487c52ff693bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/710af4ddb896a02eb2b55e65967a89411eb710cf2aeb0de1941b487c52ff693bi0></iframe></a>
	  <a href=/inscription/5da2a6669453c912772858508647e015bd2517be4f66062b19e15816d7efe33ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/5da2a6669453c912772858508647e015bd2517be4f66062b19e15816d7efe33ai0></iframe></a>
	  <a href=/inscription/00d6853ea11a7edb9e9210ba07711d2efefe992aa72acc5114c9088e5c383439i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/00d6853ea11a7edb9e9210ba07711d2efefe992aa72acc5114c9088e5c383439i0></iframe></a>
	  <a href=/inscription/6f3c56d7c6eff2ae92f1d28626d74eea0c4e1a2a5c0b8f00fb9b830d7d1c1c39i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6f3c56d7c6eff2ae92f1d28626d74eea0c4e1a2a5c0b8f00fb9b830d7d1c1c39i0></iframe></a>
	  <a href=/inscription/8c7456a771d07519b078dd569347fd255a7b018ef60aa063fcee627cc5b7da37i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8c7456a771d07519b078dd569347fd255a7b018ef60aa063fcee627cc5b7da37i0></iframe></a>
	  <a href=/inscription/05a02423dc80aeb1f1dabc43f51cb99feabd5dd3faf60e50bb4de3812b958b36i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/05a02423dc80aeb1f1dabc43f51cb99feabd5dd3faf60e50bb4de3812b958b36i0></iframe></a>
	  <a href=/inscription/376bd5243405004534c3c4c5700f0ff82e099c2fcb839317c9d4f299421afe35i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/376bd5243405004534c3c4c5700f0ff82e099c2fcb839317c9d4f299421afe35i0></iframe></a>
	  <a href=/inscription/0342cd888e4e5321f3e072a29860cae44fd66079522b06f507ed69089578a635i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/0342cd888e4e5321f3e072a29860cae44fd66079522b06f507ed69089578a635i0></iframe></a>
	  <a href=/inscription/810809716213f440243180e4603d2663372e3b7aa04881220e3df110e20c0635i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/810809716213f440243180e4603d2663372e3b7aa04881220e3df110e20c0635i0></iframe></a>
	  <a href=/inscription/29a9cadd5fd4a9f025fcc7ede0f7de36d8966353d7df1a5e346446087ef2a734i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/29a9cadd5fd4a9f025fcc7ede0f7de36d8966353d7df1a5e346446087ef2a734i0></iframe></a>
	  <a href=/inscription/ccd7834079483c7194cd09a7daa07adc68a1366bd109548e671793f82119fa33i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/ccd7834079483c7194cd09a7daa07adc68a1366bd109548e671793f82119fa33i0></iframe></a>
	  <a href=/inscription/9eecdd7edbdc39c54880d092742984a24ece8f2e565b0281d5d89973e0315433i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9eecdd7edbdc39c54880d092742984a24ece8f2e565b0281d5d89973e0315433i0></iframe></a>
	  <a href=/inscription/d06cc86970ec39fffa5d3f1c7be050880864ac06c2a47c8b8b9cb63268f03e33i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d06cc86970ec39fffa5d3f1c7be050880864ac06c2a47c8b8b9cb63268f03e33i0></iframe></a>
	  <a href=/inscription/20294f402397897a883cd92f3bd98c40e9629991748c56902b1bfb0916abd632i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/20294f402397897a883cd92f3bd98c40e9629991748c56902b1bfb0916abd632i0></iframe></a>
	  <a href=/inscription/2cf73b16fcd83b10a4b4e425083f0c9a84211e9b1e343b3a0b5997f97bd29a32i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2cf73b16fcd83b10a4b4e425083f0c9a84211e9b1e343b3a0b5997f97bd29a32i0></iframe></a>
	  <a href=/inscription/19870d85fb0ba3c73c6f7bc51a22f8e67cdd21f3280bc7ad1e64f2c9ce547132i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/19870d85fb0ba3c73c6f7bc51a22f8e67cdd21f3280bc7ad1e64f2c9ce547132i0></iframe></a>
	  <a href=/inscription/76592adc781faf531c2b07f4ba6ee5a88f0757df3510222195d48ae0153b4231i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/76592adc781faf531c2b07f4ba6ee5a88f0757df3510222195d48ae0153b4231i0></iframe></a>
	  <a href=/inscription/71dbe7f6df28bbdbc2b06ea3692d2eb56916c876d63d97946ace725a9de49c2ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/71dbe7f6df28bbdbc2b06ea3692d2eb56916c876d63d97946ace725a9de49c2ai0></iframe></a>
	  <a href=/inscription/595a53e08fb0bd36cb15e7fa14d917d76a99d20118e90a2c0f0605633789592ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/595a53e08fb0bd36cb15e7fa14d917d76a99d20118e90a2c0f0605633789592ai0></iframe></a>
	  <a href=/inscription/b7c3a8f123c9a5ebf98e2da266c274285415f8954bf6f8fa3e75cc44c614f2d3i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/b7c3a8f123c9a5ebf98e2da266c274285415f8954bf6f8fa3e75cc44c614f2d3i0></iframe></a>
	  <a href=/inscription/cfa94d041b4da69ae607c063516071219bf4b6efe7f92e8baa1225c7dd70fc29i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/cfa94d041b4da69ae607c063516071219bf4b6efe7f92e8baa1225c7dd70fc29i0></iframe></a>
	  <a href=/inscription/67bd4c54482381fbef5a904cfbae37444375b6ea28f538cd125a1b913a07b829i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/67bd4c54482381fbef5a904cfbae37444375b6ea28f538cd125a1b913a07b829i0></iframe></a>
	  <a href=/inscription/c16b74809dbeb1181fda23529d84ca5b16c7320f9541b232af5c4f3bd8adf827i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c16b74809dbeb1181fda23529d84ca5b16c7320f9541b232af5c4f3bd8adf827i0></iframe></a>
	  <a href=/inscription/533c7b2e2ecfa3edf2337b87fa7cb54ace053def3486db50a309a5765451f627i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/533c7b2e2ecfa3edf2337b87fa7cb54ace053def3486db50a309a5765451f627i0></iframe></a>
	  <a href=/inscription/84e640f11c7a92df8d6d826028991fc7ea67699eeb66a06a5ef3e6c38ba1c627i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/84e640f11c7a92df8d6d826028991fc7ea67699eeb66a06a5ef3e6c38ba1c627i0></iframe></a>
	  <a href=/inscription/429f754d269fbf4bfc6f7ba8d56928ae3e6f53ba6f4c31aa66046c2272137027i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/429f754d269fbf4bfc6f7ba8d56928ae3e6f53ba6f4c31aa66046c2272137027i0></iframe></a>
	  <a href=/inscription/65b87076c8473bfb5570a916690ad51b444a889ce1fb425fdfea863b591e4327i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/65b87076c8473bfb5570a916690ad51b444a889ce1fb425fdfea863b591e4327i0></iframe></a>
	  <a href=/inscription/7d119f1adbf8f7e65e6545f6f1eda769c22b0601b953c0624864e50ba9981b27i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7d119f1adbf8f7e65e6545f6f1eda769c22b0601b953c0624864e50ba9981b27i0></iframe></a>
	  <a href=/inscription/36b9a1154b3d996e50dba0b24229cbb0277929881bdf6ebaf1d45e336f87f526i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/36b9a1154b3d996e50dba0b24229cbb0277929881bdf6ebaf1d45e336f87f526i0></iframe></a>
	  <a href=/inscription/0a128c0be419c7e18b86494cf13a12b77c2792daa01f9253011fb55d28ea2526i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/0a128c0be419c7e18b86494cf13a12b77c2792daa01f9253011fb55d28ea2526i0></iframe></a>
	  <a href=/inscription/1f31ca01a5ad0f9b1cbff235643be241eb90f65d924faa5e9bd1431775750e25i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/1f31ca01a5ad0f9b1cbff235643be241eb90f65d924faa5e9bd1431775750e25i0></iframe></a>
	  <a href=/inscription/5c0b124f0a3aa28be88d0bfe61d12ac1f75cc7cc26c75484791ed9b50d4c2c24i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/5c0b124f0a3aa28be88d0bfe61d12ac1f75cc7cc26c75484791ed9b50d4c2c24i0></iframe></a>
	  <a href=/inscription/da377b1b045d08ccb4af25d13be06eac1e441e054a13cf4da57d16216c7f7623i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/da377b1b045d08ccb4af25d13be06eac1e441e054a13cf4da57d16216c7f7623i0></iframe></a>
	  <a href=/inscription/4bcd91c485e0278b831da32fc6daa6dafcfcb3a8c164b14af6b71142b79ddf22i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/4bcd91c485e0278b831da32fc6daa6dafcfcb3a8c164b14af6b71142b79ddf22i0></iframe></a>
	  <a href=/inscription/bc4eed03a9b42bd15ec23bac7e507ecaff76c372d2e137a836b1e52c6d27c222i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/bc4eed03a9b42bd15ec23bac7e507ecaff76c372d2e137a836b1e52c6d27c222i0></iframe></a>
	  <a href=/inscription/a56477de1f4b3794cefaadc1392bfe9d9f0e4dba656fe71664b985f70ef94a22i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a56477de1f4b3794cefaadc1392bfe9d9f0e4dba656fe71664b985f70ef94a22i0></iframe></a>
	  <a href=/inscription/8d15709577ef0df3a524f7cdbed64c9279ee0dcc9259d8118fdfb79e8328f721i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8d15709577ef0df3a524f7cdbed64c9279ee0dcc9259d8118fdfb79e8328f721i0></iframe></a>
	  <a href=/inscription/7d71d754920042e78eddfb472be1e5c3b54ae37c882aa53ffa77811dfffb8f21i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7d71d754920042e78eddfb472be1e5c3b54ae37c882aa53ffa77811dfffb8f21i0></iframe></a>
	  <a href=/inscription/c9b45efd1c6107fce1df8f224435de9bb368b99988857c99abf59d93ac6d8420i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c9b45efd1c6107fce1df8f224435de9bb368b99988857c99abf59d93ac6d8420i0></iframe></a>
	  <a href=/inscription/e17f5d0a05cefee45107813b43d28b8116d86895207cbe5c0389a78780146220i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/e17f5d0a05cefee45107813b43d28b8116d86895207cbe5c0389a78780146220i0></iframe></a>
	  <a href=/inscription/310646181f73578cb94df723c1ebca550533589307898cd1dc0c1773381c0320i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/310646181f73578cb94df723c1ebca550533589307898cd1dc0c1773381c0320i0></iframe></a>
	  <a href=/inscription/b5be2664b3766b17e289c9f572bcc7154fe189af8b1578643e87f64ca037291ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/b5be2664b3766b17e289c9f572bcc7154fe189af8b1578643e87f64ca037291ci0></iframe></a>
	  <a href=/inscription/4b090ae2b6dba6596510203fddb3e23675c993d5c96938c2c2efea4dc1fdec1bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/4b090ae2b6dba6596510203fddb3e23675c993d5c96938c2c2efea4dc1fdec1bi0></iframe></a>
	  <a href=/inscription/a5034e9bd16e9b215a0c8ebf2f703c04650961f148e0a9dc898c2423ddd0401bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a5034e9bd16e9b215a0c8ebf2f703c04650961f148e0a9dc898c2423ddd0401bi0></iframe></a>
	  <a href=/inscription/f76353ec2a85447f50e886138c0bf630afde26d31e540b03028cd644e3a8661ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f76353ec2a85447f50e886138c0bf630afde26d31e540b03028cd644e3a8661ai0></iframe></a>
	  <a href=/inscription/3fd6f1a66f2ceb2a3b2cf61c8edb8f3076dd001da7540c7f4dcd1aebdceaaf19i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/3fd6f1a66f2ceb2a3b2cf61c8edb8f3076dd001da7540c7f4dcd1aebdceaaf19i0></iframe></a>
	  <a href=/inscription/9be801aad35b885e1cfc4e2075561286a69db1d1260576589f3ea4bcebe90819i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9be801aad35b885e1cfc4e2075561286a69db1d1260576589f3ea4bcebe90819i0></iframe></a>
	  <a href=/inscription/ca748f58d6a51588171890fa5a4c8ce7473d6a4d6cc85b5e69857e39ef263c17i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/ca748f58d6a51588171890fa5a4c8ce7473d6a4d6cc85b5e69857e39ef263c17i0></iframe></a>
	  <a href=/inscription/ae914cb11bf4d49dc438c50b578ba4fdc5280d32a00a0f9f688cf40d72ebe316i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/ae914cb11bf4d49dc438c50b578ba4fdc5280d32a00a0f9f688cf40d72ebe316i0></iframe></a>
	  <a href=/inscription/2d5cc5de380e4758816bc385285e9b56ec812e3957463d1ac1ae2aaa31735516i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2d5cc5de380e4758816bc385285e9b56ec812e3957463d1ac1ae2aaa31735516i0></iframe></a>
	  <a href=/inscription/eb5d7acb8f476a4b28bb3a271f4f238ee1b3f92e7ae0f43ae331795f5af6e814i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/eb5d7acb8f476a4b28bb3a271f4f238ee1b3f92e7ae0f43ae331795f5af6e814i0></iframe></a>
	  <a href=/inscription/d78c2fc6d66e5706deb001ac60bed9bb03341b77ee19a57c9cd80300ed4cbc14i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d78c2fc6d66e5706deb001ac60bed9bb03341b77ee19a57c9cd80300ed4cbc14i0></iframe></a>
	  <a href=/inscription/2ed9c39f1e6b0f8878567728cb44f409ade4816192b126dbd13d4a9e93f4d412i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2ed9c39f1e6b0f8878567728cb44f409ade4816192b126dbd13d4a9e93f4d412i0></iframe></a>
	  <a href=/inscription/7240cbceb722ee16ad7e7e45b3e5b79f8d8c2c3e4bc4802eef7d24485305d112i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7240cbceb722ee16ad7e7e45b3e5b79f8d8c2c3e4bc4802eef7d24485305d112i0></iframe></a>
	  <a href=/inscription/e72a9c4ef68d1ca8968749d55627b141325427924a29d4995577e20651ddb611i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/e72a9c4ef68d1ca8968749d55627b141325427924a29d4995577e20651ddb611i0></iframe></a>
	  <a href=/inscription/c93e80166ae3ac64c3e6ca61be1f286eb147fd4a89423b2899b61645fcb3a511i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c93e80166ae3ac64c3e6ca61be1f286eb147fd4a89423b2899b61645fcb3a511i0></iframe></a>
	  <a href=/inscription/430fea0a71852fe6e5e3255e401d37262e32b8bf055eadaa26a252485b328d11i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/430fea0a71852fe6e5e3255e401d37262e32b8bf055eadaa26a252485b328d11i0></iframe></a>
	  <a href=/inscription/2f301a0a9988d050f67fbcd462eb762e24a5e4e332d6fa4005a76b2fd496e210i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2f301a0a9988d050f67fbcd462eb762e24a5e4e332d6fa4005a76b2fd496e210i0></iframe></a>
	  <a href=/inscription/4ab4d072c85c6b45650f39fccb162df2c508a15610ef6724aebd9fe9d681eadai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/4ab4d072c85c6b45650f39fccb162df2c508a15610ef6724aebd9fe9d681eadai0></iframe></a>
	  <a href=/inscription/f62e9ab09a87eda4c6a455fb4a76209a226024dbaa180d38dc2e9e11072e730fi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f62e9ab09a87eda4c6a455fb4a76209a226024dbaa180d38dc2e9e11072e730fi0></iframe></a>
	  <a href=/inscription/d78f1be0ac6343b46208261aa683a9b287ceb096b9a67152652485e3fd1905d1i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d78f1be0ac6343b46208261aa683a9b287ceb096b9a67152652485e3fd1905d1i0></iframe></a>
	  <a href=/inscription/40730c5055ec072c77cd4daf66dcbe46522df2a060b1ca407293b4574132b80ei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/40730c5055ec072c77cd4daf66dcbe46522df2a060b1ca407293b4574132b80ei0></iframe></a>
	  <a href=/inscription/e7eda296eaf3e47c8f5be316d747616e9dca30fbece8f2cbe7bc62f8257f180ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/e7eda296eaf3e47c8f5be316d747616e9dca30fbece8f2cbe7bc62f8257f180ci0></iframe></a>
	  <a href=/inscription/080aa97fb6cbc32e19190ce8e220a41bb68d8d78092a54912ca29e4c0a07980bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/080aa97fb6cbc32e19190ce8e220a41bb68d8d78092a54912ca29e4c0a07980bi0></iframe></a>
	  <a href=/inscription/3c94580e054b3b49220db3cbe55b9508f51ce64a13b0ed254e9208a73b205d0bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/3c94580e054b3b49220db3cbe55b9508f51ce64a13b0ed254e9208a73b205d0bi0></iframe></a>
	  <a href=/inscription/f5c0195f24face2bb42402e5941633b4f006a5091106024f63000bb2977c050ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f5c0195f24face2bb42402e5941633b4f006a5091106024f63000bb2977c050ai0></iframe></a>
	  <a href=/inscription/53311f655414d9fdf55aed0eccd6ca276f0294b0df0e5128d8107f74c4850a15i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/53311f655414d9fdf55aed0eccd6ca276f0294b0df0e5128d8107f74c4850a15i0></iframe></a>
	  <a href=/inscription/528552069dcc1fac9c4044ac957b6d1ccbf1ca14d37fd2b44940d7ad70e99a09i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/528552069dcc1fac9c4044ac957b6d1ccbf1ca14d37fd2b44940d7ad70e99a09i0></iframe></a>
	  <a href=/inscription/6b64bf829bb91b18c41fcbd06468a9c15415b220e24ebdd977e0ffe183da5509i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6b64bf829bb91b18c41fcbd06468a9c15415b220e24ebdd977e0ffe183da5509i0></iframe></a>
	  <a href=/inscription/abaef5f4d6bfa03bb108829e0b4f02cc79490dc6bc2647209cdab89db76a0769i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/abaef5f4d6bfa03bb108829e0b4f02cc79490dc6bc2647209cdab89db76a0769i0></iframe></a>
	  <a href=/inscription/85f1f842c7addd91b893fe47f699ac059c33b29df39f3b3f82dc631c8f067f08i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/85f1f842c7addd91b893fe47f699ac059c33b29df39f3b3f82dc631c8f067f08i0></iframe></a>
	  <a href=/inscription/c0d2e01adef37e45fa80fa6c3010d55427dfb0bfa2c7ca41f46ec96f605fd507i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c0d2e01adef37e45fa80fa6c3010d55427dfb0bfa2c7ca41f46ec96f605fd507i0></iframe></a>
	  <a href=/inscription/4df9d969b5980195e7c4e360047fa3ef7fea685a067dad019cd6f5878a57ba07i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/4df9d969b5980195e7c4e360047fa3ef7fea685a067dad019cd6f5878a57ba07i0></iframe></a>
	  <a href=/inscription/99c58e5d006a458a3c2c85197c0375a785f20d411824949f992ac5e7341bee43i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/99c58e5d006a458a3c2c85197c0375a785f20d411824949f992ac5e7341bee43i0></iframe></a>
	  <a href=/inscription/20300e4a9c09389c750df3f702cea19bffa982400ee75a53ae9342fc7e08cf06i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/20300e4a9c09389c750df3f702cea19bffa982400ee75a53ae9342fc7e08cf06i0></iframe></a>
	  <a href=/inscription/04e07fa7aedbe76135c891455d2f62891396b3619939e019a647b3c364040206i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/04e07fa7aedbe76135c891455d2f62891396b3619939e019a647b3c364040206i0></iframe></a>
	  <a href=/inscription/376b714e998bc0ad174b6d3e9110a06864f5358689ae9a23c30de5507f45d504i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/376b714e998bc0ad174b6d3e9110a06864f5358689ae9a23c30de5507f45d504i0></iframe></a>
	  <a href=/inscription/82000c63cf30a7ef22848e2e9581f9a069f9861e3ec4b17f2b66d58cc1c86104i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/82000c63cf30a7ef22848e2e9581f9a069f9861e3ec4b17f2b66d58cc1c86104i0></iframe></a>
	  <a href=/inscription/8fee732e38865c1bec71b780488a912e4fa02badbcb5b9e16ab8059c85f04c04i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8fee732e38865c1bec71b780488a912e4fa02badbcb5b9e16ab8059c85f04c04i0></iframe></a>
	  <a href=/inscription/7d9ccb7928bdc1cd56d92fb570094b6dcfcd97d3eae19116a105a983f46a37d9i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7d9ccb7928bdc1cd56d92fb570094b6dcfcd97d3eae19116a105a983f46a37d9i0></iframe></a>
	  <a href=/inscription/e495de1841d7e5a85a1056f9b8d476836e87960d5ef3b248595b166435fc1a03i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/e495de1841d7e5a85a1056f9b8d476836e87960d5ef3b248595b166435fc1a03i0></iframe></a>
	  <a href=/inscription/6753c0262ec9bfb92404dbbe447b335ac17b4a53628d14222338fbf5c8d7c302i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6753c0262ec9bfb92404dbbe447b335ac17b4a53628d14222338fbf5c8d7c302i0></iframe></a>
	  <a href=/inscription/f7f4b2b51ad1f2f1672199106036d4f9953e6381d27d9e68edeb92865cb87702i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f7f4b2b51ad1f2f1672199106036d4f9953e6381d27d9e68edeb92865cb87702i0></iframe></a>
	  <a href=/inscription/a2128de6897eb0b1a06de84201cb9b45b84e209fabfdc19ec78bf9ed52236a01i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a2128de6897eb0b1a06de84201cb9b45b84e209fabfdc19ec78bf9ed52236a01i0></iframe></a>
	  <a href=/inscription/eaf6e493d2f450d275385c69af05e10efd9e4fff304e36b74c6a13208d331e01i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/eaf6e493d2f450d275385c69af05e10efd9e4fff304e36b74c6a13208d331e01i0></iframe></a>
	  <a href=/inscription/8f004a661b17752efcc4033bd6bb1e18c52c27a0870ddfb1f5262666124264f1i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8f004a661b17752efcc4033bd6bb1e18c52c27a0870ddfb1f5262666124264f1i0></iframe></a>
	  <a href=/inscription/a793e2e1f230469deb014433cad8afc50340705806854cf3e63ab3b039c788ebi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a793e2e1f230469deb014433cad8afc50340705806854cf3e63ab3b039c788ebi0></iframe></a>
	  <a href=/inscription/d0a8478bebc508a8e42e01c0eccfa2781d4f9e5a86acf86ee240d49c43e6e666i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d0a8478bebc508a8e42e01c0eccfa2781d4f9e5a86acf86ee240d49c43e6e666i0></iframe></a>
	  <a href=/inscription/f8104820699da7712103e27507cc62267f5cda9b8cbdf977317c0a7b0e397837i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f8104820699da7712103e27507cc62267f5cda9b8cbdf977317c0a7b0e397837i0></iframe></a>
	  <a href=/inscription/7288de0f0ba5e599206e6d77c3b4bb2f9f3172c443ce9edcd0455411cba653abi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7288de0f0ba5e599206e6d77c3b4bb2f9f3172c443ce9edcd0455411cba653abi0></iframe></a>
	  <a href=/inscription/a4b2eeb520a383aafdf00e3b67127736df5e1ec41cd6ffe3ee716e096f6863c7i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a4b2eeb520a383aafdf00e3b67127736df5e1ec41cd6ffe3ee716e096f6863c7i0></iframe></a>
	  <a href=/inscription/dab9a69ea184d50438d317a35b6e990a76c2ffcf942882af2d0a11eaf9c398c8i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/dab9a69ea184d50438d317a35b6e990a76c2ffcf942882af2d0a11eaf9c398c8i0></iframe></a>
	  <a href=/inscription/7265a0a85413ffb0a5eed30627af504f6078c1ced8a52a3f1cc068fb2176bae9i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7265a0a85413ffb0a5eed30627af504f6078c1ced8a52a3f1cc068fb2176bae9i0></iframe></a>
	  <a href=/inscription/50adb8845f991b9f0a0059f553644a231a5c1fb75796859ad4aed04023593f48i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/50adb8845f991b9f0a0059f553644a231a5c1fb75796859ad4aed04023593f48i0></iframe></a>
	  <a href=/inscription/889626980b33da6a96e8db061e5007133cc781df5d3c8f9670d2ca5ff76dabd5i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/889626980b33da6a96e8db061e5007133cc781df5d3c8f9670d2ca5ff76dabd5i0></iframe></a>
	</div>
	<div class=center>
	<a class=prev href=/inscriptions/10400370>prev</a>
	next
	</div>
	
	  </main>
	  </body>
	</html>`)
	mockHTTPResult("http://localhost:8080/inscriptions", homePageBody)

	inscriptionsPage := NewInscriptionsPage()
	data, err := parser.Parse(inscriptionsPage)

	r := require.New(t)
	r.Nil(err)
	inscriptions, ok := data.(*Inscriptions)
	r.True(ok)
	r.Equal(100, len(inscriptions.UIDs))
	r.Equal("20b16bba26122ef01fbe800a32a4950712d36abe0a5d8e0cc50264117aaf4d3ei0", inscriptions.UIDs[0])
	r.Equal("cfc486c622248cc4bb8b167b9dbdd0988b421358017d0b65aa9e0805ea38ec3di0", inscriptions.UIDs[1])
	r.Equal("889626980b33da6a96e8db061e5007133cc781df5d3c8f9670d2ca5ff76dabd5i0", inscriptions.UIDs[99])
	r.Nil(inscriptions.NextID)
}

func TestInscriptionsPage(t *testing.T) {
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
	pageBody := []byte(`

	<!doctype html>
	<html lang=en>
	  <head>
		<meta charset=utf-8>
		<meta name=format-detection content='telephone=no'>
		<meta name=viewport content='width=device-width,initial-scale=1.0'>
		<meta property=og:title content='Inscriptions'>
		<meta property=og:image content='https://ip-172-31-9-253/static/favicon.png'>
		<meta property=twitter:card content=summary>
		<title>Inscriptions</title>
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
	<h1>Inscriptions</h1>
	<div class=thumbnails>
	  <a href=/inscription/018e5bd5eb839ad7804d63ed09338cc52576697db536eaf2c950a5912ce96e0di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/018e5bd5eb839ad7804d63ed09338cc52576697db536eaf2c950a5912ce96e0di0></iframe></a>
	  <a href=/inscription/01d96aeeb0d790bb5396af4e9176590b214cb9508eb90f5aab595b918ad78bf8i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/01d96aeeb0d790bb5396af4e9176590b214cb9508eb90f5aab595b918ad78bf8i0></iframe></a>
	  <a href=/inscription/531a952bbaac1e04253848ce300816d34448e96712259a49d5dbbc5958566859i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/531a952bbaac1e04253848ce300816d34448e96712259a49d5dbbc5958566859i0></iframe></a>
	  <a href=/inscription/6435f34e3680b44affec353a9d5bfb5b00ed6aade3d3f40b274c5861e8eed126i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6435f34e3680b44affec353a9d5bfb5b00ed6aade3d3f40b274c5861e8eed126i0></iframe></a>
	  <a href=/inscription/a8168a99ed983bbf922f93c1cec5b2b433f7d1e7e06170c2a96ab0295e4c404ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a8168a99ed983bbf922f93c1cec5b2b433f7d1e7e06170c2a96ab0295e4c404ai0></iframe></a>
	  <a href=/inscription/2916f4820f6e7616c64cd15fefdeff02e6bea9ca40232616c281870498e3b693i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2916f4820f6e7616c64cd15fefdeff02e6bea9ca40232616c281870498e3b693i0></iframe></a>
	  <a href=/inscription/d66f6f50dbae3658f5dca0a6db7c97229b1c6ba8069ee69a82a1d2918e930a63i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d66f6f50dbae3658f5dca0a6db7c97229b1c6ba8069ee69a82a1d2918e930a63i0></iframe></a>
	  <a href=/inscription/2c5d0ac30bffcbb104ee7c0830d8092b65951c3b96bdc30b00676841aaf0fd7ei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2c5d0ac30bffcbb104ee7c0830d8092b65951c3b96bdc30b00676841aaf0fd7ei0></iframe></a>
	  <a href=/inscription/bf92c3a2dc27eb5d509832b088df5f98fe9d1c557ac02196b435e426a0b57d0ei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/bf92c3a2dc27eb5d509832b088df5f98fe9d1c557ac02196b435e426a0b57d0ei0></iframe></a>
	  <a href=/inscription/77c24b9a5662d7582ddc69b30babdb3d3da789080a6d1ca87b88fca494b9e43di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/77c24b9a5662d7582ddc69b30babdb3d3da789080a6d1ca87b88fca494b9e43di0></iframe></a>
	  <a href=/inscription/51300cca90c1858985c177e399ee69cd37562757bf47fc031fee9c5e67450204i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/51300cca90c1858985c177e399ee69cd37562757bf47fc031fee9c5e67450204i0></iframe></a>
	  <a href=/inscription/eb5dd7f875d1aa7da66199bea245b4156f9888cf12a1fae167fd96af5fca54d3i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/eb5dd7f875d1aa7da66199bea245b4156f9888cf12a1fae167fd96af5fca54d3i0></iframe></a>
	  <a href=/inscription/ffafc53e3964ae28ccbc1a02cb37035512d1c6e1f2da67ae28fa1dab1fa54db6i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/ffafc53e3964ae28ccbc1a02cb37035512d1c6e1f2da67ae28fa1dab1fa54db6i0></iframe></a>
	  <a href=/inscription/3fd8f15d6a8021413ad66fab0cf8bc040441bc99191769104968ab2e701eb68ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/3fd8f15d6a8021413ad66fab0cf8bc040441bc99191769104968ab2e701eb68ci0></iframe></a>
	  <a href=/inscription/69f2c75f2dddf2902aeeb1113302e5e0310a025fb6073a38911543a14dac3474i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/69f2c75f2dddf2902aeeb1113302e5e0310a025fb6073a38911543a14dac3474i0></iframe></a>
	  <a href=/inscription/6302e1b53030c3edfdccb246863dfabade59d2d1a28ef490c675127a37d83c68i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6302e1b53030c3edfdccb246863dfabade59d2d1a28ef490c675127a37d83c68i0></iframe></a>
	  <a href=/inscription/28680e6ac59a1f4fde8278a2471328f18c0f90c686330bf1e0ea474b8931dcebi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/28680e6ac59a1f4fde8278a2471328f18c0f90c686330bf1e0ea474b8931dcebi0></iframe></a>
	  <a href=/inscription/46a2641e55b8078dd57e922ab4d13f4a601999f2b7ee4c1994bc1cb27bc856ddi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/46a2641e55b8078dd57e922ab4d13f4a601999f2b7ee4c1994bc1cb27bc856ddi0></iframe></a>
	  <a href=/inscription/5b43fdebc86167faa1e497a97fd445f191a21e6ef25b9533446d12dfe200059bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/5b43fdebc86167faa1e497a97fd445f191a21e6ef25b9533446d12dfe200059bi0></iframe></a>
	  <a href=/inscription/0a81df9f39fb111a7a9207e4ffc9f7588f3cb332ad893708cba98ee34beee709i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/0a81df9f39fb111a7a9207e4ffc9f7588f3cb332ad893708cba98ee34beee709i0></iframe></a>
	  <a href=/inscription/c784211890d3b0dec8991eb3cc2a73319d888c5ed3da7934e38a78b7184902e3i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c784211890d3b0dec8991eb3cc2a73319d888c5ed3da7934e38a78b7184902e3i0></iframe></a>
	  <a href=/inscription/d330c53937062af3c38a0c60bf85b372682e26ab8e417e3d6ae47f0478ae14eei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d330c53937062af3c38a0c60bf85b372682e26ab8e417e3d6ae47f0478ae14eei0></iframe></a>
	  <a href=/inscription/8df1e86d5186e23704cf5bdabe681455df5f0df169bb208ee40f6ce446ccea25i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8df1e86d5186e23704cf5bdabe681455df5f0df169bb208ee40f6ce446ccea25i0></iframe></a>
	  <a href=/inscription/9f1297802e33b8971dbd049da649c25dad0948be5f62d7fc38516ca650441b97i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9f1297802e33b8971dbd049da649c25dad0948be5f62d7fc38516ca650441b97i0></iframe></a>
	  <a href=/inscription/171803a4d2add0b3a95f5e4107ced58379cc1f2cff6776697d067647db592851i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/171803a4d2add0b3a95f5e4107ced58379cc1f2cff6776697d067647db592851i0></iframe></a>
	  <a href=/inscription/cae7b88e5e04628aaa796bf8187955b306ae380d097b1d2d4d24ff9f2f27324ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/cae7b88e5e04628aaa796bf8187955b306ae380d097b1d2d4d24ff9f2f27324ci0></iframe></a>
	  <a href=/inscription/9c9ec808cfdc03a81eb966e714e1fb6536e836d0e1b03a7f29b00a6941804338i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9c9ec808cfdc03a81eb966e714e1fb6536e836d0e1b03a7f29b00a6941804338i0></iframe></a>
	  <a href=/inscription/2f6296cfb211299f474395166f596ab1738ed1c7afcf9b6190c1d930b141fe15i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2f6296cfb211299f474395166f596ab1738ed1c7afcf9b6190c1d930b141fe15i0></iframe></a>
	  <a href=/inscription/b2f16aef117e154c79a6c1c33f700293da132a0b081979bce89b40f9736719edi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/b2f16aef117e154c79a6c1c33f700293da132a0b081979bce89b40f9736719edi0></iframe></a>
	  <a href=/inscription/f1944fefdedc92d2cd99974cd725e1ec3d92043a2a8e12b2bff645e49ce28fe7i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f1944fefdedc92d2cd99974cd725e1ec3d92043a2a8e12b2bff645e49ce28fe7i0></iframe></a>
	  <a href=/inscription/13c42b0e35496311912520df618d04d8316aebee9cd5f18fa107866db61761dai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/13c42b0e35496311912520df618d04d8316aebee9cd5f18fa107866db61761dai0></iframe></a>
	  <a href=/inscription/38452851efeba810382363aa41d7479e72a56443041b956beeeac5a8d63411d9i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/38452851efeba810382363aa41d7479e72a56443041b956beeeac5a8d63411d9i0></iframe></a>
	  <a href=/inscription/6d5df64ff642c28a1fa1fbfc6b30708543c6c55a74223bf6e3a59d6e60a21fc4i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6d5df64ff642c28a1fa1fbfc6b30708543c6c55a74223bf6e3a59d6e60a21fc4i0></iframe></a>
	  <a href=/inscription/50ae5eb72b94cc74a338a9580f43c6974e0a4905a939e1eb8008a819186716bei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/50ae5eb72b94cc74a338a9580f43c6974e0a4905a939e1eb8008a819186716bei0></iframe></a>
	  <a href=/inscription/5673e3f6657c3638c2ede531fc0c60c95d71cf7e301236a397fa00d4b5e721b2i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/5673e3f6657c3638c2ede531fc0c60c95d71cf7e301236a397fa00d4b5e721b2i0></iframe></a>
	  <a href=/inscription/d6ff2a99c09c3ea9f3d8585d9ff4835da3c0604ee569cd105cae50158ddaf2a2i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d6ff2a99c09c3ea9f3d8585d9ff4835da3c0604ee569cd105cae50158ddaf2a2i0></iframe></a>
	  <a href=/inscription/10657de81698b786615eae967468ba8ee25db5fab238f5450ab4ff23f933f085i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/10657de81698b786615eae967468ba8ee25db5fab238f5450ab4ff23f933f085i0></iframe></a>
	  <a href=/inscription/acd3cdb90f12af6dbaa7575168c7424ee0b86cd8ca408f5e03a51ed7f7730776i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/acd3cdb90f12af6dbaa7575168c7424ee0b86cd8ca408f5e03a51ed7f7730776i0></iframe></a>
	  <a href=/inscription/d337cba1b34a87de14d3e21a0512c31fcbfe31ec02d530591d5a1ddcc24b9267i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d337cba1b34a87de14d3e21a0512c31fcbfe31ec02d530591d5a1ddcc24b9267i0></iframe></a>
	  <a href=/inscription/aca86f6dc564e8f9ffdd0cbaec9d6f17fe44a82504f57e75050a7f44c0d64363i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/aca86f6dc564e8f9ffdd0cbaec9d6f17fe44a82504f57e75050a7f44c0d64363i0></iframe></a>
	  <a href=/inscription/2e7cc2380db129e516371edd63ea7d0987a6710a5876c2652e0e009ff9f8375di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2e7cc2380db129e516371edd63ea7d0987a6710a5876c2652e0e009ff9f8375di0></iframe></a>
	  <a href=/inscription/6706ef0a54f221f1859157f13e5c14838960baefe625f07d35e3cad31ced8b57i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6706ef0a54f221f1859157f13e5c14838960baefe625f07d35e3cad31ced8b57i0></iframe></a>
	  <a href=/inscription/2311c20f8abbec874a7efea90fab4765200c712324c34047f12830720c030d54i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2311c20f8abbec874a7efea90fab4765200c712324c34047f12830720c030d54i0></iframe></a>
	  <a href=/inscription/dc640eeff8c9a552856f3e0676db693fa8a0ceae2eb505417f32ad26150cce53i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/dc640eeff8c9a552856f3e0676db693fa8a0ceae2eb505417f32ad26150cce53i0></iframe></a>
	  <a href=/inscription/6f15427ff6c8145e4f1793ac961fd1a5e349b292fa4b3555cd932b1568143c51i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/6f15427ff6c8145e4f1793ac961fd1a5e349b292fa4b3555cd932b1568143c51i0></iframe></a>
	  <a href=/inscription/ce6920504665b615b016bb499473dfcb2d804ca89f0c63afa551d7f51f01404di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/ce6920504665b615b016bb499473dfcb2d804ca89f0c63afa551d7f51f01404di0></iframe></a>
	  <a href=/inscription/c3197ae09c9dd62559fe96686440a5e96920c454f3953ce71f6a85df2d779123i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c3197ae09c9dd62559fe96686440a5e96920c454f3953ce71f6a85df2d779123i0></iframe></a>
	  <a href=/inscription/b9d42f7c7d9cfed1129db751407a4eaa8626a83996ac38c9d80d3a38add19a01i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/b9d42f7c7d9cfed1129db751407a4eaa8626a83996ac38c9d80d3a38add19a01i0></iframe></a>
	  <a href=/inscription/c0f1f77b86fb0ca17ff0ddd6f8f11c2c49d675a2cb36f04d19382d56e026b6f1i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c0f1f77b86fb0ca17ff0ddd6f8f11c2c49d675a2cb36f04d19382d56e026b6f1i0></iframe></a>
	  <a href=/inscription/04b70aa72305f41b37b4d288b580f89f76e35d73db24f523dde2615c6f8644abi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/04b70aa72305f41b37b4d288b580f89f76e35d73db24f523dde2615c6f8644abi0></iframe></a>
	  <a href=/inscription/f297c4465b52400785241c14091c4717fa6bfa441531eac1c574cadff0af2ba4i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f297c4465b52400785241c14091c4717fa6bfa441531eac1c574cadff0af2ba4i0></iframe></a>
	  <a href=/inscription/d0c5f6a9671b296e97eec5bab271ecce07c806331dea9fe3f068c89d4ed60d64i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d0c5f6a9671b296e97eec5bab271ecce07c806331dea9fe3f068c89d4ed60d64i0></iframe></a>
	  <a href=/inscription/fe13c0c67231f78f317e04c476a0e3bf95325280d9cdd31905ba2227b131ed5ei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/fe13c0c67231f78f317e04c476a0e3bf95325280d9cdd31905ba2227b131ed5ei0></iframe></a>
	  <a href=/inscription/cba3376f17308f415daff57e2a8fd10a8445f5bbda555cbfa8b2aa3977738245i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/cba3376f17308f415daff57e2a8fd10a8445f5bbda555cbfa8b2aa3977738245i0></iframe></a>
	  <a href=/inscription/8e8d82c5508af946ea2c8f152a3d694c98ad0184fffd5a3a465e6edc3978ab32i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8e8d82c5508af946ea2c8f152a3d694c98ad0184fffd5a3a465e6edc3978ab32i0></iframe></a>
	  <a href=/inscription/4e84ef7e5b8474be2280f777ea172008e23ccfafb13d61a02b80e46817aa7bcei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/4e84ef7e5b8474be2280f777ea172008e23ccfafb13d61a02b80e46817aa7bcei0></iframe></a>
	  <a href=/inscription/d71b47de798064efa2b6b6031ca685139c9547455a11cebfa49de9affa369edei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d71b47de798064efa2b6b6031ca685139c9547455a11cebfa49de9affa369edei0></iframe></a>
	  <a href=/inscription/1d6352f61cf3516651b9bd2fd5abc0dcebb1468fae58e42696725156842133d8i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/1d6352f61cf3516651b9bd2fd5abc0dcebb1468fae58e42696725156842133d8i0></iframe></a>
	  <a href=/inscription/a712bcbbfe0ce3e825167c2bc28d9e27f7e82254f2c8a338deb9d2e008170879i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a712bcbbfe0ce3e825167c2bc28d9e27f7e82254f2c8a338deb9d2e008170879i0></iframe></a>
	  <a href=/inscription/2f2b50235651c129db14e2f8eabefe40b145a6b9e56fcb75783adac761e661aci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2f2b50235651c129db14e2f8eabefe40b145a6b9e56fcb75783adac761e661aci0></iframe></a>
	  <a href=/inscription/10a88195cafc934641c5078d988754f3bef28bf8b318f612c8af3569c5ee4f3ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/10a88195cafc934641c5078d988754f3bef28bf8b318f612c8af3569c5ee4f3ci0></iframe></a>
	  <a href=/inscription/1c9320478fbf2d9d350b9e745c719f38685c4e08f35fee7ef2abacc00a2d101ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/1c9320478fbf2d9d350b9e745c719f38685c4e08f35fee7ef2abacc00a2d101ci0></iframe></a>
	  <a href=/inscription/9d9ddae4caf06b12c7bd7da6905a529d2d53d285bb9d7cca89e7ad94f9b85a0ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9d9ddae4caf06b12c7bd7da6905a529d2d53d285bb9d7cca89e7ad94f9b85a0ai0></iframe></a>
	  <a href=/inscription/3bc39a2fdd9301acfad5cf52bedddc3763045d4df8d17f904ee428d943fe0854i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/3bc39a2fdd9301acfad5cf52bedddc3763045d4df8d17f904ee428d943fe0854i0></iframe></a>
	  <a href=/inscription/e9f6a8629644e26aa7c1ab4d2f92d83604a8fbb6aa827d648ecb6469e9766df3i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/e9f6a8629644e26aa7c1ab4d2f92d83604a8fbb6aa827d648ecb6469e9766df3i0></iframe></a>
	  <a href=/inscription/c7dd7e3f2b7c48ee2494e9d725cbce41992bc34a0e95c3438bb9baf5218caca8i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c7dd7e3f2b7c48ee2494e9d725cbce41992bc34a0e95c3438bb9baf5218caca8i0></iframe></a>
	  <a href=/inscription/82a42c3405d599b63197b3ebdcebb0af5832a47e04a7856b2a5c6fe83cd94ed0i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/82a42c3405d599b63197b3ebdcebb0af5832a47e04a7856b2a5c6fe83cd94ed0i0></iframe></a>
	  <a href=/inscription/d3b9b12abd2eca6be91d96a5a139af407a52d6c2b357da3de0825c4ecf2440bbi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d3b9b12abd2eca6be91d96a5a139af407a52d6c2b357da3de0825c4ecf2440bbi0></iframe></a>
	  <a href=/inscription/9d7d20bac95a6b420a02e4da713f899c0473db45a72844a7c6520ecaa69aa6a2i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9d7d20bac95a6b420a02e4da713f899c0473db45a72844a7c6520ecaa69aa6a2i0></iframe></a>
	  <a href=/inscription/565b4b3f80e9dcb8d29dbdca37c0f09139f4e13a717ffc15c66ccda4792753f6i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/565b4b3f80e9dcb8d29dbdca37c0f09139f4e13a717ffc15c66ccda4792753f6i0></iframe></a>
	  <a href=/inscription/a52fec547fe0f96a2380c8e95870a8d9f9a1ec52d09e6605593dd3499df641e4i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a52fec547fe0f96a2380c8e95870a8d9f9a1ec52d09e6605593dd3499df641e4i0></iframe></a>
	  <a href=/inscription/94cdca7ed84fd124523ebb5ce6e87415aee37e4772a045e8f4315c7aa3fdacd5i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/94cdca7ed84fd124523ebb5ce6e87415aee37e4772a045e8f4315c7aa3fdacd5i0></iframe></a>
	  <a href=/inscription/5509de903ca998960882c21918fb0e9fac30790195d4e86c4ae94dae95d258c4i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/5509de903ca998960882c21918fb0e9fac30790195d4e86c4ae94dae95d258c4i0></iframe></a>
	  <a href=/inscription/08336bcfdc925be766457546afd148c40c846a3bbfc7574e14ba4423c621adb7i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/08336bcfdc925be766457546afd148c40c846a3bbfc7574e14ba4423c621adb7i0></iframe></a>
	  <a href=/inscription/4a2e872efb55aa387f0054e2da9e1235feefddca7c2e5264c67071e1fa5aa1a4i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/4a2e872efb55aa387f0054e2da9e1235feefddca7c2e5264c67071e1fa5aa1a4i0></iframe></a>
	  <a href=/inscription/7dfb3eedb5ac26f66fd9c376953cb4e7f5f52bc986afedcb816985ab141f348ei0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7dfb3eedb5ac26f66fd9c376953cb4e7f5f52bc986afedcb816985ab141f348ei0></iframe></a>
	  <a href=/inscription/f88ae413a1aa97da18170952e4fd9182fd8ef919dc2065711ffbfd84378cda8di0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f88ae413a1aa97da18170952e4fd9182fd8ef919dc2065711ffbfd84378cda8di0></iframe></a>
	  <a href=/inscription/456a6d2c4fa5ee871b3a7d61b4a67319b8eca3d4f44d6cfd3c52dcceb8d7217ci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/456a6d2c4fa5ee871b3a7d61b4a67319b8eca3d4f44d6cfd3c52dcceb8d7217ci0></iframe></a>
	  <a href=/inscription/c7235e673f2d77afa30863734e2e45b2ca9acda22bf737f404ebf1b0cdad8061i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c7235e673f2d77afa30863734e2e45b2ca9acda22bf737f404ebf1b0cdad8061i0></iframe></a>
	  <a href=/inscription/2eecd5affa1c71d60c244e79a2b3f98676a592d215f0607cf58b91d3c6e2cf59i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/2eecd5affa1c71d60c244e79a2b3f98676a592d215f0607cf58b91d3c6e2cf59i0></iframe></a>
	  <a href=/inscription/51e7b5b79426027abae007b7cd25a3bf08972ea2a09a1cb86c94837a79358459i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/51e7b5b79426027abae007b7cd25a3bf08972ea2a09a1cb86c94837a79358459i0></iframe></a>
	  <a href=/inscription/67f8173727c1480e26af934b3a35156505eae21292fe2d6de71e8d8ecd30964bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/67f8173727c1480e26af934b3a35156505eae21292fe2d6de71e8d8ecd30964bi0></iframe></a>
	  <a href=/inscription/a1c144066585174da147502438a82f94c142390eadbaafc3eaabd18a159a8548i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/a1c144066585174da147502438a82f94c142390eadbaafc3eaabd18a159a8548i0></iframe></a>
	  <a href=/inscription/345110ea4bf0ea7c266a3eb4071f5b0964b534f44a6d96f97856598d0842c047i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/345110ea4bf0ea7c266a3eb4071f5b0964b534f44a6d96f97856598d0842c047i0></iframe></a>
	  <a href=/inscription/cbefc9bf84ff9c021c86b452ede5a8aa6c7df6987108b823e2baf40a50dd4839i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/cbefc9bf84ff9c021c86b452ede5a8aa6c7df6987108b823e2baf40a50dd4839i0></iframe></a>
	  <a href=/inscription/1eaadf51e20b0eaf5c34db3d97ee460f90c00208bb94a1bc004ac16c81e72327i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/1eaadf51e20b0eaf5c34db3d97ee460f90c00208bb94a1bc004ac16c81e72327i0></iframe></a>
	  <a href=/inscription/8cb30a6b04202b2c93d6aa4f30c87e2bb712ea1fad2999d4cf731ef2fbddc520i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/8cb30a6b04202b2c93d6aa4f30c87e2bb712ea1fad2999d4cf731ef2fbddc520i0></iframe></a>
	  <a href=/inscription/b66d148ec972a75567838dcd04e5fdb5e13b6cc465d08e4c5a3c0f947081091fi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/b66d148ec972a75567838dcd04e5fdb5e13b6cc465d08e4c5a3c0f947081091fi0></iframe></a>
	  <a href=/inscription/310fe88ac4758606cccdf023dbd4eca56ef19ab650ded1daf6c887ec4186061ai0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/310fe88ac4758606cccdf023dbd4eca56ef19ab650ded1daf6c887ec4186061ai0></iframe></a>
	  <a href=/inscription/9eadc337fbb3913c3effc40ca68a09e1a04dd406cdb2060b17f1b2f545cbc815i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9eadc337fbb3913c3effc40ca68a09e1a04dd406cdb2060b17f1b2f545cbc815i0></iframe></a>
	  <a href=/inscription/ee158458a03f17e8ce470674259e52bf4903fd0815ea617d28b07fb75a8b3f08i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/ee158458a03f17e8ce470674259e52bf4903fd0815ea617d28b07fb75a8b3f08i0></iframe></a>
	  <a href=/inscription/b67e1e95f684c12fe536bee47dbe97899906de2c63b56d4e20580718fcd39401i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/b67e1e95f684c12fe536bee47dbe97899906de2c63b56d4e20580718fcd39401i0></iframe></a>
	  <a href=/inscription/1e279ffe2871fa2b087bacdb338305780333502683e0ed7ba6ebaf3f36453b86i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/1e279ffe2871fa2b087bacdb338305780333502683e0ed7ba6ebaf3f36453b86i0></iframe></a>
	  <a href=/inscription/f93aca8fbb2a4a5fa131d90119ec82a75f22dee5aa091e2b964f5e08abac5480i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/f93aca8fbb2a4a5fa131d90119ec82a75f22dee5aa091e2b964f5e08abac5480i0></iframe></a>
	  <a href=/inscription/29ad489f14c3a5e0764effd97dde4c400c29e6f81bacb075778284b2167ac8cfi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/29ad489f14c3a5e0764effd97dde4c400c29e6f81bacb075778284b2167ac8cfi0></iframe></a>
	  <a href=/inscription/7eceae16aa1e819e2873efe42d081984d4eaf76d5dd98fa1a7de333bfe6d34d0i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/7eceae16aa1e819e2873efe42d081984d4eaf76d5dd98fa1a7de333bfe6d34d0i0></iframe></a>
	  <a href=/inscription/d08c6b4da3a5f2a24a1a93c140867cdbed35b3bc45de3c34a3c60de5e87db64bi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/d08c6b4da3a5f2a24a1a93c140867cdbed35b3bc45de3c34a3c60de5e87db64bi0></iframe></a>
	  <a href=/inscription/9471ef0ad655080402f5cf9179feb9a055625fca5b69c539a0b6821a005e150fi0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/9471ef0ad655080402f5cf9179feb9a055625fca5b69c539a0b6821a005e150fi0></iframe></a>
	  <a href=/inscription/c8ae123ba72523a5a9d9582a2817a8b88cfc2012f134daa8e2b8faf45c52b6dci0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/c8ae123ba72523a5a9d9582a2817a8b88cfc2012f134daa8e2b8faf45c52b6dci0></iframe></a>
	  <a href=/inscription/560e3f7d62c945a53dfa05e3410391b6d8c663b430dc0aa605edd7503227e2e1i0><iframe sandbox=allow-scripts scrolling=no loading=lazy src=/preview/560e3f7d62c945a53dfa05e3410391b6d8c663b430dc0aa605edd7503227e2e1i0></iframe></a>
	</div>
	<div class=center>
	<a class=prev href=/inscriptions/10400270>prev</a>
	<a class=next href=/inscriptions/10400470>next</a>
	</div>
	
	  </main>
	  </body>
	</html>	`)
	mockHTTPResult("http://localhost:8080/inscriptions/10400370", pageBody)

	inscriptionsPage := NewInscriptionsPage(10400370)
	data, err := parser.Parse(inscriptionsPage)

	r := require.New(t)
	r.Nil(err)
	inscriptions, ok := data.(*Inscriptions)
	r.True(ok)
	r.Equal(100, len(inscriptions.UIDs))
	r.Equal("018e5bd5eb839ad7804d63ed09338cc52576697db536eaf2c950a5912ce96e0di0", inscriptions.UIDs[0])
	r.Equal("01d96aeeb0d790bb5396af4e9176590b214cb9508eb90f5aab595b918ad78bf8i0", inscriptions.UIDs[1])
	r.Equal("560e3f7d62c945a53dfa05e3410391b6d8c663b430dc0aa605edd7503227e2e1i0", inscriptions.UIDs[99])
	r.Equal(int64(10400470), *inscriptions.NextID)
}
