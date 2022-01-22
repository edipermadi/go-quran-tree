package orthography

// Verse represents quran verse
type Verse interface {
	Chapter() Chapter
	ChapterNumber() int
	VerseNumber() int
	Location() Location
	TokenCount() int
	Token(location Location) Token
	Tokens() []Token
}

type verse struct {
	location    Location
	chapter     Chapter
	verseNumber int
	tokens      []Token
}

// NewVerse return an instance of verse
func NewVerse(location Location, chapter Chapter, verseNumber int, tokens []Token) Verse {
	return &verse{
		location:    location,
		chapter:     chapter,
		verseNumber: verseNumber,
		tokens:      tokens,
	}
}

func (v *verse) Chapter() Chapter {
	return v.chapter
}

func (v *verse) ChapterNumber() int {
	return v.chapter.ChapterNumber()
}

func (v *verse) VerseNumber() int {
	return v.verseNumber
}

func (v *verse) Location() Location {
	return v.location
}

func (v *verse) TokenCount() int {
	return len(v.tokens)
}

func (v *verse) Token(location Location) Token {
	return v.tokens[location.TokenNumber()]
}

func (v *verse) Tokens() []Token {
	return v.tokens
}
