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

var nonAsciiRegpex = regexp.MustCompile(`[^\x00-\x7F]`)

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
		return 0
	}

	regexps := computeRegexps()
	for _, c := range characters {
		occ := len(regexps[c.char].FindAllIndex(bytes, -1))
		score += float64(occ) * c.rate
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
