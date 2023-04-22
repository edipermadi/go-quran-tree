package quran_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/edipermadi/go-quran-tree"
	"github.com/stretchr/testify/require"
)

func TestAya_Tokens(t *testing.T) {
	file, err := os.Open("quran-uthmani.xml")
	require.NoError(t, err)
	defer func() { _ = file.Close() }()

	// decode
	q := &quran.Quran{}
	require.NoError(t, xml.NewDecoder(file).Decode(q))

	tokens := q.Suras[0].Ayas[0].Tokens()
	require.NotEmpty(t, tokens)
	for _, v := range tokens {
		t.Logf(v.Text)
	}
}

func TestAya_RemoveTashkeels(t *testing.T) {
	file, err := os.Open("quran-uthmani.xml")
	require.NoError(t, err)
	defer func() { _ = file.Close() }()

	// decode
	q := &quran.Quran{}
	require.NoError(t, xml.NewDecoder(file).Decode(q))

	for _, v := range q.Suras[0].Ayas {
		t.Logf(v.RemoveTashkeels())
	}
}
