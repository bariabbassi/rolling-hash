package main

type Chunk struct {
	index uint
	chunk []byte
	file  []byte
}

func New(file []byte) *Chunk {
	return &Chunk{
		index: 0,
		chunk: file[:chunkSize],
		file:  file,
	}
}

func (c *Chunk) String() string {
	return string(c.chunk)
}

func (c *Chunk) Roll() {
	c.index += 1
	c.chunk = c.chunk[1:]
	c.chunk = append(c.chunk, c.file[c.index+chunkSize])
}
