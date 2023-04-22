package quran

import (
	"encoding/xml"
)

// Quran represents the whole scripture
type Quran struct {
	XMLName xml.Name `xml:"quran"`
	Suras   []Sura   `xml:"sura"`
}
