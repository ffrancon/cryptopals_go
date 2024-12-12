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
		by := scanner.Bytes()
		check(err)
		bytes = append(bytes, by...)
	}
	check(scanner.Err())
	return bytes
}
func main() {
	// get data from file
	data := scan("./data/6.txt")
	hex := pkg.BytesToHexStr(data)
	bytes, err := pkg.HexStrToBytes(hex)
	check(err)
	// determine key size
	ks := pkg.DetermineBestKeySize(bytes, 2, 80)
	chunks := pkg.ChunkBytes(bytes, ks)

	// transpose chunks
	// [[1, 2], [3, 4], [5, 6], [7, 8]] -> [[1, 3, 5, 7], [2, 4, 6, 8]] clen = 4
	t := make([][]byte, ks)
	clen := len(chunks)
	for i := 0; i < ks; i++ {
		for j := 0; j < clen; j++ {
			if i < len(chunks[j]) {
				t[i] = append(t[i], chunks[j][i])
			}
		}
	}
	fmt.Println(t)
	key := make([]byte, ks)

	for x := range t {
		m := pkg.DecryptXorSingleByte(t[x])
		fmt.Printf("dec: %s, key: %d, score: %f\n", string(m.Decrypted), m.Key, m.Score)
		key[x] = m.Key
	}
	fmt.Printf("key: %v", key)
}
