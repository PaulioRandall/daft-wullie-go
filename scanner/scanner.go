package scanner

import (
	"strings"

	"github.com/PaulioRandall/daft-wullie-go/token"
)

type ScanLine func() ([]token.Lexeme, ScanLine)

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

func NewScanner(s string) ScanLine {
	ss := &scriptScanner{
		lines: splitLines(s),
	}
	return scanner(ss)
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
