package parser

import (
	//"github.com/PaulioRandall/daft-wullie-go/types"
	"github.com/PaulioRandall/daft-wullie-go/token"
)

func Parse(lines [][]token.Lexeme) []Node {
	r := newLineReader(lines)
	return parseLines(r)
}

func parseLines(r *lineReader) []Node {
	panic("Not implemented yet!")
}
