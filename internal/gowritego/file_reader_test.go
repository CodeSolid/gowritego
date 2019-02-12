package gowritego

import (
	"testing"
	"os"
)


func TestReadFileCanReadTestFile(t  *testing.T) {
	blocks, _:= ReadFile(filename)
	if len(blocks) != 2 {
		t.Errorf("For %v, expected %v blocks, actual was %v\n", filename, 2, len(blocks))
	}
}

func TestNonExistentFilenameCreatesError(t  *testing.T) {
	_, err := ReadFile("bogus_file.go")
	if ! os.IsNotExist(err) {
		t.Errorf("Expected ErrNotExist, got %v", err)
	}
}


