package orthography

import "sync"

// Document represents quran document
type Document interface {
	Name() string
	ChapterCount() int
	Chapter(location Location) Chapter
	Chapters() []Chapter
	VerseCount() int
	Verse(location Location) Verse
	Verses(location Location) []Verse
	TokenCount() int
	Token(location Location) Token
	Tokens(location Location) []Token
}

// NewDocument return an instance of quran Document
func NewDocument(chapters []Chapter) Document {
	return &document{chapters: chapters}
}

type document struct {
	chapters       []Chapter
	tokenCount     int
	verseCount     int
	tokenCountOnce sync.Once
	verseCountOnce sync.Once
}

func (d *document) Name() string {
	return "The Holy Quran"
}

func (d *document) ChapterCount() int {
	return ChapterCount
}

func (d *document) Chapter(location Location) Chapter {
	return d.chapters[location.ChapterNumber()]
}

func (d *document) Chapters() []Chapter {
	return d.chapters
}

func (d *document) VerseCount() int {
	d.verseCountOnce.Do(func() {
		verseCount := 0
		for _, v := range d.chapters {
			verseCount += v.VerseCount()
		}

		d.verseCount = verseCount
	})

	return d.verseCount
}

func (d *document) Verse(location Location) Verse {
	return d.Chapter(location).Verse(location)
}

func (d *document) Verses(location Location) []Verse {
	return d.Chapter(location).Verses()
}

func (d *document) TokenCount() int {
	d.tokenCountOnce.Do(func() {
		tokenCount := 0
		for _, v := range d.chapters {
			tokenCount += v.TokenCount()
		}

		d.tokenCount = tokenCount
	})

	return d.tokenCount
}

func (d *document) Token(location Location) Token {
	return d.Chapter(location).Verse(location).Token(location)
}

func (d *document) Tokens(location Location) []Token {
	return d.Chapter(location).Verse(location).Tokens()
}
