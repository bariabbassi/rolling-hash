package main

import (
	"fmt"
)

type Chunk struct {
	index     int
	chunk     []byte
	file      []byte
	chunkSize int
	fileSize  int
	hash      int
}

func New(file []byte, csize int) *Chunk {
	return &Chunk{
		index:     0,
		chunk:     file[:csize],
		file:      file,
		chunkSize: csize,
		fileSize:  len(file),
		hash:      newHash(file[:csize]),
	}
}

func newHash(chunk []byte) int {
	hash := 0
	for i := 0; i < len(chunk); i++ {
		hash += int(chunk[i]) - 96
	}
	return hash
}

func (c *Chunk) String() string {
	return fmt.Sprintf("%d %s %d", c.index, string(c.chunk), c.hash)
}

func (c *Chunk) Roll() error {
	if c.index >= c.fileSize-c.chunkSize-1 {
		return fmt.Errorf("Can't roll. End of file reached")
	}
	newByte := c.file[c.index+chunkSize]
	c.hash = c.hash + int(newByte) - int(c.chunk[0])
	c.chunk = c.chunk[1:]
	c.chunk = append(c.chunk, newByte)
	c.index++
	return nil
}

func (c1 *Chunk) Diff(c2 *Chunk) (bool, []int, error) {
	if c1.fileSize != c2.fileSize {
		return false, nil, fmt.Errorf("The 2 file sizes are not equal")
	}
	if c1.chunkSize != c2.chunkSize {
		return false, nil, fmt.Errorf("The 2 chunk sizes are not equal")
	}
	var diffs []int
	for i := 0; i < c1.fileSize-c1.chunkSize-1; i++ {
		if c1.hash != c2.hash {
			diffs = append(diffs, c1.index)
		}
		err := c1.Roll()
		if err != nil {
			return false, nil, err
		}
		err = c2.Roll()
		if err != nil {
			return false, nil, err
		}
	}
	return true, diffs, nil
}
