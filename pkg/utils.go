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

var nonEnglishCharRegexp = regexp.MustCompile(`[^A-Za-z0-9 .,?!:;']`)

func computeRegexps() map[string]*regexp.Regexp {
	regexps := make(map[string]*regexp.Regexp)
	for _, c := range characters {
		regexps[c.char] = regexp.MustCompile(`(?i)` + c.char)
	}
	return regexps
}

// returns a score based on the frequency of english characters in the input bytes
func EvaluateEnglish(bytes []byte) (score float64) {
	regexps := computeRegexps()
	for _, c := range characters {
		occ := len(regexps[c.char].FindAllIndex(bytes, -1))
		score += float64(occ) * c.rate
	}
	// non-english characters impact score as well
	occ := len(nonEnglishCharRegexp.FindAllIndex(bytes, -1))
	score -= float64(occ) * 10

	return score
}
