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
	d := float64(-1)
	k := 2
	for i := 2; i < 41; i++ {
		nd, _ := pkg.NormalizeKeySizeHammingDistance(data, i)
		if d == -1 || nd < d {
			d = nd
			k = i
		}
	}
	chunks := pkg.ChunkBytes(data, k)
	fmt.Println(chunks)
}
