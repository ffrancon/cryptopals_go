package pkg

import (
	"bufio"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes = append(bytes, scanner.Bytes()...)
	}
	check(scanner.Err())

	return bytes
}
