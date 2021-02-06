package parser

import (
	"github.com/PaulioRandall/daft-wullie-go/node"
	"github.com/PaulioRandall/daft-wullie-go/token"
)

func Parse(lines [][]token.Lexeme) []node.Node {
	r := newLineReader(lines)
	return parseLines(r)
}

func parseLines(r *lineReader) []node.Node {
	panic("Not implemented yet!")
}
