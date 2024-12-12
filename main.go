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
func scan(path string) (bytes []byte) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		check(err)
		fmt.Println(scanner.Text(), pkg.BytesToBase64(scanner.Bytes()))
		bytes = append(bytes, scanner.Bytes()...)
	}
	check(scanner.Err())
	return bytes
}

func main() {
	scan("./data/6.txt")
}

/* func main() {
	// get data from file
	data := scan("./data/6.txt")
	bytes, err := pkg.Base64ToBytes(data)
	check(err)
	// determine key size
	ks := pkg.DetermineBestKeySize(bytes, 2, 80)
	chunks := pkg.ChunkBytes(bytes, ks)
	transposed := pkg.TransposeBytesChunks(chunks)
	key := make([]byte, ks)

	for x := range transposed {
		m := pkg.DecryptXorSingleByte(transposed[x])
		fmt.Printf("hex: %s, dec: %s, key: %d, score: %f\n\n", pkg.BytesToHexStr(transposed[x]), string(m.Decrypted), m.Key, m.Score)
		key[x] = m.Key
	}
	fmt.Printf("key: %v", key)
} */
