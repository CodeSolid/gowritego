/*#

# main.go

This is the main file for gowritego, a simple tool that reads special comment blocks
like this in a golang program to allow.  For more information, see the code and the readme
at [https://github.com/CodeSolid/gowritego](https://github.com/CodeSolid/gowritego)

#*/
package main

import (
	"fmt"
	"github.com/CodeSolid/gowritego/internal/gowritego"
	"os"
)


func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		fmt.Printf("Usage %v inputFile (Go) outputFile (markdown)\n", os.Args[0])
		os.Exit(-1)
	}
	input := os.Args[1]
	output := os.Args[2]

	fmt.Printf("Processing %v to %v\n", input, output)
	// Read input file to blocks of either markdown or code
	blocks, err := gowritego.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Render and write the markdown
	markdown := gowritego.BlocksToMarkdown(blocks)
	_, err = gowritego.WriteFile(output, markdown)
	if err != nil {
		panic(err)
	}
}
