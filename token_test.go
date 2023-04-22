package quran_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"testing"
	
	"github.com/edipermadi/go-quran-tree"
	"github.com/stretchr/testify/require"
)

func TestToken_ParseAllCharacters(t *testing.T) {
	file, err := os.Open("quran-uthmani.xml")
	require.NoError(t, err)
	defer func() { _ = file.Close() }()

	// decode
	q := &quran.Quran{}
	require.NoError(t, xml.NewDecoder(file).Decode(q))

	for suraNum, sura := range q.Suras {
		for ayaNum, aya := range sura.Ayas {
			for tokenNum, token := range aya.Tokens() {
				// decode all character without panic
				require.NotPanics(t, func() {
					characters := token.Characters()
					require.NotEmpty(t, characters, fmt.Sprintf("error while parsing sura %d aya %d token %d: %s", suraNum+1, ayaNum+1, tokenNum+1, token.Dump()))
				}, fmt.Sprintf("error while parsing sura %d aya %d token %d: %s", suraNum+1, ayaNum+1, tokenNum+1, token.Dump()))
			}
		}
	}

	tokens := q.Suras[0].Ayas[0].Tokens()
	require.NotEmpty(t, tokens)
	characters := tokens[0].Characters()
	for _, v := range characters {
		t.Logf("%s", string(v.Letter))
	}
}

func TestToken_RemoveTashkeels(t *testing.T) {
	file, err := os.Open("quran-uthmani.xml")
	require.NoError(t, err)
	defer func() { _ = file.Close() }()

	// decode
	q := &quran.Quran{}
	require.NoError(t, xml.NewDecoder(file).Decode(q))

	tokens := q.Suras[0].Ayas[0].Tokens()
	require.NotEmpty(t, tokens)
	t.Logf(tokens[0].Text)
	t.Logf(tokens[0].RemoveTashkeels())
}

func TestToken_Characters(t *testing.T) {
	token := quran.Token{Text: strings.TrimSpace("بِسْم")}
	characters := token.Characters()
	require.Equal(t, 3, len(characters))
}
