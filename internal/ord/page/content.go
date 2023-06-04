package page

import (
	"fmt"
	"io"

	"github.com/adshao/ordinals-indexer/internal/ord/parser"
)

type Content struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
}

type ContentPage struct {
	inscriptionUid string
}

var (
	_ Page = (*ContentPage)(nil)
)

func NewContentPage(inscriptionUid string) *ContentPage {
	return &ContentPage{
		inscriptionUid: inscriptionUid,
	}
}

func (p *ContentPage) URL() string {
	return fmt.Sprintf("/content/%s", p.inscriptionUid)
}

func (p *ContentPage) Parse(r io.Reader) (interface{}, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, p := range parser.ParserList() {
		data, valid, err := p.Parse(body)
		if err != nil {
			continue
		}
		if !valid {
			continue
		}
		return &Content{
			Data: data,
			Type: p.Name(),
		}, nil
	}
	return &Content{
		Data: body,
		Type: "raw",
	}, nil
}
