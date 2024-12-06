package pkg

import (
	"math"
	"regexp"
)

type character struct {
	char string
	rate float64
}

var characters = []character{
	{"E", 0.127},
	{"T", 0.091},
	{"A", 0.082},
	{"O", 0.075},
	{"I", 0.070},
	{"N", 0.067},
	{" ", 0.063},
	{"S", 0.063},
	{"H", 0.061},
	{"R", 0.060},
	{"D", 0.043},
	{"L", 0.040},
	{"U", 0.028},
}

var nonAsciiRegpex = regexp.MustCompile(`[^A-Za-z]`)

func computeRegexps() map[string]*regexp.Regexp {
	regexps := make(map[string]*regexp.Regexp)
	for _, c := range characters {
		regexps[c.char] = regexp.MustCompile(`(?i)` + c.char)
	}
	return regexps
}

// returns a score based on the frequency of english characters in the input bytes
func EvaluateEnglish(bytes []byte) (score float64) {
	// if there are non-english characters, it's not english
	nonAscii := len(nonAsciiRegpex.FindAllIndex(bytes, -1))
	if nonAscii > 0 {
		return 999
	}

	regexps := computeRegexps()
	for _, c := range characters {
		rate := float64(len(regexps[c.char].FindAllIndex(bytes, -1))) / float64(len(bytes))
		score += math.Pow(rate-c.rate, 2) / c.rate
	}

	return score
}

func ChunkBytes(bytes []byte, size int) [][]byte {
	res := make([][]byte, len(bytes)/size)
	for i := 0; i < len(bytes); i += size {
		if i+size > len(bytes) {
			res = append(res, bytes[i:])
		}
		res[i/size] = bytes[i : i+size]
	}
	return res
}
