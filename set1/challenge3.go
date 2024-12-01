package set1

import (
	"ffrancon/cryptopals-go/utils"
	"fmt"
)

func Challenge3(str string) {
	bytes := utils.HexStrToBytes(str)
	if bytes == nil {
		return
	}
	for i := 0; i < 256; i++ {
		copy := make([]byte, len(bytes))
		byte := byte(i)
		for j := range bytes {
			copy[j] = bytes[j] ^ byte
		}
		fmt.Printf("%d => %s\n", i, string(copy[:]))
	}
}
