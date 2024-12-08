package pkg

import (
	"math"
	"regexp"
)

var freqTable = map[byte]float64{
	'A':  0.08167, // A
	'B':  0.01492, // B
	'C':  0.02782, // C
	'D':  0.04253, // D
	'E':  0.12702, // E
	'F':  0.02228, // F
	'G':  0.02015, // G
	'H':  0.06094, // H
	'I':  0.06966, // I
	'J':  0.00153, // J
	'K':  0.00772, // K
	'L':  0.04025, // L
	'M':  0.02406, // M
	'N':  0.06749, // N
	'O':  0.07507, // O
	'P':  0.01929, // P
	'Q':  0.00095, // Q
	'R':  0.05987, // R
	'S':  0.06327, // S
	'T':  0.09056, // T
	'U':  0.02758, // U
	'V':  0.00978, // V
	'W':  0.02360, // W
	'X':  0.00150, // X
	'Y':  0.01974, // Y
	'Z':  0.00074, // Z
	' ':  0.13000, // Space
	'!':  0.01000, // !
	'?':  0.01000, // ?
	'"':  0.01000, // "
	'\'': 0.01000, // '
	',':  0.01000, // ,
	'.':  0.01000, // .
	':':  0.01000, // :
	';':  0.01000, // ;
	'\n': 0.01000, // \n
}

var nonEnglishCharRegexp = regexp.MustCompile(`[^a-zA-Z\s!?":;.,']`)

func ScoringEnglish(bytes []byte) (score float64) {
	if len(nonEnglishCharRegexp.FindAllIndex(bytes, -1)) > 0 {
		return -1
	}
	charCount := make(map[byte]int)
	for _, b := range bytes {
		// convert lowercase to uppercase
		if b >= 97 && b <= 122 {
			charCount[b-32] = charCount[b-32] + 1
		} else {
			charCount[b] = charCount[b] + 1
		}
	}
	// Chi-square test
	for b, r := range freqTable {
		occ := float64(charCount[b])
		expOcc := float64(len(bytes)) * r
		score += math.Pow(occ-expOcc, 2) / expOcc
	}
	return score
}

func IsBetterScore(score, bestScore float64) bool {
	return bestScore == -1 || (score >= 0 && score < bestScore)
}
