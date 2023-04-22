package quran

import (
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

type runeForms struct {
	Isolated   rune
	Beginning  rune
	Medium     rune
	Final      rune
	Tashkeel   bool
	Annotation bool
}

var allRuneForms = map[rune]runeForms{
	0x0621: runeForms{
		// ARABIC LETTER HAMZA
		Isolated:  0xfe80,
		Beginning: 0x0621,
		Medium:    0x0621,
		Final:     0x0621,
	},
	0x0623: runeForms{
		// ARABIC LETTER ALEF WITH HAMZA ABOVE
		Isolated:  0xfe83,
		Beginning: 0x0623,
		Medium:    0xfe84,
		Final:     0xfe84,
	},
	0x0624: runeForms{
		// ARABIC LETTER WAW WITH HAMZA ABOVE
		Isolated:  0xfe85,
		Beginning: 0x0624,
		Medium:    0xfe86,
		Final:     0xfe86,
	},
	0x0625: runeForms{
		// ARABIC LETTER ALEF WITH HAMZA BELOW
		Isolated:  0xfe87,
		Beginning: 0x0625,
		Medium:    0xfe88,
		Final:     0xfe88,
	},
	0x0626: runeForms{
		// ARABIC LETTER YEH WITH HAMZA ABOVE
		Isolated:  0xfe89,
		Beginning: 0xfe8b,
		Medium:    0xfe8c,
		Final:     0xfe8a,
	},
	0x0627: runeForms{
		// ARABIC LETTER ALEF
		Isolated:  0xfe8d,
		Beginning: 0x0627,
		Medium:    0xfe8e,
		Final:     0xfe8e,
	},
	0x0628: runeForms{
		// ARABIC LETTER BEH
		Isolated:  0xfe8f,
		Beginning: 0xfe91,
		Medium:    0xfe92,
		Final:     0xfe90,
	},
	0x0629: runeForms{
		// ARABIC LETTER TEH MARBUTA
		Isolated:  0xfe93,
		Beginning: 0x0629,
		Medium:    0x0629,
		Final:     0xfe94,
	},
	0x062a: runeForms{
		// ARABIC LETTER TEH
		Isolated:  0xfe95,
		Beginning: 0xfe97,
		Medium:    0xfe98,
		Final:     0xfe96,
	},
	0x062b: runeForms{
		// ARABIC LETTER THEH
		Isolated:  0xfe99,
		Beginning: 0xfe9b,
		Medium:    0xfe9c,
		Final:     0xfe9a,
	},
	0x062c: runeForms{
		// ARABIC LETTER JEEM
		Isolated:  0xfe9d,
		Beginning: 0xfe9f,
		Medium:    0xfea0,
		Final:     0xfe9e,
	},
	0x062d: runeForms{
		// ARABIC LETTER HAH
		Isolated:  0xfea1,
		Beginning: 0xfea3,
		Medium:    0xfea4,
		Final:     0xfea2,
	},
	0x062e: runeForms{
		// ARABIC LETTER KHAH
		Isolated:  0xfea5,
		Beginning: 0xfea7,
		Medium:    0xfea8,
		Final:     0xfea6,
	},
	0x062f: runeForms{
		// ARABIC LETTER DAL
		Isolated:  0xfea9,
		Beginning: 0x062f,
		Medium:    0xfeaa,
		Final:     0xfeaa,
	},
	0x0630: runeForms{
		// ARABIC LETTER THAL
		Isolated:  0xfeab,
		Beginning: 0x0630,
		Medium:    0xfeac,
		Final:     0xfeac,
	},
	0x0631: runeForms{
		// ARABIC LETTER REH
		Isolated:  0x0fead,
		Beginning: 0x00631,
		Medium:    0x0feae,
		Final:     0x0feae,
	},
	0x0632: runeForms{
		// ARABIC LETTER ZAIN
		Isolated:  0xfeaf,
		Beginning: 0x0632,
		Medium:    0xfeb0,
		Final:     0xfeb0,
	},
	0x0633: runeForms{
		// ARABIC LETTER SEEN
		Isolated:  0xfeb1,
		Beginning: 0xfeb3,
		Medium:    0xfeb4,
		Final:     0xfeb2,
	},
	0x0634: runeForms{
		// ARABIC LETTER SHEEN
		Isolated:  0xfeb5,
		Beginning: 0xfeb7,
		Medium:    0xfeb8,
		Final:     0xfeb6,
	},
	0x0635: runeForms{
		// ARABIC LETTER SAD
		Isolated:  0xfeb9,
		Beginning: 0xfebb,
		Medium:    0xfebc,
		Final:     0xfeba,
	},
	0x0636: runeForms{
		// ARABIC LETTER DAD
		Isolated:  0xfebd,
		Beginning: 0xfebf,
		Medium:    0xfec0,
		Final:     0xfebe,
	},
	0x0637: runeForms{
		// ARABIC LETTER TAH
		Isolated:  0xfec1,
		Beginning: 0xfec3,
		Medium:    0xfec4,
		Final:     0xfec2,
	},
	0x0638: runeForms{
		// ARABIC LETTER ZAH
		Isolated:  0xfec5,
		Beginning: 0xfec7,
		Medium:    0xfec8,
		Final:     0xfec6,
	},
	0x0639: runeForms{
		// ARABIC LETTER AIN
		Isolated:  0xfec9,
		Beginning: 0xfecb,
		Medium:    0xfecc,
		Final:     0xfeca,
	},
	0x063a: runeForms{
		// ARABIC LETTER GHAIN
		Isolated:  0xfecd,
		Beginning: 0xfecf,
		Medium:    0xfed0,
		Final:     0xfece,
	},
	0x0640: runeForms{
		// ARABIC TATWEEL
		Isolated:  0x0640,
		Beginning: 0x0640,
		Medium:    0x0640,
		Final:     0x0640,
	},
	0x0641: runeForms{
		// ARABIC LETTER FEH
		Isolated:  0xfed1,
		Beginning: 0xfed3,
		Medium:    0xfed4,
		Final:     0xfed2,
	},
	0x0642: runeForms{
		// ARABIC LETTER QAF
		Isolated:  0xfed5,
		Beginning: 0xfed7,
		Medium:    0xfed8,
		Final:     0xfed6,
	},
	0x0643: runeForms{
		// ARABIC LETTER KAF
		Isolated:  0xfed9,
		Beginning: 0xfedb,
		Medium:    0xfedc,
		Final:     0xfeda,
	},
	0x0644: runeForms{
		// ARABIC LETTER LAM
		Isolated:  0xedd,
		Beginning: 0xfedf,
		Medium:    0xfee0,
		Final:     0xfede,
	},
	0x0645: runeForms{
		// ARABIC LETTER MEEM
		Isolated:  0xfee1,
		Beginning: 0xfee3,
		Medium:    0xfee4,
		Final:     0xfee2,
	},
	0x0646: runeForms{
		// ARABIC LETTER NOON
		Isolated:  0xfee5,
		Beginning: 0xfee7,
		Medium:    0xfee8,
		Final:     0xfee6,
	},
	0x0647: runeForms{
		// ARABIC LETTER HEH
		Isolated:  0xfee9,
		Beginning: 0xfeeb,
		Medium:    0xfeec,
		Final:     0xfeea,
	},
	0x0648: runeForms{
		// ARABIC LETTER WAW
		Isolated:  0xfeed,
		Beginning: 0x0648,
		Medium:    0xfeee,
		Final:     0xfeee,
	},
	0x0649: runeForms{
		// ARABIC LETTER ALEF MAKSURA
		Isolated:  0xfeef,
		Beginning: 0x0649,
		Medium:    0xfef0,
		Final:     0xfef0,
	},
	0x064a: runeForms{
		// ARABIC LETTER YEH
		Isolated:  0xfef1,
		Beginning: 0xfef3,
		Medium:    0xfef4,
		Final:     0xfef2,
	},
	0x064b: runeForms{
		// ARABIC FATHATAN
		Tashkeel: true,
	},
	0x064c: runeForms{
		// ARABIC DAMMATAN
		Tashkeel: true,
	},
	0x064d: runeForms{
		// ARABIC KASRATAN
		Tashkeel: true,
	},
	0x064e: runeForms{
		// ARABIC FATHA
		Tashkeel: true,
	},
	0x064f: runeForms{
		// ARABIC DAMMA
		Tashkeel: true,
	},
	0x0650: runeForms{
		// ARABIC KASRA
		Tashkeel: true,
	},
	0x0651: runeForms{
		// ARABIC SHADDA
		Tashkeel: true,
	},
	0x0652: runeForms{
		// ARABIC SUKUN
		Tashkeel: true,
	},
	0x0653: runeForms{
		// ARABIC MADDAH ABOVE
		Tashkeel: true,
	},
	0x0654: runeForms{
		// ARABIC HAMZA ABOVE
		Tashkeel: true,
	},
	0x0670: runeForms{
		// ARABIC LETTER SUPERSCRIPT ALEF
		Tashkeel: true,
	},
	0x0671: runeForms{
		// ARABIC LETTER ALEF WASLA
		Tashkeel: true,
	},
	0x06d6: runeForms{
		// ARABIC SMALL HIGH LIGATURE SAD WITH LAM WITH ALEF MAKSURA
		Annotation: true,
	},
	0x06d7: runeForms{
		// ARABIC SMALL HIGH LIGATURE QAF WITH LAM WITH ALEF MAKSURA
		Annotation: true,
	},
	0x06d8: runeForms{
		// ARABIC SMALL HIGH MEEM INITIAL FORM
		Annotation: true,
	},
	0x06d9: runeForms{
		// ARABIC SMALL HIGH LAM ALEF
		Annotation: true,
	},
	0x06da: runeForms{
		// ARABIC SMALL HIGH JEEM
		Annotation: true,
	},
	0x06db: runeForms{
		// ARABIC SMALL HIGH THREE DOTS
		Annotation: true,
	},
	0x06dc: runeForms{
		// ARABIC SMALL HIGH SEEN
		Annotation: true,
	},
	0x06de: runeForms{
		// ARABIC START OF RUB EL HIZB
		Annotation: true,
	},
	0x06df: runeForms{
		// ARABIC SMALL HIGH ROUNDED ZERO
		Annotation: true,
	},
	0x06e0: runeForms{
		// ARABIC SMALL HIGH UPRIGHT RECTANGULAR ZERO
		Annotation: true,
	},
	0x06e2: runeForms{
		// ARABIC SMALL HIGH MEEM ISOLATED FORM
		Annotation: true,
	},
	0x06e3: runeForms{
		// ARABIC SMALL LOW SEEN
		Annotation: true,
	},
	0x06e5: runeForms{
		// ARABIC SMALL WAW
		Annotation: true,
	},
	0x06e6: runeForms{
		// ARABIC SMALL YEH
		Annotation: true,
	},
	0x06e7: runeForms{
		// ARABIC SMALL HIGH YEH
		Annotation: true,
	},
	0x06e8: runeForms{
		// ARABIC SMALL HIGH NOON
		Annotation: true,
	},
	0x06e9: runeForms{
		// ARABIC PLACE OF SAJDAH
		Annotation: true,
	},
	0x06ea: runeForms{
		// ARABIC EMPTY CENTRE LOW STOP
		Annotation: true,
	},
	0x06eb: runeForms{
		// ARABIC EMPTY CENTRE HIGH STOP
		Annotation: true,
	},
	0x06ec: runeForms{
		// ARABIC ROUNDED HIGH STOP WITH FILLED CENTRE
		Annotation: true,
	},
	0x06ed: runeForms{
		// ARABIC SMALL LOW MEEM
		Annotation: true,
	},
}



var normalizable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0640, 0x0640, 1}, // ARABIC TATWEEL
		{0x064b, 0x064b, 1}, // ARABIC FATHATAN
		{0x064c, 0x064c, 1}, // ARABIC DAMMATAN
		{0x064d, 0x064d, 1}, // ARABIC KASRATAN
		{0x064e, 0x064e, 1}, // ARABIC FATHA
		{0x064f, 0x064f, 1}, // ARABIC DAMMA
		{0x0650, 0x0650, 1}, // ARABIC KASRA
		{0x0651, 0x0651, 1}, // ARABIC SHADDA
		{0x0652, 0x0652, 1}, // ARABIC SUKUN
		{0x0653, 0x0653, 1}, // ARABIC MADDAH ABOVE
		{0x0670, 0x0670, 1}, // ARABIC LETTER SUPERSCRIPT ALEF
	},
}

func removeTashkeels(input string)string{
	tm := transform.Chain(runes.Remove(runes.In(normalizable)))
	output, _, _ := transform.String(tm, input)
	return output
}

func isAllArabicAnnotation(s string) bool {
	for _, v := range s {
		form, found := allRuneForms[v]
		if !found {
			panic(fmt.Errorf("unexpected rune 0x%04x", v))
		} else if !form.Annotation {
			return false
		}
	}
	return true
}
