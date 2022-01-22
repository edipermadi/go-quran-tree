package orthography

type Token interface {
	Location() Location
	Chapter() Chapter
	Verse() Verse
	ChapterNumber() int
	VerseNumber() int
	TokenNumber() int
}

type token struct {
	location    Location
	chapter     Chapter
	verse       Verse
	tokenNumber int
}

func NewToken(location location, chapter Chapter, verse Verse, tokenNumber int) Token {
	return &token{
		location:    location,
		chapter:     chapter,
		verse:       verse,
		tokenNumber: tokenNumber,
	}
}

func (t *token) Chapter() Chapter {
	return t.chapter
}

func (t *token) ChapterNumber() int {
	return t.chapter.ChapterNumber()
}

func (t *token) Verse() Verse {
	return t.verse
}

func (t *token) VerseNumber() int {
	return t.verse.VerseNumber()
}

func (t *token) TokenNumber() int {
	return t.tokenNumber
}

func (t *token) Location() Location {
	return t.location
}
