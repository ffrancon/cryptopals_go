package pkg

import (
	"fmt"
)

type Message struct {
	Key       byte
	Decrypted []byte
	Score     float64
}

func DecryptXorSingleByte(str string) (m Message) {
	bytes, err := HexStrToBytes(str)
	if err != nil {
		fmt.Println(err)
		return Message{}
	}
	m.Score = 9999
	for i := range 256 {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := EvaluateEnglish(xor)
		if score < m.Score {
			m = Message{byte, xor, score}
		}
	}
	fmt.Printf("Key: %d, Decrypted: %s, Score: %f\n", m.Key, string(m.Decrypted), m.Score)
	return m
}

func DecryptXorSingleByteBatch(batch [][]byte) (m Message) {
	for _, bytes := range batch {
		mm := DecryptXorSingleByte(BytesToHexStr(bytes))
		if mm.Score < m.Score {
			m = mm
		}
	}
	return m
}
