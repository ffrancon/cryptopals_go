package set1

import (
	"ffrancon/cryptopals-go/utils"
	"fmt"
)

func Chal2(buf1, buf2 string) string {
	bytes1 := utils.HexStrToBytes(buf1)
	if bytes1 == nil {
		fmt.Printf("invalid hex string: %s", buf1)
		return ""
	}
	bytes2 := utils.HexStrToBytes(buf2)
	if bytes2 == nil {
		fmt.Printf("invalid hex string: %s", buf2)
		return ""
	}
	if len(bytes1) != len(bytes2) {
		fmt.Println("buffers are not of the same length")
	}

	bytes3 := make([]byte, len(bytes1))
	for i := range bytes3 {
		bytes3[i] = bytes1[i] ^ bytes2[i]
	}
	return utils.BytesToHexStr(bytes3)
}
