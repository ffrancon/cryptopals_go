package pkg

import (
	"regexp"
)

type character struct {
	char string
	rate float64
}

var characters = []character{
	{"E", 12.7},
	{"T", 9.1},
	{"A", 8.2},
	{"O", 7.5},
	{"I", 7.0},
	{"N", 6.7},
	{" ", 6.3},
	{"S", 6.3},
	{"H", 6.1},
	{"R", 6.0},
	{"D", 4.3},
	{"L", 4.0},
	{"U", 2.8},
}

func EvaluateEnglish(bytes []byte) (score float64) {
	for _, c := range characters {
		reg := regexp.MustCompile(`(?i)` + c.char)
		occ := len(reg.FindAllIndex(bytes, -1))
		score += float64(occ) * c.rate
	}

	// Unrecognized english characters impact score as well
	unrecognizedReg := regexp.MustCompile(`[^A-Za-z0-9 .,?!:;']`)
	occ := len(unrecognizedReg.FindAllIndex(bytes, -1))
	score -= float64(occ) * 10

	return score
}
