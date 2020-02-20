package main

import (
	"fmt"
)

// Chunker contains a chunk that rolls through a file
type Chunker struct {
	index     int    // The index of the first byte of the chunk in the file
	chunk     []byte // A subset of bytes from file
	file      []byte // data in bytes
	chunkSize int
	fileSize  int
	hash      int
}

// New creates a new chunker
func New(file []byte, csize int) *Chunker {
	return &Chunker{
		index:     0,
		chunk:     file[:csize],
		file:      file,
		chunkSize: csize,
		fileSize:  len(file),
		hash:      newHash(file[:csize]),
	}
}

// newHash creates the hash of the first chunk
func newHash(chunk []byte) int {
	hash := 0
	for i := 0; i < len(chunk); i++ {
		hash += int(chunk[i]) - 96
	}
	return hash
}

// String creates a readible string from chunker
func (c *Chunker) String() string {
	return fmt.Sprintf("%d %s %d", c.index, string(c.chunk), c.hash)
}

// rollHash creates the next hash
func (c *Chunker) rollHash(newByte, oldByte byte) {
	c.hash = c.hash + int(newByte) - int(oldByte)
}

// Roll rolls the chunk through the file
func (c *Chunker) Roll() error {
	if c.index+c.chunkSize >= c.fileSize {
		return fmt.Errorf("Can't roll. End of file reached")
	}
	newByte := c.file[c.index+c.chunkSize]

	c.rollHash(newByte, c.chunk[0])
	c.chunk = append(c.chunk, newByte)
	c.chunk = c.chunk[1:]
	c.index++
	return nil
}

// Reset sets the chunker back to the initial state
func (c *Chunker) Reset() {
	c.index = 0
	c.chunk = c.file[:c.chunkSize]
	c.file = c.file[:c.fileSize]
	c.hash = newHash(c.chunk)
}

// Diff finds the indexes of the unequal chunks in file
func (c1 *Chunker) Diff(c2 *Chunker) ([]int, error) {

	// The 2 chunkers must have the same file size and chunk size
	if c1.fileSize != c2.fileSize {
		return nil, fmt.Errorf("The 2 file sizes are not equal")
	}
	if c1.chunkSize != c2.chunkSize {
		return nil, fmt.Errorf("The 2 chunk sizes are not equal")
	}

	// Reset both chunkers
	c1.Reset()
	c2.Reset()

	// Roll through the 2 files and find unequal hashes
	var diffs []int
	for i := 0; i < c1.fileSize-c1.chunkSize; i++ {

		if c1.hash != c2.hash {
			// The indexes of unequal chunks are stored in diffs
			diffs = append(diffs, c1.index)
			fmt.Println(c1)
			fmt.Println(c2)
			fmt.Println()
		}

		// Roll through file  1
		err := c1.Roll()
		if err != nil {
			return nil, err
		}

		// Roll through file 2
		err = c2.Roll()
		if err != nil {
			return nil, err
		}
	}

	// Return the indexes of unequel chunks
	return diffs, nil
}
