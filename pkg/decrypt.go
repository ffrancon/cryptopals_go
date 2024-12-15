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

func DecryptXorSingleByteRange(bytes []byte, start, end int, c chan Message) {
	m := Message{Score: -1}
	for i := start; i < end; i++ {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := ScoringEnglish(xor)
		if IsBetterScore(score, m.Score) {
			m = Message{Key: byte, Decrypted: xor, Score: score}
		}
	}
	c <- m
}

func DecryptXorSingleByte(bytes []byte, index int) Message {
	winner := make(chan Message)
	for i := 0; i <= 256; i += 64 {
		go DecryptXorSingleByteRange(bytes, i, i+63, winner)
	}
	return <-winner
}

func DecryptXorSingleByteFromBatchFile(path string) (m Message) {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	m.Score = -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		bytes := HexStrToBytes(str)
		mm := DecryptXorSingleByte(bytes, 999)
		if IsBetterScore(mm.Score, m.Score) {
			m = mm
		}
	}
	Check(scanner.Err())

	return m
}
