package pkg

import (
	"bufio"
	"os"
)

type Message struct {
	Key       byte
	Decrypted []byte
	score     float64
}

func DecryptXorSingleByte(bytes []byte) (m Message) {
	/* bytes, err := HexStrToBytes(str)
	if err != nil {
		fmt.Println(err)
		return Message{}
	} */
	m.score = -1
	for i := range 256 {
		byte := byte(i)
		xor := XorSingleByte(bytes, byte)
		score := ScoringEnglish(xor)
		if IsBetterScore(score, m.score) {
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

	m.score = -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		bytes, err := HexStrToBytes(str)
		check(err)
		mm := DecryptXorSingleByte(bytes)
		if IsBetterScore(mm.score, m.score) {
			m = mm
		}
	}
	check(scanner.Err())

	return m
}
