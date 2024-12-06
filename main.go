package main

import (
	"ffrancon/cryptopals-go/pkg"
	"fmt"
)

func main() {
	file := pkg.ReadFile("./data/14.txt")
	batch := pkg.ChunkBytes(file, 60)
	res := pkg.DecryptXorSingleByteBatch(batch)
	fmt.Printf("\nKey: %d, Decrypted: %s, Score: %f\n", res.Key, string(res.Decrypted), res.Score)
}
