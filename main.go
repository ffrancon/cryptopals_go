package main

import (
	"bufio"
	"ffrancon/cryptopals-go/pkg"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func scan(path string) (str string) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		check(err)
		str += pkg.Base64ToString(scanner.Text())
	}
	check(scanner.Err())
	return str
}

/* func main() {
	scan("./data/6.txt")
} */

func main() {
	// get data from file
	dec := scan("./data/6.txt")
	bytes := []byte(dec)
	// determine key size
	ks := pkg.DetermineBestKeySize(bytes, 2, 80)
	chunks := pkg.ChunkBytes(bytes, ks)
	transposed := pkg.TransposeBytesChunks(chunks)
	key := make([]byte, ks)

	for x := range transposed {
		m := pkg.DecryptXorSingleByte(transposed[x])
		// fmt.Printf("dec: %s, key: %d, score: %f\n\n", string(m.Decrypted), m.Key, m.Score)
		key[x] = m.Key
	}
	fmt.Printf("key: %v", key)
}
