package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	data, err := hex.DecodeString(hexStr)

	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))
}
