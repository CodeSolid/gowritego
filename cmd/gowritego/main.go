/*#

# main.go

This comment will turn into markdown.

#*/
package main

import (
	"fmt"
	"github.com/CodeSolid/gowritego/internal/gowritego"
	"github.com/spf13/viper"
	"os"

)

// Currently not needed
func loadViperConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.Get("greeting"))

}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		fmt.Printf("Usage %v inputFile (Go) outputFile (markdown)\n", os.Args[0])
		os.Exit(-1)
	}
	input := os.Args[1]
	output := os.Args[2]

	fmt.Printf("Processing %v to %v\n", input, output)

	blocks, err := gowritego.ReadFile(input)
	if err != nil {
		panic(err)
	}

	markdown := gowritego.BlocksToMarkdown(blocks)
	_, err = gowritego.WriteFile(output, markdown)
	if err != nil {
		panic(err)
	}
}
