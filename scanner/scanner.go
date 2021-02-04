package scanner

import (
	"strings"
)

type ScanLine func() ([]Lexeme, ScanLine)

func ScanAll(s string) [][]Lexeme {

	var (
		f   = NewScanner(s)
		r   = [][]Lexeme{}
		lxs []Lexeme
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
	return func() ([]Lexeme, ScanLine) {
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
