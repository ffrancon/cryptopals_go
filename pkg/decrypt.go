package pkg

import (
	"bufio"
	"os"
)

type Message struct {
	Key       byte
	Decrypted []byte
	Score     float64
}

func DecryptXorSingleByte(bytes []byte) (m Message) {
	m.Score = -1
	for i := range 256 {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := ScoringEnglish(xor)
		/* fmt.Printf("dec: %s, key: %d, score: %f\n", string(xor), byte, score) */
		if IsBetterScore(score, m.Score) {
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

	m.Score = -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		bytes, err := HexStrToBytes(str)
		check(err)
		mm := DecryptXorSingleByte(bytes)
		if IsBetterScore(mm.Score, m.Score) {
			m = mm
		}
	}
	check(scanner.Err())

	return m
}
