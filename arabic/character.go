package arabic

type Character interface {
	DiacriticCount() int
}

type character struct {
	buffer []rune
	offset int
}

// NewArabicCharacter returns an instance of arabic character
func NewArabicCharacter(buffer []rune, offset int) Character {
	return &character{
		buffer: buffer,
		offset: offset,
	}
}

func (c *character) DiacriticCount() int {

}
