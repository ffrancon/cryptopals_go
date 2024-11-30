package set1

import (
	"ffrancon/cryptopals-go/utils"
)

func Chal2(buf1 string, buf2 string) string {
	by1 := utils.HexStrToBytes(buf1)
	by2 := utils.HexStrToBytes(buf2)

	if by1 == nil || by2 == nil {
		return "Error decoding hex string"
	}
	if len(by1) != len(by2) {
		return "Buffers are not of the same length"
	}

	i := len(by1)
	res := make([]byte, i)

	for j := 0; j < i; j++ {
		res[j] = by1[j] ^ by2[j]
	}

	return utils.BytesToHexStr(res)
}
