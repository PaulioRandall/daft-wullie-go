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

func parseNotes() types.Phrase {
	panic("Not implemented yet")
}

func parseLine() types.Phrase {
	panic("Not implemented yet")
}

func parseNode() types.NodePhrase {
	panic("Not implemented yet")
}

func parsePhrase() types.Phrase {
	panic("Not implemented yet")
}

func parseTime() types.TimePhrase {
	panic("Not implemented yet")
}
