package page

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Inscription struct {
	ID            int64     `json:"inscription_id"`
	UID           string    `json:"uid"`
	Address       string    `json:"address,omitempty"`
	OutputValue   uint64    `json:"output_value,omitempty"`
	ContentLength uint64    `json:"content_length,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	GenesisHeight uint64    `json:"genesis_height,omitempty"`
	GenesisFee    uint64    `json:"genesis_fee,omitempty"`
	GenesisTx     string    `json:"genesis_tx,omitempty"`
	Location      string    `json:"location,omitempty"`
	Output        string    `json:"output,omitempty"`
	Offset        uint64    `json:"offset,omitempty"`
}

type InscriptionPage struct {
	UID string
}

var (
	_ Page = (*InscriptionPage)(nil)
)

func NewInscriptionPage(uid string) *InscriptionPage {
	return &InscriptionPage{
		UID: uid,
	}
}

func (p *InscriptionPage) URL() string {
	return fmt.Sprintf("/inscription/%s", p.UID)
}

func (p *InscriptionPage) Parse(r io.Reader) (interface{}, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	inscription := new(Inscription)
	inscriptionIDText := doc.Find("h1").First().Text()
	inscriptionIDText = strings.Replace(inscriptionIDText, "Inscription ", "", -1)
	// convert inscriptionID string to int64
	inscriptionID, err := strconv.ParseInt(inscriptionIDText, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inscriptionID %s to int64: %v", inscriptionIDText, err)
	}
	inscription.ID = inscriptionID

	dtElements := doc.Find("dl dt")
	ddElements := doc.Find("dl dd")
	dtElements.Each(func(i int, dt *goquery.Selection) {
		key := dt.Text()
		dd := ddElements.Eq(i)
		value := dd.Text()
		if aTag := dd.Find("a"); aTag.Length() > 0 {
			value = aTag.Text()
		}
		key = strings.ToLower(key)
		switch key {
		case "id":
			inscription.UID = value
		case "address":
			inscription.Address = value
		case "output value":
			v, _ := strconv.ParseUint(value, 10, 64)
			inscription.OutputValue = v
		case "content length":
			// conver "3440 bytes" to 3440
			value = strings.Replace(value, " bytes", "", -1)
			v, _ := strconv.ParseUint(value, 10, 64)
			inscription.ContentLength = v
		case "content type":
			inscription.ContentType = value
		case "timestamp":
			// convert "2023-05-28 03:28:17 UTC" to time.Time
			v, _ := time.Parse("2006-01-02 15:04:05 UTC", value)
			inscription.Timestamp = v
		case "genesis height":
			v, _ := strconv.ParseUint(value, 10, 64)
			inscription.GenesisHeight = v
		case "genesis fee":
			v, _ := strconv.ParseUint(value, 10, 64)
			inscription.GenesisFee = v
		case "genesis transaction":
			inscription.GenesisTx = value
		case "location":
			inscription.Location = value
		case "output":
			inscription.Output = value
		case "offset":
			v, _ := strconv.ParseUint(value, 10, 64)
			inscription.Offset = v
		}
	})
	return inscription, nil
}
