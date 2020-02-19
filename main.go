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

	chunk1 := New(file, chunkSize)
	chunk2 := New(file, chunkSize)

	isEqual, diffs, err := chunk1.Diff(chunk2)
	if err != nil {
		panic(err)
	}
	if isEqual {
		fmt.Println("True")
	} else {
		fmt.Printf("False")
		fmt.Println(diffs)
	}
}
