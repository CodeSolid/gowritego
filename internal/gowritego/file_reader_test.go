package gowritego

import (
	"fmt"
)

func ExampleReadFileCanBeCalled() {
	blocks, err := ReadFile("./file_reader_test.go")
	fmt.Printf("%v\n", err == nil)
	fmt.Printf("%v\n", len(blocks))
	// Output:
	// true
	// 1
}

func ExampleNonExistentFilenameCreatesError() {
	_, err := ReadFile("./somefile.go")
	fmt.Printf("%v\n", err != nil)
	fmt.Printf("%v\n", err)
	// Output:
	// true
	// stat ./somefile.go: no such file or directory
}
