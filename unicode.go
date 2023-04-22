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
