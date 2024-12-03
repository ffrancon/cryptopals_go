package pkg

import (
	"fmt"
)

type Message struct {
	Key       byte
	Decrypted []byte
	score     float64
}

func DecryptXorSingleByte(str string) (m Message) {
	bytes, err := HexStrToBytes(str)
	if err != nil {
		fmt.Println(err)
		return Message{}
	}
	for i := range 256 {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := EvaluateEnglish(xor)
		if score > m.score {
			m = Message{byte, xor, score}
		}
	}
	return m
}
