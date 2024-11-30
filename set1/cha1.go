package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexStrToBase64(str string) []byte {
	data, err := hex.DecodeString(str)

	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return nil
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)

	return dst
}
