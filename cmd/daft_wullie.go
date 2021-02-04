package main

import (
	"github.com/PaulioRandall/daft-wullie-go/cmd/reader"
	"github.com/PaulioRandall/daft-wullie-go/parser"
	"github.com/PaulioRandall/daft-wullie-go/types"
)

const testData = `
# Title
## Topic
### Sub-topic
`

func main() {
	parse(testData)
}

func parse(s string) {
	rr := reader.NewReader(s)
	notes := parser.Parse(rr)
	println(types.DebugNotesString(notes))
}
