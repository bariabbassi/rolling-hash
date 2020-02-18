package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	chunk := New(file, chunkSize)

	for i := 0; i < chunk.fileSize-chunk.chunkSize-1; i++ {
		fmt.Println(chunk)
		err := chunk.Roll()
		if err != nil {
			log.Fatal(err)
		}
	}
}
