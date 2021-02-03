package parser

import (
	"github.com/PaulioRandall/daft-wullie-go/types"
)

type RuneReader interface {
	More() bool             // True if more runes remain
	Match(int, string) bool // True if the string was matched
	Accept(string) bool     // True if the string was matched and sliced off
	MatchNewline() bool     // True if the next symbol represents a newline
	AcceptNewline() bool    // MatchNewline() + slices off the newline
	Drain() []rune          // Slices off all remaining runes
	Read(int) []rune        // Slices off the amount of runes
	ReadLine() RuneReader   // Slices the current line of runes, returns it as a RuneReader
}

func Parse(rr RuneReader) types.Notes {
	return parseNotes(rr)
}

func matchAny(rr RuneReader, pats ...string) bool {
	for _, pat := range pats {
		if rr.Match(0, pat) {
			return true
		}
	}
	return false
}

func parseNotes(rr RuneReader) types.Notes {
	n := types.Notes{Lines: []types.Phrase{}}

	for rr.More() {
		if rr.AcceptNewline() {
			n.Lines = append(n.Lines, types.EmptyLine{}) // Add one empty line only
			for rr.AcceptNewline() {                     // Discard consecutive empty lines
			}
			continue
		}

		p := parseLine(rr)
		n.Lines = append(n.Lines, p)

		if !rr.AcceptNewline() {
			panic("Expected newline")
		}
	}

	return n
}

func parseLine(rr RuneReader) types.Phrase {
	switch {
	case !rr.More():
		panic("Unexpected end of file")

	case rr.Accept("###"):
		return types.SubTopic{Text: parseTextLine(rr)}

	case rr.Accept("##"):
		return types.Topic{Text: parseTextLine(rr)}

	case rr.Accept("#"):
		return types.Title{Text: parseTextLine(rr)}

	case rr.Accept("."):
		return types.BulletItem{Node: parseNodeLine(rr)}

	case matchAny(rr, "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"):
		return parseNumItem(rr)

	case rr.Match(0, `"`):
		return parseQuote(rr)

	default:
		return parsePhrase(rr)
	}
}

func parsePhrase(rr RuneReader) types.Phrase {
	switch {
	case !rr.More():
		panic("Unexpected end of file")

	case rr.Accept("/"):
		return parseText(rr)

	case rr.Accept("+"):
		return types.Positive{Node: parseNodeUntil(rr, "+")}

	case rr.Accept("-"):
		return types.Negative{Node: parseNodeUntil(rr, "-")}

	case rr.Accept("*"):
		return types.Strong{Node: parseNodeUntil(rr, "*")}

	case rr.Accept("{"):
		return types.KeyPhrase{Text: parseTextUntil(rr, "}")}

	case rr.Accept("`"):
		return types.Snippet{Text: parseTextUntil(rr, "`")}

	case rr.Accept("&"):
		return parsePlace(rr)

	case rr.Accept("@"):
		return parseTime(rr)

	default:
		return parseText(rr)
	}
}

func parseText(rr RuneReader) types.Text {
	panic("Not implemented yet")
}

func parseTextLine(rr RuneReader) types.Text {
	panic("Not implemented yet")
}

func parseNode(rr RuneReader) types.Node {
	panic("Not implemented yet")
}

func parseNodeLine(rr RuneReader) types.Node {
	panic("Not implemented yet")
}

func parseNumItem(rr RuneReader) types.Node {
	panic("Not implemented yet")
}

func parseTextUntil(rr RuneReader, endDelim string) types.Text {
	panic("Not implemented yet")
}

func parseNodeUntil(rr RuneReader, endDelim string) types.Node {
	panic("Not implemented yet")
}

func parseQuote(rr RuneReader) types.Quote {
	panic("Not implemented yet")
}

func parsePlace(rr RuneReader) types.Place {
	panic("Not implemented yet")
}

func parseTime(rr RuneReader) types.Time {
	panic("Not implemented yet")
}

func maybeQuestion(rr RuneReader, left types.Phrase) types.Phrase {
	panic("Not implemented yet")
}
