package pkg

import (
	"errors"
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
	if len(bytes) < 3*s {
		return -1, errors.New("byte array is too short")
	}
	d1, _ := ComputeHammingDistance(bytes[:s], bytes[s:2*s])
	d2, _ := ComputeHammingDistance(bytes[s:2*s], bytes[2*s:3*s])
	avgDist := float64((d1 + d2) / 2)
	return avgDist / float64(s), nil
}

func ChunkBytes(bytes []byte, s int) (chunks [][]byte) {
	for i := 0; i < len(bytes); i += s {
		chunks = append(chunks, bytes[i:i+s])
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
