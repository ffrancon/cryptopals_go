package main

import (
	"ffrancon/cryptopals-go/pkg"
	"fmt"
)

func main() {
	res := pkg.DecryptXorSingleByteFromBatchFile("./data/14.txt")
	fmt.Printf("\nKey: %d, Decrypted: %s\n", res.Key, string(res.Decrypted))
}
