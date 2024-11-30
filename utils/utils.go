package utils

import (
	"encoding/hex"
	"fmt"
)

func HexStrToBytes(str string) []byte {
	data, err := hex.DecodeString(str)

	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return nil
	}

	return data
}

func BytesToHexStr(data []byte) string {
	return hex.EncodeToString(data)
}
