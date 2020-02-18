package main

import "fmt"

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
	return fmt.Sprintf("%s %d", string(c.chunk), c.hash)
}

func (c *Chunk) Roll() error {
	if c.index >= c.fileSize-c.chunkSize-1 {
		return fmt.Errorf("Can't roll. No more  chunks left")
	}
	newByte := c.file[c.index+chunkSize]
	c.hash = c.hash + int(newByte) - int(c.chunk[0])
	c.chunk = c.chunk[1:]
	c.chunk = append(c.chunk, newByte)
	c.index += 1
	return nil
}
