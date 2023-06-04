package page

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	mockResults = make(map[string][]byte)
)

func mockHTTPGet(url string) (*http.Response, error) {
	body, ok := mockResults[url]
	if !ok {
		return nil, fmt.Errorf("url %s not found", url)
	}
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}, nil
}

func mockHTTPResult(url string, body []byte) {
	mockResults[url] = body
}

func resetHTTPGet() {
	mockResults = make(map[string][]byte)
}
