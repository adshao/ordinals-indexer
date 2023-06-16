package page

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/wire"

	"github.com/adshao/ordinals-indexer/internal/conf"
)

var ProviderSet = wire.NewSet(NewPageParser)

func NewPageParser(c *conf.Ord) PageParser {
	return &pageParser{
		httpGet: http.Get,
		c:       c,
	}
}

type PageParser interface {
	Parse(Page) (interface{}, error)
}

type pageParser struct {
	httpGet func(string) (*http.Response, error)
	c       *conf.Ord
}

func (parser *pageParser) parsePageRaw(p Page) (io.Reader, error) {
	u := p.URL()
	if !strings.HasPrefix(u, "http") {
		u, _ = url.JoinPath(parser.c.Server.Addr, u)
	}
	resp, err := parser.httpGet(u)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (parser *pageParser) Parse(p Page) (interface{}, error) {
	r, err := parser.parsePageRaw(p)
	if err != nil {
		return nil, err
	}
	if rc, ok := r.(io.ReadCloser); ok {
		defer rc.Close()
	}
	return p.Parse(r)
}

type Page interface {
	URL() string
	// Parse parses the page.
	Parse(io.Reader) (interface{}, error)
}
