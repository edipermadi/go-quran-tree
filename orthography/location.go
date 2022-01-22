package orthography

// Location represents location of a word in quran
type Location interface {
	ChapterNumber() int
	VerseNumber() int
	TokenNumber() int
}

// NewLocation returns an instance of location
func NewLocation(chapter, verse, token int) Location {
	return location{chapter: chapter, verse: verse, token: token}
}

type location struct {
	chapter int
	verse   int
	token   int
}

func (l location) ChapterNumber() int {
	return l.chapter
}

func (l location) TokenNumber() int {
	return l.token
}

func (l location) VerseNumber() int {
	return l.verse
}
