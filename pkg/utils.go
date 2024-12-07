package pkg

import "fmt"

func ComputeHammingDistance(bytes1, bytes2 []byte) (d int) {
	if len(bytes1) != len(bytes2) {
		fmt.Println("bytes arrays are not of the same length")
		return 0
	}
	for i, b := range bytes1 {
		xor := int(b ^ bytes2[i])
		for xor > 0 {
			d += xor & 1
			xor >>= 1
		}
	}
	return d
}
