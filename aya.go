package quran

type Aya struct {
	Index int    `xml:"index,attr,str"`
	Text  string `xml:"text,attr"`
}
