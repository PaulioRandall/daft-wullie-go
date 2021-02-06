package parser

import (
	"github.com/PaulioRandall/daft-wullie-go/node"
	"github.com/PaulioRandall/daft-wullie-go/token"
)

type ParseLine func() (node.Node, ParseLine)

func NewParser(lines [][]token.Lexeme) ParseLine {
	r := &lineReader{lines: lines}
	if !r.more() {
		return nil
	}
	return parser(r)
}

func ParseAll(lines [][]token.Lexeme) []node.Node {
	var (
		f = NewParser(lines)
		r = []node.Node{}
		n node.Node
	)
	for f != nil {
		n, f = f()
		r = append(r, n)
	}
	return r
}

func parser(r *lineReader) ParseLine {
	return func() (node.Node, ParseLine) {
		lr := r.nextLine()
		ns := parseLine(lr)
		if r.more() {
			return ns, parser(r)
		}
		return ns, nil
	}
}

func parseLine(r *tokenReader) node.Node {
	switch {
	case r.accept(token.H1):
		return node.H1{M_Text: parsePhraseLine(r)}

	case r.accept(token.H2):
		return node.H2{M_Text: parsePhraseLine(r)}

	case r.accept(token.H3):
		return node.H3{M_Text: parsePhraseLine(r)}

	case r.accept(token.QUOTE):
		return node.Quote{M_Text: parsePhraseLine(r)}

	default:
		panic("Unknown token")
	}
}

// TEXT_LINE := TEXT
func parsePhraseLine(r *tokenReader) string {
	if !r.match(token.TEXT) {
		panic("Expected TEXT token")
	}
	return r.read().Val
}

//func parse...(r *tokenReader) node.??? {
//panic("Not implemented yet!")
//}
