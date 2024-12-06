package main

import (
	"ffrancon/cryptopals-go/pkg"
	"fmt"
)

func main() {
	file := pkg.ReadFile("./data/14.txt")
	batch := pkg.ChunkBytes(file, 60)
	/* for _, b := range batch {
		fmt.Println(string(b))
	} */
	res := pkg.DecryptXorSingleByteBatch(batch)
	fmt.Printf("\nKey: %d, Decrypted: %s\n", res.Key, string(res.Decrypted))
}
