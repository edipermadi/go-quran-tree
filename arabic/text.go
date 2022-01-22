package arabic

import "github.com/edipermadi/unicode"

type Text interface {
}

type text struct {
	txt     string
	buffers []rune
}

var table = map[rune]struct{}{
	unicode.ArabicLetterHamza:{},
	unicode.ArabicLetterAlifW
}

// NewText returns arabic text from string
func NewText(txt string) Text {
	for _, v := range txt {

	}

	return &text{txt: txt}
}
func (t *text) Length() int {

}

func (t *text) LetterCount() int {

}

func (t *text) Character(index int) rune {

}
