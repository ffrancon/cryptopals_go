package set1

import (
	"encoding/base64"
	"ffrancon/cryptopals-go/utils"
)

func Chal1(str string) []byte {
	data := utils.HexStrToBytes(str)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)

	return dst
}
