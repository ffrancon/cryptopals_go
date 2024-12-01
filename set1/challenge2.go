package set1

import (
	"ffrancon/cryptopals-go/utils"
	"fmt"
)

func handleStrToBytesConversion(buf1, buf2 string) ([]byte, []byte, error) {
	bytes1 := utils.HexStrToBytes(buf1)
	if bytes1 == nil {
		return nil, nil, fmt.Errorf("invalid hex string: %s", buf1)
	}
	bytes2 := utils.HexStrToBytes(buf2)
	if bytes2 == nil {
		return nil, nil, fmt.Errorf("invalid hex string: %s", buf2)
	}
	return bytes1, bytes2, nil
}

func Challenge2(buf1, buf2 string) (result string) {
	bytes1, bytes2, err := handleStrToBytesConversion(buf1, buf2)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if len(bytes1) != len(bytes2) {
		fmt.Println("buffers are not of the same length")
		return ""
	}

	bytes3 := make([]byte, len(bytes1))
	for i := range bytes3 {
		bytes3[i] = bytes1[i] ^ bytes2[i]
	}
	result = utils.BytesToHexStr(bytes3)
	return result
}
