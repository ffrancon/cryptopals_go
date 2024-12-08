package main

import (
	"ffrancon/cryptopals-go/pkg"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	data, err := os.ReadFile("./data/6.txt")
	check(err)
	k := pkg.DetermineBestKeySize(data, 2, 40)
	chunks := pkg.ChunkBytes(data, k)
	fmt.Println(chunks)
}
