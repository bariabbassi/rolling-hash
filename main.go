package main

import (
	"fmt"
	"io/ioutil"
)

const (
	fileName1 = "testfiles/longtext1.txt"
	fileName2 = "testfiles/longtext3.txt"
	chunkSize = 13
)

func main() {
	// Read file 1
	file1, err := ioutil.ReadFile(fileName1)
	if err != nil {
		panic(err)
	}

	// Read file 2
	file2, err := ioutil.ReadFile(fileName2)
	if err != nil {
		panic(err)
	}

	// Create chunkers
	chunker1 := New(file1, chunkSize)
	chunker2 := New(file2, chunkSize)

	// Diffing
	diffs, err := chunker1.Diff(chunker2)
	if err != nil {
		panic(err)
	}

	// Print unequal chunks' indexes
	fmt.Println(diffs)
}
