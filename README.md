# gowritego
A simple tool for publishing go files, gowritego looks for standard go / C-style multiline comment blocks
with an extra hash symbol like this in .go source:

```
/*#

...

#*/


```

Anything between those blocks is treated as markdown, and the comment delimiters are removed,
while the go source itself is rendered in a markdown code block.


****
## Usage

With go 1.11 or later and $GOPATH\bin on your system path:

```
go install github.com/CodeSolid/gowritego/cmd/gowritego
gowritego <golang_source_file> <markdown_input>
```

