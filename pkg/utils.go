package pkg

func ComputeHammingDistance(bytes1, bytes2 []byte) (d int) {
	for i, byte := range bytes1 {
		xor := int(byte ^ bytes2[i])
		for i := 1; i <= xor; i *= 2 {
			if xor&i > 0 {
				d++
			}
		}
	}
	return d
}
