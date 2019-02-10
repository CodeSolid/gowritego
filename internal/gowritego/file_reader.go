package gowritego

import (
	"bufio"
	"os"
	"strings"
)

type BlockType int
const (
	Code BlockType = 0
	Markdown BlockType = 1
)

type Block struct {
	TypeOfBlock BlockType
	Content     string
}

type State int
const (
	BeginFile State = 0
	InsideMarkdown = 1
	MarkdownEnded = 2
	InsideCode = 3

)

func ReadFile(filename string) ([]Block, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	blocks 	:= make([]Block, 0)
	var state = BeginFile

	for scanner.Scan() {
		var text = scanner.Text()
		if strings.Index(text, "/*#") == 0 {
			blocks = append(blocks, Block{Markdown, ""})
			state = InsideMarkdown
			continue
		}
		if strings.Index(text, "#*/") == 0 {
			state = MarkdownEnded
			continue
		}
		switch(state) {
		case BeginFile:
			state = InsideCode
			blocks = append(blocks, Block{Code, text + "\n"})
		case MarkdownEnded:
			state = InsideCode
			blocks = append(blocks, Block{Code, text + "\n"})
		default: // Case InsideMarkdown or insideCode
			if len(blocks) > 0 {
				blocks[len(blocks)-1].Content += (text + "\n")
			}
		}





	}


	//fmt.Printf("File length:  %v\n", file_length)


	return blocks, nil
}


