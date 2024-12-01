package pkg

import (
	"fmt"
	"regexp"
)

type Message struct {
	Key       byte
	Decrypted []byte
	score     int
}

func DecryptXorSingleByte(str string) (m Message) {
	bytes, err := HexStrToBytes(str)
	if err != nil {
		fmt.Println(err)
		return Message{}
	}
	reg := regexp.MustCompile(`(?i)[ETAOIN SHRDLU]`)
	for i := 0; i < 256; i++ {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := len(reg.FindAll(xor, -1))
		if score > m.score {
			m = Message{byte, xor, score}
		}
	}
	return m
}
