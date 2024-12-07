package pkg

import (
	"bufio"
	"fmt"
	"os"
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
	m.Score = 999
	for i := range 256 {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := ScoringEnglish(xor)
		if score < 900 {
			fmt.Printf("Key: %d, Decrypted: %s, Score: %f\n", byte, string(xor), score)
		}
		if score < m.Score {
			m = Message{byte, xor, score}
		}
	}
	return m
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func DecryptXorSingleByteFromBatchFile(path string) (m Message) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	m.Score = 999
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mm := DecryptXorSingleByte(scanner.Text())
		if mm.Score < m.Score {
			m = mm
		}
	}
	check(scanner.Err())

	return m
}
