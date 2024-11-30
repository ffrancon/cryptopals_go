package utils

import (
	"encoding/hex"
	"fmt"
)

func HexStrToBytes(str string) []byte {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return nil
	}
	return bytes
}

func BytesToHexStr(data []byte) string {
	return hex.EncodeToString(data)
}
