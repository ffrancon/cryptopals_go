package pkg

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexStrToBytes(str string) ([]byte, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, fmt.Errorf("error decoding hex string: %s", err)
	}
	return bytes, nil
}

func BytesToHexStr(data []byte) string {
	return hex.EncodeToString(data)
}

func HexStrToBase64(str string) ([]byte, error) {
	data, err := HexStrToBytes(str)
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(bytes, data)
	return bytes, nil
}

func BytesToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64ToString(data string) string {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Sprintf("error decoding base64 string: %s", err)
	}
	return string(bytes)
}
