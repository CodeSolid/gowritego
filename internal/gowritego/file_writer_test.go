package gowritego

import (
	"os"
	"testing"
	"strings"
)

func TestBlocksToMarkdown(t *testing.T) {
	const FILENAME = "./hello_world_test.go"
	blocks, err := ReadFile(FILENAME)
	if err != nil {
		t.Errorf("Could not read file %v, %v",FILENAME, err)
	}

	markdown := BlocksToMarkdown(blocks)
	if (len( markdown) == 0)  {
		t.Errorf("No output from markdown generated for %v",FILENAME)
	}

	markers := strings.Count(markdown, "```")
	if markers != 2 {
		t.Errorf("Wrong number of code block markers for %v; expected 2, got %v",FILENAME, markers)
	}

	trimmed := strings.Trim(markdown, "\n\t ")
	if strings.Index(trimmed, "```") == 0 {
		t.Errorf("Expected %v to not start with a code block",FILENAME)
	}

	if ! strings.HasSuffix(trimmed, "```") {
		t.Errorf("Expected %v to end with a code block",FILENAME)
	}
}

func TestWriteFileReturnsErrExistIfFileExists(t *testing.T) {
	const FILENAME = "./hello_world_test.go"
	_, err := WriteFile(FILENAME, "Unused")
	if (! os.IsExist(err)) {
		t.Errorf("Expected WriteFile to return %v for file %v", os.ErrExist, FILENAME)
	}
}


func TestDirectoryDisplay(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log(wd)
}
