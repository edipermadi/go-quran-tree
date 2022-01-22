package tree_test

import (
	"testing"

	tree "github.com/edipermadi/go-quran-tree"
)

func TestTanzilReader_ReadXML(t *testing.T) {
	reader := tree.TanzilReader{}
	reader.ReadXML("./res/quran-uthmani.xml")
}
