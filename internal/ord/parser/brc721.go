package parser

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	NameBRC721Deploy = "brc-721-deploy"
	NameBRC721Mint   = "brc-721-mint"
	NameBRC721Update = "brc-721-update"
)

const (
	OpBRC721 = "brc-721"
)

var (
	_ Parser = (*BRC721DeployParser)(nil)
	_ Parser = (*BRC721MintParser)(nil)
	_ Parser = (*BRC721UpdateParser)(nil)

	_ Validator = (*BRC721Deploy)(nil)
	_ Validator = (*BRC721Mint)(nil)
	_ Validator = (*BRC721Update)(nil)
)

type BRC721Meta struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Image       string                   `json:"image"`
	Attributes  []map[string]interface{} `json:"attributes"`
}

type BRC721Deploy struct {
	P       string      `json:"p"`
	Op      string      `json:"op"`
	Tick    string      `json:"tick"`
	Max     string      `json:"max"`
	BaseURI *string     `json:"buri"`
	Meta    *BRC721Meta `json:"meta"`
}

func (m BRC721Deploy) Validate() bool {
	if m.P != OpBRC721 {
		return false
	}
	if m.Op != "deploy" {
		return false
	}
	if m.Tick == "" {
		return false
	}
	if m.Max == "" {
		return false
	}
	if m.BaseURI == nil && m.Meta == nil {
		return false
	}
	return true
}

type BRC721DeployParser struct {
}

func (p *BRC721DeployParser) Name() string {
	return NameBRC721Deploy
}

func (p *BRC721DeployParser) Parse(data []byte) (interface{}, bool, error) {
	var deploy BRC721Deploy
	err := json.Unmarshal(data, &deploy)
	if err != nil {
		return nil, false, err
	}
	return &deploy, deploy.Validate(), nil
}

type BRC721Mint struct {
	P    string `json:"p"`
	Op   string `json:"op"`
	Tick string `json:"tick"`
}

func (m BRC721Mint) Validate() bool {
	if m.P != OpBRC721 {
		return false
	}
	if m.Op != "mint" {
		return false
	}
	if m.Tick == "" {
		return false
	}
	return true
}

type BRC721MintParser struct {
}

func (p *BRC721MintParser) Name() string {
	return NameBRC721Mint
}

func (p *BRC721MintParser) Parse(data []byte) (interface{}, bool, error) {
	var mint BRC721Mint
	err := json.Unmarshal(data, &mint)
	if err != nil {
		return nil, false, err
	}
	return &mint, mint.Validate(), nil
}

type BRC721Update struct {
	P       string  `json:"p"`
	Op      string  `json:"op"`
	Tick    string  `json:"tick"`
	BaseURI *string `json:"buri"`
}

func (m BRC721Update) Validate() bool {
	if m.P != OpBRC721 {
		return false
	}
	if m.Op != "update" {
		return false
	}
	if m.Tick == "" {
		return false
	}
	if m.BaseURI == nil {
		return false
	}
	return true
}

type BRC721UpdateParser struct {
}

func (p *BRC721UpdateParser) Name() string {
	return NameBRC721Update
}

func (p *BRC721UpdateParser) Parse(data []byte) (interface{}, bool, error) {
	var update BRC721Update
	err := json.Unmarshal(data, &update)
	if err != nil {
		return nil, false, err
	}
	return &update, update.Validate(), nil
}

func init() {
	registerParser(&BRC721DeployParser{})
	registerParser(&BRC721MintParser{})
	registerParser(&BRC721UpdateParser{})
}
