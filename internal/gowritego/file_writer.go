package gowritego
import (
	"os"

)
func BlocksToMarkdown(blocks[] Block) string {
	markdown := ""

	for _, block := range blocks {

		switch block.TypeOfBlock {
			case Markdown:
				markdown = markdown + block.Content + "\n"
			case Code:
				markdown = markdown + "```\n" + block.Content + "\n```\n"
		}
	}

	return markdown
}

func WriteFile(filename string, contents string) (int, error) {
	if fileInfo, _:= os.Stat(filename); fileInfo != nil {
		return 0, os.ErrExist
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.WriteString(contents)
	if err != nil {
		panic(err)
	}
	return n, nil
}
