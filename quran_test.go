package quran_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/edipermadi/go-quran-tree"
	"github.com/stretchr/testify/require"
)

func TestDocument_ParseTanzilXML(t *testing.T) {
	file, err := os.Open("quran-uthmani.xml")
	require.NoError(t, err)
	defer func() { _ = file.Close() }()

	// decode
	q := &quran.Quran{}
	require.NoError(t, xml.NewDecoder(file).Decode(q))

	// check first total sura, first sura, first aya
	require.NotEmpty(t, q.Suras)
	require.Equal(t, 114, len(q.Suras))
	require.Equal(t, 1, q.Suras[0].Index)
	require.Equal(t, "الفاتحة", q.Suras[0].Name)
	require.Equal(t, "بِسْمِ ٱللَّهِ ٱلرَّحْمَـٰنِ ٱلرَّحِيمِ", q.Suras[0].Ayas[0].Text)
}
