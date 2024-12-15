package main

import (
	"encoding/base64"
	"ffrancon/cryptopals-go/pkg"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// get data from file
	data, err := pkg.ReadFile("./data/6.txt")
	check(err)
	bytes, err := base64.StdEncoding.DecodeString(string(data))
	check(err)
	// determine key size
	ks := pkg.DetermineBestKeySize(bytes, 2, 40)

	chunks := pkg.ChunkBytes(bytes, ks)
	transposed := pkg.TransposeBytesChunks(chunks)
	key := make([]byte, ks)

	for x := range transposed {
		m := pkg.DecryptXorSingleByte(transposed[x], x)
		// fmt.Printf("dec: %s, key: %d, score: %f\n\n", string(m.Decrypted), m.Key, m.Score)
		key[x] = m.Key
	}
	decrypted := pkg.XorRepeatingKey(bytes, key)
	fmt.Println(string(decrypted))
}
