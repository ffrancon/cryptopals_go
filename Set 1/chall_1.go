package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func hexStrToBase64(str string) {
	data, err := hex.DecodeString(str)

	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))
}
