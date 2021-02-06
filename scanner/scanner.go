package scanner

import (
	"strings"

	"github.com/PaulioRandall/daft-wullie-go/token"
)

type ScanLine func() ([]token.Lexeme, ScanLine)

func NewScanner(s string) ScanLine {
	ss := &scriptScanner{
		lines: splitLines(s),
	}
	return scanner(ss)
}

func ScanAll(s string) [][]token.Lexeme {

	var (
		f   = NewScanner(s)
		r   = [][]token.Lexeme{}
		lxs []token.Lexeme
	)

	for f != nil {
		lxs, f = f()
		r = append(r, lxs)
	}

	return r
}

func scanner(ss *scriptScanner) ScanLine {
	return func() ([]token.Lexeme, ScanLine) {
		tks := ss.scanLine()
		if ss.more() {
			return tks, scanner(ss)
		}
		return tks, nil
	}
}

func splitLines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(s, "\n")
}

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
