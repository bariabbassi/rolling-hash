package main

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	fileName := "testfiles/alphabet1.txt"

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}
	if string(file) != "abcdefghijklmnopqrstuvwxyz" {
		t.Error("File read incorrectly")
	}
}

func TestNewHash(t *testing.T) {
	if newHash([]byte("")) != 0 {
		t.Error("Incorrect hash")
	}
	if newHash([]byte("x")) != 24 {
		t.Error("Incorrect hash", newHash([]byte("x")))
	}
	if newHash([]byte("abcdefghijklmnopqrstuvwxyz")) != 351 {
		t.Error("Incorrect hash")
	}
	if newHash([]byte("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")) != 702 {
		t.Error("Incorrect hash")
	}
	if newHash([]byte("Rolling Hash!!")) != -131 {
		t.Error("Incorrect hash")
	}
}

func TestNew(t *testing.T) {
	fileName := "testfiles/alphabet1.txt"
	chunkSize := 4

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}
	if string(file) != "abcdefghijklmnopqrstuvwxyz" {
		t.Error("File read incorrectly")
	}

	chunker := New(file, chunkSize)
	if chunker.index != 0 {
		t.Error("Incorrect index")
	}
	if string(chunker.chunk) != "abcd" {
		t.Error("Incorrect chunk")
	}
	if string(chunker.file) != "abcdefghijklmnopqrstuvwxyz" {
		t.Error("Incorrect file")
	}
	if chunker.chunkSize != chunkSize {
		t.Error("Incorrect chunkSize")
	}
	if chunker.fileSize != len(file) {
		t.Error("Incorrect fileSize")
	}
	if chunker.hash != 10 {
		t.Error("Incorrect hash")
	}
}

func TestRoll(t *testing.T) {
	fileName := "testfiles/alphabet1.txt"
	chunkSize := 4

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}

	// Check the first
	chunker := New(file, chunkSize)
	if string(chunker.chunk) != "abcd" {
		t.Error("Incorrect chunk")
	}

	// Check the second chunk
	err = chunker.Roll()
	if err != nil {
		t.Error(err)
	}
	if string(chunker.chunk) != "bcde" {
		t.Error("Incorrect chunk")
	}

	// check the chunk before the last
	for i := 0; i < 20; i++ {
		err = chunker.Roll()
		if err != nil {
			t.Error(err)
		}
	}
	if string(chunker.chunk) != "vwxy" {
		t.Error("Incorrect chunk")
	}

	// Check the last chunk
	err = chunker.Roll()
	if err != nil {
		t.Error(err)
	}
	if string(chunker.chunk) != "wxyz" {
		t.Error("Incorrect chunk")
	}

	// The chunk should stay at wxyz
	err = chunker.Roll()
	if err == nil {
		t.Error("Should give an end of file error")
	}
	if string(chunker.chunk) != "wxyz" {
		t.Error("Incorrect chunk")
	}
}

func TestRoll2(t *testing.T) {
	fileName := "testfiles/alphabet1.txt"
	chunkSize := 5

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}

	// Check the first
	chunker := New(file, chunkSize)
	if string(chunker.chunk) != "abcde" {
		t.Error("Incorrect chunk")
	}

	// Check the second chunk
	err = chunker.Roll()
	if err != nil {
		t.Error(err)
	}
	if string(chunker.chunk) != "bcdef" {
		t.Error("Incorrect chunk")
	}

	// check the chunk before the last
	for i := 0; i < 19; i++ {
		err = chunker.Roll()
		if err != nil {
			t.Error(err)
		}
	}
	if string(chunker.chunk) != "uvwxy" {
		t.Error("Incorrect chunk")
	}

	// Check the last chunk
	err = chunker.Roll()
	if err != nil {
		t.Error(err)
	}
	if string(chunker.chunk) != "vwxyz" {
		t.Error("Incorrect chunk")
	}

	// The chunk should stay at wxyz
	err = chunker.Roll()
	if err == nil {
		t.Error("Should give an end of file error")
	}
	if string(chunker.chunk) != "vwxyz" {
		t.Error("Incorrect chunk")
	}
}

func TestRoll3(t *testing.T) {
	fileName := "testfiles/alphabet1.txt"
	chunkSize := 8

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}

	// Check the first
	chunker := New(file, chunkSize)
	if string(chunker.chunk) != "abcdefgh" {
		t.Error("Incorrect chunk")
	}

	// Check the second chunk
	err = chunker.Roll()
	if err != nil {
		t.Error(err)
	}
	if string(chunker.chunk) != "bcdefghi" {
		t.Error("Incorrect chunk")
	}

	// check the chunk before the last
	for i := 0; i < 16; i++ {
		err = chunker.Roll()
		if err != nil {
			t.Error(err)
		}
	}
	if string(chunker.chunk) != "rstuvwxy" {
		t.Error("Incorrect chunk")
	}

	// Check the last chunk
	err = chunker.Roll()
	if err != nil {
		t.Error(err)
	}
	if string(chunker.chunk) != "stuvwxyz" {
		t.Error("Incorrect chunk")
	}

	// The chunk should stay at wxyz
	err = chunker.Roll()
	if err == nil {
		t.Error("Should give an end of file error")
	}
	if string(chunker.chunk) != "stuvwxyz" {
		t.Error("Incorrect chunk")
	}
}

func TestDiff(t *testing.T) {
	fileName1 := "testfiles/alphabet1.txt"
	fileName2 := "testfiles/alphabet2.txt"
	fileName3 := "testfiles/alphabet3.txt"
	chunkSize := 4

	file1, err := ioutil.ReadFile(fileName1)
	file2, err := ioutil.ReadFile(fileName2)
	file3, err := ioutil.ReadFile(fileName3)

	chunker1 := New(file1, chunkSize)
	chunker1p := New(file1, chunkSize)
	chunker2 := New(file2, chunkSize)
	chunker3 := New(file3, chunkSize)

	// Diff 2 equal files
	diffs, err := chunker1.Diff(chunker1p)
	if err != nil {
		t.Error(err)
	}
	if diffs != nil {
		t.Error("Incorrect diff")
	}

	// not the same size
	diffs, err = chunker1.Diff(chunker2)
	if err == nil {
		t.Error("Incorrect diff")
	}
	if diffs != nil {
		t.Error("Incorrect diffs", diffs)
	}

	// 2 lettres
	diffs, err = chunker1.Diff(chunker3)
	if err != nil {
		t.Error(err)
	}
	if len(diffs) != 4 || diffs[0] != 16 || diffs[1] != 17 || diffs[2] != 18 || diffs[3] != 19 {
		t.Error("Incorrect diffs", diffs)
	}

}

func TestDiff2(t *testing.T) {
	fileName1 := "testfiles/alphabet1.txt"
	fileName2 := "testfiles/alphabet2.txt"
	fileName3 := "testfiles/alphabet3.txt"
	chunkSize := 5

	file1, err := ioutil.ReadFile(fileName1)
	file2, err := ioutil.ReadFile(fileName2)
	file3, err := ioutil.ReadFile(fileName3)

	chunker1 := New(file1, chunkSize)
	chunker1p := New(file1, chunkSize)
	chunker2 := New(file2, chunkSize)
	chunker3 := New(file3, chunkSize)

	// Diff 2 equal files
	diffs, err := chunker1.Diff(chunker1p)
	if err != nil {
		t.Error(err)
	}
	if diffs != nil {
		t.Error("Incorrect diff")
	}

	// not the same size
	diffs, err = chunker1.Diff(chunker2)
	if err == nil {
		t.Error("Incorrect diff")
	}
	if diffs != nil {
		t.Error("Incorrect diffs", diffs)
	}

	// 2 lettres
	diffs, err = chunker1.Diff(chunker3)
	if err != nil {
		t.Error(err)
	}
	if len(diffs) != 5 || diffs[0] != 15 || diffs[1] != 16 || diffs[2] != 17 || diffs[3] != 18 || diffs[4] != 19 {
		t.Error("Incorrect diffs", diffs)
	}
}

func TestDiff3(t *testing.T) {
	fileName1 := "testfiles/longtext1.txt"
	fileName2 := "testfiles/longtext2.txt"
	fileName3 := "testfiles/longtext3.txt"
	chunkSize := 16

	file1, err := ioutil.ReadFile(fileName1)
	file2, err := ioutil.ReadFile(fileName2)
	file3, err := ioutil.ReadFile(fileName3)

	chunker1 := New(file1, chunkSize)
	chunker1p := New(file1, chunkSize)
	chunker2 := New(file2, chunkSize)
	chunker3 := New(file3, chunkSize)

	// Diff 2 equal files
	diffs, err := chunker1.Diff(chunker1p)
	if err != nil {
		t.Error(err)
	}
	if diffs != nil {
		t.Error("Incorrect diff")
	}

	// not the same size
	diffs, err = chunker1.Diff(chunker2)
	if err == nil {
		t.Error("Incorrect diff")
	}
	if diffs != nil {
		t.Error("Incorrect diffs", diffs)
	}

	// 2 lettres
	diffs, err = chunker1.Diff(chunker3)
	if err != nil {
		t.Error(err)
	}
	if len(diffs) != 16 {
		t.Error("Incorrect diffs", diffs)
	}
	for i := 0; i < 16; i++ {
		if diffs[i] != 338+i {
			t.Error("Incorrect diffs value ", diffs[i])
		}
	}
}
