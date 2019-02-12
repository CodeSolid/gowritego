package gowritego

import (
	"os"
	"testing"
	"strings"
)

const filename = "./testdata/hello_world_test.go"

func TestBlocksToMarkdown(t *testing.T) {

	blocks, err := ReadFile(filename)
	if err != nil {
		t.Errorf("Could not read file %v, %v",filename, err)
	}

	markdown := BlocksToMarkdown(blocks)
	if (len( markdown) == 0)  {
		t.Errorf("No output from markdown generated for %v",filename)
	}

	markers := strings.Count(markdown, "```")
	if markers != 2 {
		t.Errorf("Wrong number of code block markers for %v; expected 2, got %v",filename, markers)
	}

	trimmed := strings.Trim(markdown, "\n\t ")
	if strings.Index(trimmed, "```") == 0 {
		t.Errorf("Expected %v to not start with a code block",filename)
	}

	if ! strings.HasSuffix(trimmed, "```") {
		t.Errorf("Expected %v to end with a code block",filename)
	}
}

func TestWriteFileReturnsErrExistIfFileExists(t *testing.T) {
	_, err := WriteFile(filename, "Unused")
	if (! os.IsExist(err)) {
		t.Errorf("Expected WriteFile to return %v for file %v", os.ErrExist, filename)
	}
}


