package main

import (
	"fmt"
	"io/ioutil"
)

const (
	fileName  = "file.txt"
	chunkSize = 4
)

func main() {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	chunk := New(file)
	chunk.Roll()

	fmt.Println(chunk)
}
