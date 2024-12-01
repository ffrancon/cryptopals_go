package set1

import (
	"encoding/base64"
	"ffrancon/cryptopals-go/utils"
)

func Challenge1(str string) (result []byte) {
	data := utils.HexStrToBytes(str)
	result = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(result, data)
	return result
}
