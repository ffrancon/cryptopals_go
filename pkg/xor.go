package pkg

import (
	"fmt"
)

func XorBytes(b1, b2 []byte) (bytes []byte) {
	bytes = make([]byte, len(b1))
	for i := range bytes {
		bytes[i] = b1[i] ^ b2[i]
	}
	return bytes
}

func XorSingleByte(bytes []byte, key byte) []byte {
	xor := make([]byte, len(bytes))
	for i := range bytes {
		xor[i] = bytes[i] ^ key
	}
	return xor
}

func XorBuffers(buf1, buf2 string) string {
	by1, e1 := HexStrToBytes(buf1)
	if e1 != nil {
		fmt.Println(e1)
		return ""
	}
	by2, e2 := HexStrToBytes(buf2)
	if e2 != nil {
		fmt.Println(e2)
	}
	if len(by1) != len(by2) {
		fmt.Println("buffers are not of the same length")
		return ""
	}
	bytes := XorBytes(by1, by2)
	result := BytesToHexStr(bytes)
	return result
}
