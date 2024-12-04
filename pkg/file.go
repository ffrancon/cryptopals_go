package pkg

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	fs, err := file.Stat()
	check(err)

	bytes := make([]byte, fs.Size())
	_, err = file.Read(bytes)
	check(err)

	return bytes
}
