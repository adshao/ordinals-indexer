package parser

import (
	"sync"
)

var (
	lock    sync.Mutex
	parsers = make([]Parser, 0)
)

type Parser interface {
	Parse(content []byte) (interface{}, bool, error)
	Name() string
}

type Validator interface {
	Validate() bool
}

func ParserList() []Parser {
	return parsers
}

func registerParser(parser Parser) {
	lock.Lock()
	parsers = append(parsers, parser)
	lock.Unlock()
}
