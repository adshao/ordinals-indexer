package page

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Inscriptions struct {
	UIDs   []string `json:"uids"`
	NextID *int64   `json:"next_id,omitempty"`
	PrevID *int64   `json:"prev_id,omitempty"`
}

type InscriptionsPage struct {
	ID *int64 `json:"id,omitempty"`
}

var (
	_ Page = (*InscriptionsPage)(nil)
)

func NewInscriptionsPage(id ...int64) *InscriptionsPage {
	if len(id) == 0 {
		return &InscriptionsPage{
			ID: nil,
		}
	}
	return &InscriptionsPage{
		ID: &id[0],
	}
}

func (p *InscriptionsPage) URL() string {
	if p.ID == nil {
		return "/inscriptions"
	}
	return fmt.Sprintf("/inscriptions/%d", *p.ID)
}

func (p *InscriptionsPage) Parse(r io.Reader) (interface{}, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	inscriptions := &Inscriptions{
		UIDs: make([]string, 0),
	}
	links := doc.Find("div.thumbnails a")
	links.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		uid := strings.Replace(href, "/inscription/", "", -1)
		if uid == "" {
			return
		}
		// save the inscription uids, so that we can process them in order
		inscriptions.UIDs = append(inscriptions.UIDs, uid)
	})

	nextLink := doc.Find("a.next")
	if nextLink.Length() > 0 {
		href, _ := nextLink.Attr("href")
		nextIDText := strings.Replace(href, "/inscriptions/", "", -1)
		nextID, err := strconv.ParseInt(nextIDText, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert nextID %s to int64: %v", nextIDText, err)
		}
		inscriptions.NextID = &nextID
	}
	prevLink := doc.Find("a.prev")
	if prevLink.Length() > 0 {
		href, _ := prevLink.Attr("href")
		prevIDText := strings.Replace(href, "/inscriptions/", "", -1)
		prevID, err := strconv.ParseInt(prevIDText, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert prevID %s to int64: %v", prevIDText, err)
		}
		inscriptions.PrevID = &prevID
	}
	return inscriptions, nil
}
