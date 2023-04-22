package quran

import "strings"

type Aya struct {
	Index int    `xml:"index,attr,str"`
	Text  string `xml:"text,attr"`
}


func (a *Aya) RemoveTashkeels()string{
	return removeTashkeels(a.Text)
}

func (a *Aya) Tokens() []Token {
	parts := strings.Fields(a.Text)
	tokens := make([]Token, 0)

	for i := 0; i < len(parts); i++ {
		// all annotation, skip
		if len(parts[i]) == 0 || isAllArabicAnnotation(parts[i]) {
			continue
		}

		tokens = append(tokens, Token{Text: parts[i]})
	}

	return tokens
}
