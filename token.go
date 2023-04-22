package quran

import (
	"fmt"
	"strings"
)

type Token struct {
	Text string
}

func (t *Token) Characters() []Character {
	characters := make([]Character, 0)

	// 0 expecting for letter
	// 1 expect diacritic
	var state int

	var c Character
	for _, v := range t.Text {
		if isArabicLetter(v) {
			if state == 0 {
				// found letter while expecting letter, store and expect tashkeel
				c.Letter = v
				state++
			} else {
				// found letter while expecting tashkeel, clear diacritic and mark as complete character
				characters = append(characters, Character{
					Letter:    c.Letter,
					Tashkeels: c.Tashkeels,
				})

				// reset parser
				state = 0
				c.Letter = 0
				c.Tashkeels = nil
			}
		} else if isArabicTashkeel(v) {
			// accumulate tashkeels
			if c.Tashkeels == nil {
				c.Tashkeels = []rune{v}
			} else {
				c.Tashkeels = append(c.Tashkeels, v)
			}
		} else if isArabicAnnotation(v) {
			// ignore annotation accumulate letter
			if c.Letter > 0 {
				characters = append(characters, Character{
					Letter:    c.Letter,
					Tashkeels: c.Tashkeels,
				})

				// reset parser
				state = 0
				c.Letter = 0
				c.Tashkeels = nil
			}
		} else {
			panic(fmt.Errorf("unexpected rune %04x", v))
		}
	}

	// handle leftovers
	if len(characters) == 0 && c.Letter > 0 {
		characters = append(characters, Character{
			Letter:    c.Letter,
			Tashkeels: c.Tashkeels,
		})
	}

	return characters
}

func (t *Token) Dump() string {
	values := make([]string, 0)
	for _, v := range t.Text {
		values = append(values, fmt.Sprintf("%04x", v))
	}

	return fmt.Sprintf("[ %s ]", strings.Join(values, ", "))
}
