package scanner

import (
	"strings"
)

type RuneReader interface {
	More() bool             // True if more runes remain
	InRange(int) bool       // True if the index is within range
	Match(int, string) bool // True if the string was matched
	Accept(string) bool     // True if the string was matched and sliced off
	MatchNewline() bool     // True if the next symbol represents a newline
	AcceptNewline() bool    // MatchNewline() + slices off the newline
	Drain() []rune          // Slices off all remaining runes
	Read() rune             // Slices off the amount of runes
	ReadMany(int) []rune    // Slices off some amount of runes
	ReadLine() []rune       // Slices off the current line of runes
}

type TokenType int

const (
	UNDEFINED TokenType = iota
	TITLE
	TOPIC
	SUB_TOPIC
	BUL_POINT
	NUM_POINT
	QUESTION
	QUOTE
	PLUS
	MINUS
	ASTERISK
	L_BRACE
	R_BRACE
	BACK_TICK
	AMPERSAND
	AT
)

type Pos struct{ Line, Start, End int }

func (p Pos) GetLine() int  { return p.Line }
func (p Pos) GetStart() int { return p.Start }
func (p Pos) GetEnd() int   { return p.End }

type Token struct {
	TokenType
	Pos
	Val string
}

type ScanLine func() ([]Token, ScanLine, error)

func NewScanLine(s string) ScanLine {

	s = strings.ReplaceAll(s, "\r\n", "\n")
	sc := &scanner{data: strings.Split(s, "\n")}
	return sc.scan
}

func (sc *scanner) scan() ([]Token, ScanLine, error) {

	tks, e := sc.scanLine()
	if e != nil {
		return nil, nil, e
	}

	if sc.empty() {
		return tks, nil, nil
	}
	return tks, sc.scan, nil
}

type scanner struct {
	line int
	data []string
	col  int
	curr []rune
}

func (sc *scanner) scanLine() ([]Token, error) {
	sc.line++
	sc.col = 0
	sc.curr = []rune(sc.data[0])
	sc.data = sc.data[1:]

	return nil, nil
}

func (sc *scanner) empty() bool {
	return len(sc.data) == 0
}
