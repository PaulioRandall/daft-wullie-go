package reader

import (
	"github.com/PaulioRandall/daft-wullie-go/parser"
)

type runeReader struct {
	runes []rune
}

func NewReader(s string) parser.RuneReader {
	return &runeReader{runes: []rune(s)}
}

func (rr *runeReader) More() bool {
	return len(rr.runes) > 0
}

func (rr *runeReader) match(start int, s string) bool {

	needle := []rune(s)
	if len(rr.runes) < len(needle) {
		return false
	}

	for i, ru := range needle {
		if rr.runes[i] != ru {
			return false
		}
	}

	return true
}

func (rr *runeReader) Accept(s string) bool {
	if rr.match(0, s) {
		rr.Read(len([]rune(s)))
		return true
	}
	return false
}

func (rr *runeReader) MatchNewline() bool {
	return rr.match(0, "\n") || rr.match(0, "\r\n")
}

func (rr *runeReader) AcceptNewline() bool {
	return rr.Accept("\n") || rr.Accept("\r\n")
}

func (rr *runeReader) Drain() []rune {
	return rr.Read(len(rr.runes))
}

func (rr *runeReader) Read(n int) []rune {

	if len(rr.runes) < n {
		panic("Index out of range, reading too many runes")
	}

	if n == 0 {
		return []rune{}
	}

	r := rr.runes[:n]
	rr.runes = rr.runes[n:]
	return r
}

func (rr *runeReader) ReadLine() parser.RuneReader {

	i := 0

	for i = range rr.runes {

		if rr.match(i, "\n") {
			defer rr.Read(1)
			break
		}

		if rr.match(i, "\r\n") {
			defer rr.Read(2)
			break
		}
	}

	return &runeReader{runes: rr.Read(i)}
}
