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
	if len(bytes) < 8*s {
		return -1, errors.New("byte array is too short")
	}
	// compute the normalized hamming distance for the first 4 chunks
	d1, _ := ComputeHammingDistance(bytes[:s], bytes[s:2*s])
	n1 := float64(d1 / s)
	d2, _ := ComputeHammingDistance(bytes[2*s:3*s], bytes[3*s:4*s])
	n2 := float64(d2 / s)
	d3, _ := ComputeHammingDistance(bytes[4*s:5*s], bytes[5*s:6*s])
	n3 := float64(d3 / s)
	d4, _ := ComputeHammingDistance(bytes[6*s:7*s], bytes[7*s:8*s])
	n4 := float64(d4 / s)
	// return the average of the normalized hamming distances
	avg := (n1 + n2 + n3 + n4) / 4
	return avg, nil
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

// [1, 2, 3, 4, 5, 6, 7, 8] -> [[1, 2], [3, 4], [5, 6], [7, 8]]
func ChunkBytes(bytes []byte, s int) (chunks [][]byte) {
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

// [[1, 2], [3, 4], [5, 6], [7]] -> [[1, 3, 5, 7], [2, 4, 6,]]
func TransposeBytesChunks(chunks [][]byte) [][]byte {
	chunksLength := len(chunks)
	singleChunkLength := len(chunks[0])
	transpose := make([][]byte, singleChunkLength)
	// create a new array with the length of the first chunk
	for x := 0; x < singleChunkLength; x++ {
		transpose[x] = make([]byte, chunksLength)
		// iterate over the chunks and add the byte to the new array
		for y := 0; y < chunksLength; y++ {
			if x < len(chunks[y]) {
				transpose[x][y] = chunks[y][x]
			}
		}
	}
	return transpose
}
