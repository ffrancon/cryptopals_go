package set1

import (
	"ffrancon/cryptopals-go/utils"
	"regexp"
)

type message struct {
	key       byte
	decrypted []byte
	score     int
}

func Challenge3(str string) {
	bytes := utils.HexStrToBytes(str)
	if bytes == nil {
		return
	}
	messages := make([]message, 256)
	reg := regexp.MustCompile(`(?i)[ETAOIN SHRDLU]`)
	for i := 0; i < 256; i++ {
		copy := make([]byte, len(bytes))
		byte := byte(i)
		for j := range bytes {
			copy[j] = bytes[j] ^ byte
		}
		score := len(reg.FindAll(copy, -1))
		messages[byte] = message{byte, copy, score}
	}
}
