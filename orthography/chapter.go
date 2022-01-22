package orthography

import (
	"github.com/edipermadi/go-quran-tree/arabic"
	"sync"
)

type Chapter interface {
	Location() Location
	Bismillah() arabic.ArabicText
	ChapterNumber() int
	Name() arabic.ArabicText
	VerseCount() int
	TokenCount() int
	Verse(location Location) Verse
	Verses() []Verse
}

type chapter struct {
	location       Location
	chapterNumber  int
	name           arabic.ArabicText
	verses         []Verse
	tokenCount     int
	tokenCountOnce sync.Once
}

// NewChapter returns an instance of quran chapter
func NewChapter(location Location, chapterNumber int, name arabic.ArabicText, verses []Verse) Chapter {
	return &chapter{
		location:      location,
		chapterNumber: chapterNumber,
		name:          name,
		verses:        verses,
	}
}

func (c *chapter) Location() Location {
	return c.location
}

func (c *chapter) Bismillah() arabic.ArabicText {
	//
	return nil
}

func (c *chapter) ChapterNumber() int {
	return c.chapterNumber
}

func (c *chapter) Name() arabic.ArabicText {
	return c.name
}

func (c *chapter) TokenCount() int {
	c.tokenCountOnce.Do(func() {
		tokenCount := 0
		for _, v := range c.verses {
			tokenCount += v.TokenCount()
		}

		c.tokenCount = tokenCount
	})

	return c.tokenCount
}

func (c *chapter) VerseCount() int {
	return len(c.verses)
}

func (c *chapter) Verse(location Location) Verse {
	return c.verses[location.VerseNumber()]
}

func (c *chapter) Verses() []Verse {
	return c.verses
}
