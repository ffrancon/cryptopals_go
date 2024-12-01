package set1

import (
	"ffrancon/cryptopals-go/utils"
	"regexp"
)

type message struct {
	Key       byte
	Decrypted []byte
	score     int
}

func Challenge3(str string) (m message) {
	bytes := utils.HexStrToBytes(str)
	if bytes == nil {
		return message{}
	}
	reg := regexp.MustCompile(`(?i)[ETAOIN SHRDLU]`)
	for i := 0; i < 256; i++ {
		copy := make([]byte, len(bytes))
		byte := byte(i)
		for j := range bytes {
			copy[j] = bytes[j] ^ byte
		}
		score := len(reg.FindAll(copy, -1))
		if score > m.score {
			m = message{byte, copy, score}
		}
	}
	return m
}
