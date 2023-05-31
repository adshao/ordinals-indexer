package ord

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/ord/parser"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	httpGet = http.Get
)

type Worker struct {
	wid        int
	baseURL    string
	data       *data.Data
	uidChan    chan string
	resultChan chan (*result)
	stopC      chan struct{}
	logger     *log.Helper
}

func (w *Worker) Start() {
	for {
		select {
		case uid := <-w.uidChan:
			w.logger.Debugf("[worker %d]: processing inscription %s", w.wid, uid)
			w.resultChan <- w.processInscription(uid)
		case <-w.stopC:
			w.logger.Infof("[worker %d]: stopping", w.wid)
			return
		}
	}
}

func (w *Worker) processInscription(uid string) *result {
	info, err := w.parseInscriptionInfo(uid)
	if info == nil {
		info = make(map[string]interface{})
	}
	info["uid"] = uid
	// FIXME: inscription_id is not always available
	inscriptionID, ok := info["inscription_id"].(int64)
	if !ok {
		if err == nil {
			err = fmt.Errorf("failed to get inscription_id")
		}
		return &result{inscriptionUid: uid, inscriptionId: 0, info: info, err: err}
	}
	w.logger.Debugf("[worker %d] parsed inscription %d", w.wid, inscriptionID)
	return &result{inscriptionUid: uid, inscriptionId: inscriptionID, info: info, err: err}
}

func (w *Worker) parseInscriptionInfo(uid string) (map[string]interface{}, error) {
	inscriptionURL, _ := url.JoinPath(w.baseURL, "inscription", uid)
	w.logger.Debugf("[worker %d] fetching %s...", w.wid, inscriptionURL)
	resp, err := httpGet(inscriptionURL)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	details := make(map[string]interface{})
	inscriptionIDText := doc.Find("h1").First().Text()
	inscriptionIDText = strings.Replace(inscriptionIDText, "Inscription ", "", -1)
	// convert inscriptionID string to int64
	inscriptionID, err := strconv.ParseInt(inscriptionIDText, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inscriptionID %s to int64: %v", inscriptionIDText, err)
	}
	details["inscription_id"] = inscriptionID

	dtElements := doc.Find("dl dt")
	ddElements := doc.Find("dl dd")
	dtElements.Each(func(i int, dt *goquery.Selection) {
		key := dt.Text()
		dd := ddElements.Eq(i)
		value := dd.Text()
		if aTag := dd.Find("a"); aTag.Length() > 0 {
			value = aTag.Text()
		}
		key = strings.Replace(strings.ToLower(key), " ", "_", -1)
		switch key {
		case "output_value":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "content_length":
			// conver "3440 bytes" to 3440
			value = strings.Replace(value, " bytes", "", -1)
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "timestamp":
			// convert "2023-05-28 03:28:17 UTC" to time.Time
			v, _ := time.Parse("2006-01-02 15:04:05 UTC", value)
			details[key] = v
		case "genesis_height":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "genesis_fee":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "offset":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		default:
			details[key] = value
		}
	})

	err = w.parseContent(details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (w *Worker) parseContent(info map[string]interface{}) error {
	contentURL, _ := url.JoinPath(w.baseURL, "content", info["id"].(string))
	w.logger.Debugf("[worker %d] fetching %s...", w.wid, contentURL)
	resp, err := httpGet(contentURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var found bool
	for _, p := range parser.ParserList() {
		data, valid, err := p.Parse(body)
		if err != nil {
			continue
		}
		if !valid {
			continue
		}
		found = true
		info["content"] = data
		info["content_parser"] = p.Name()
		break
	}
	if !found {
		info["content"] = body
		info["content_parser"] = "raw"
	}
	return nil
}
