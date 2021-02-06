package scanner

import (
	"github.com/PaulioRandall/daft-wullie-go/token"
)

type scriptScanner struct {
	idx   int
	lines []string
}

func (ss *scriptScanner) more() bool {
	return ss.idx < len(ss.lines)
}

func (ss *scriptScanner) scanLine() []token.Lexeme {
	ls := &lineScanner{
		text: []rune(ss.lines[ss.idx]),
	}
	ss.idx++
	return ls.scanLine()
}
