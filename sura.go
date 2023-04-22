package quran

type Sura struct {
	Index int    `xml:"index,attr,str"`
	Name  string `xml:"name,attr"`
	Ayas  []Aya  `xml:"aya"`
}
