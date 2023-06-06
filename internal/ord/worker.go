package ord

import (
	"fmt"
	"net/http"

	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/ord/page"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	httpGet = http.Get
)

type Worker struct {
	wid        int
	baseURL    string
	pageParser page.PageParser
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
		info = &page.Inscription{}
	}
	if info.UID == "" {
		info.UID = uid
	}
	// FIXME: inscription_id is not always available
	if info.ID == 0 {
		if err == nil {
			err = fmt.Errorf("failed to get inscription_id")
		}
		return &result{info: info, err: err}
	}
	w.logger.Debugf("[worker %d] parsed inscription %d", w.wid, info.ID)
	return &result{info: info, err: err}
}

func (w *Worker) parseInscriptionInfo(uid string) (*page.Inscription, error) {
	inscriptionPage := page.NewInscriptionPage(uid)
	w.logger.Debugf("[worker %d] fetching %s...", w.wid, inscriptionPage.URL())
	data, err := w.pageParser.Parse(inscriptionPage)
	if err != nil {
		return nil, err
	}
	inscription, ok := data.(*page.Inscription)
	if !ok {
		return nil, fmt.Errorf("invalid inscription page: %T", data)
	}
	content, err := w.parseContent(uid)
	if err != nil {
		return nil, err
	}
	inscription.Content = content
	return inscription, nil
}

func (w *Worker) parseContent(uid string) (*page.Content, error) {
	contentPage := page.NewContentPage(uid)
	w.logger.Debugf("[worker %d] fetching %s...", w.wid, contentPage.URL())
	data, err := w.pageParser.Parse(contentPage)
	if err != nil {
		return nil, err
	}
	content, ok := data.(*page.Content)
	if !ok {
		return nil, fmt.Errorf("invalid content page: %T", data)
	}
	return content, nil
}
