package quran

const (
	beginArabicLetter = ArabicLetterKashmiriYeh
	endArabicLetter   = ArabicLetterYeh

	beginArabicDiacritic = ArabicFathatan
	endArabicDiacritic   = ArabicWavyHamzaBelow
)
const (
	ArabicSmallHighZain     rune = 0x0617
	ArabicSmallFatha        rune = 0x0618
	ArabicSmallDamma        rune = 0x0619
	ArabicSmallKasra        rune = 0x061a
	ArabicLetterKashmiriYeh rune = 0x0620
	ArabicLetterYeh         rune = 0x064a
	ArabicFathatan          rune = 0x064b
	ArabicWavyHamzaBelow    rune = 0x065f

	ArabicSmallHighLigatureSadWithLamWithAlefMaksura rune = 0x06d6
	ArabicSmallHighLigatureQafWithLamWithAlefMaksura rune = 0x06d7
	ArabicSmallHighMeemInitialForm                   rune = 0x06d8
	ArabicSmallHighLamAlef                           rune = 0x06d9
	ArabicSmallHighJeem                              rune = 0x06da
	ArabicSmallHighThreeDots                         rune = 0x06db
	ArabicSmallHighSeen                              rune = 0x06dc
	ArabicEndOfAyah                                  rune = 0x06dd
	ArabicStartOfRubElHizb                           rune = 0x06de
	ArabicSmallHighROUNDEDZero                       rune = 0x06df
	ArabicSmallHighUprightRectangularZero            rune = 0x06e0
	ArabicSmallHighDotlessHeadOfKhah                 rune = 0x06e1
	ArabicSmallHighMeemIsolatedForm                  rune = 0x06e2
	ArabicSmallLowSeen                               rune = 0x06e3
	ArabicSmallHighMadda                             rune = 0x06e4
	ArabicSmallWaw                                   rune = 0x06e5
	ArabicSmallYeh                                   rune = 0x06e6
	ArabicSmallHighYeh                               rune = 0x06e7
	ArabicSmallHighNoon                              rune = 0x06e8
	ArabicPlaceOfSajdah                              rune = 0x06e9
	ArabicEmptyCentreLowStop                         rune = 0x06aa
	ArabicEmptyCentreHighStop                        rune = 0x06eb
	ArabicRoundedHighStopWithFilledCentre            rune = 0x06ec
	ArabicSmallLowMeem                               rune = 0x06ed

	ArabicLetterAlefWasla              rune = 0x0671
	ArabicLetterAlefWithWavyHamzaAbove rune = 0x0672
	ArabicLetterSuperscriptAlef        rune = 0x0670
)

func isArabicLetter(v rune) bool {
	switch {
	case v >= beginArabicLetter && v <= endArabicLetter:
		return true
	case v == ArabicLetterAlefWasla,
		v == ArabicLetterAlefWithWavyHamzaAbove:
		return true
	default:
		return false
	}
}

func isArabicTashkeel(v rune) bool {
	switch {
	case v >= beginArabicDiacritic && v <= endArabicDiacritic:
		return true
	case v == ArabicLetterSuperscriptAlef:
		return true
	default:
		return false
	}
}

func isArabicAnnotation(v rune) bool {
	switch {
	case v >= ArabicSmallHighLigatureSadWithLamWithAlefMaksura && v <= ArabicSmallLowMeem:
		return true
	default:
		return false
	}
}

func isAllArabicAnnotation(s string) bool {
	for _, v := range s {
		if !isArabicAnnotation(v) {
			return false
		}
	}

	return true
}

var allRunes = []rune{0x0020, 0x0621, 0x0623, 0x0624, 0x0625, 0x0626, 0x0627, 0x0628, 0x0629, 0x062a, 0x062b, 0x062c, 0x062d, 0x062e, 0x062f, 0x0630, 0x0631, 0x0632, 0x0633, 0x0634, 0x0635, 0x0636, 0x0637, 0x0638, 0x0639, 0x063a, 0x0640, 0x0641, 0x0642, 0x0643, 0x0644, 0x0645, 0x0646, 0x0647, 0x0648, 0x0649, 0x064a, 0x064b, 0x064c, 0x064d, 0x064e, 0x064f, 0x0650, 0x0651, 0x0652, 0x0653, 0x0654, 0x0670, 0x0671, 0x06d6, 0x06d7, 0x06d8, 0x06d9, 0x06da, 0x06db, 0x06dc, 0x06de, 0x06df, 0x06e0, 0x06e2, 0x06e3, 0x06e5, 0x06e6, 0x06e7, 0x06e8, 0x06e9, 0x06ea, 0x06eb, 0x06ec, 0x06ed}
