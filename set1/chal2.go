package set1

import (
	"errors"
	"ffrancon/cryptopals-go/utils"
	"fmt"
)

func Chal2(buf1 string, buf2 string) (result string, error error) {
	bytes1 := utils.HexStrToBytes(buf1)
	bytes2 := utils.HexStrToBytes(buf2)

	if bytes1 == nil {
		return "", fmt.Errorf("error decoding hex string: %v", buf1)
	} else if bytes2 == nil {
		return "", fmt.Errorf("error decoding hex string: %v", buf2)
	} else if len(bytes1) != len(bytes2) {
		return "", errors.New("buffers are not of the same length")
	}

	bytes3 := make([]byte, len(bytes1))
	for i := range bytes3 {
		bytes3[i] = bytes1[i] ^ bytes2[i]
	}

	result = utils.BytesToHexStr(bytes3)
	return result, nil
}
