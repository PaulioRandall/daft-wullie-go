package parser

import (
	"github.com/PaulioRandall/daft-wullie-go/types"
)

type RuneReader interface {
	More() bool
	Accept(string) bool
	MatchNewline() bool
	AcceptNewline() bool
	Drain() []rune
	Read(int) []rune
	ReadLine() RuneReader
}

func Parse(rr RuneReader) types.Notes {

	panic("Not implemented yet")
}

/*
func parseHeading(rr RuneReader) types.Phrase {
	panic("Not implemented yet")
}

func parseNotes(rr RuneReader) types.Phrase {
	panic("Not implemented yet")
}

func parseLine(rr RuneReader) types.Phrase {
	panic("Not implemented yet")
}

func parseNode(rr RuneReader) types.NodePhrase {
	panic("Not implemented yet")
}

func parsePhrase(rr RuneReader) types.Phrase {
	panic("Not implemented yet")
}

func parseTime(rr RuneReader) types.TimePhrase {
	panic("Not implemented yet")
}
*/
