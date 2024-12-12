package pkg

import (
	"errors"
	"fmt"
)

func ComputeHammingDistance(bytes1, bytes2 []byte) (d int, e error) {
	if len(bytes1) != len(bytes2) {
		return -1, errors.New("byte arrays are not of the same length")
	}
	for i, b := range bytes1 {
		xor := int(b ^ bytes2[i])
		for xor > 0 {
			d += xor & 1
			xor >>= 1
		}
	}
	return d, nil
}

func ComputeNormalizedHammingDistance(bytes []byte, s int) (float64, error) {
	if len(bytes) < 4*s {
		return -1, errors.New("byte array is too short")
	}
	fs := float64(s)

	d1, _ := ComputeHammingDistance(bytes[:s], bytes[s:2*s])
	n1 := float64(d1) / fs

	d2, _ := ComputeHammingDistance(bytes[s:2*s], bytes[2*s:3*s])
	n2 := float64(d2) / fs

	d3, _ := ComputeHammingDistance(bytes[:s], bytes[2*s:3*s])
	n3 := float64(d3) / fs

	d4, _ := ComputeHammingDistance(bytes[2*s:3*s], bytes[3*s:4*s])
	n4 := float64(d4) / fs

	return (n1 + n2 + n3 + n4) / 4, nil
}

// [1, 2, 3, 4, 5, 6, 7, 8] -> [[1, 2], [3, 4], [5, 6], [7, 8]]
func ChunkBytes(bytes []byte, s int) (chunks [][]byte) {
	fmt.Printf("len(bytes): %d\n", len(bytes))
	end := len(bytes)
	for i := 0; i < end; i += s {
		to := i + s
		if to > end {
			to = end
		}
		chunks = append(chunks, bytes[i:to])

	}
	return chunks
}

func DetermineBestKeySize(bytes []byte, min, max int) (s int) {
	d := float64(-1)
	for i := min; i < max; i++ {
		nd, _ := ComputeNormalizedHammingDistance(bytes, i)
		if d == -1 || nd < d {
			d = nd
			s = i
		}
	}
	return s
}
