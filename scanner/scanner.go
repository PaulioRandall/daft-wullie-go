package scanner

import (
	//"errors"
	"strings"
	"unicode"
)

type (
	Token    int
	ScanLine func() ([]Lexeme, ScanLine, error)
	Pos      struct{ Line, Start, End int }

	Lexeme struct {
		Token
		Pos
		Val string
	}

	scanner struct {
		line int
		data []string
		col  int
		curr []rune
	}
)

const (
	UNDEFINED Token = iota
	TEXT
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

func NewScanLine(s string) ScanLine {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	sc := &scanner{data: strings.Split(s, "\n")}
	return sc.scan
}

func (p Pos) GetLine() int  { return p.Line }
func (p Pos) GetStart() int { return p.Start }
func (p Pos) GetEnd() int   { return p.End }

func (sc *scanner) scan() ([]Lexeme, ScanLine, error) {

	tks, e := sc.scanLine()
	if e != nil {
		return nil, nil, e
	}

	if sc.empty() {
		return tks, nil, nil
	}
	return tks, sc.scan, nil
}

func (sc *scanner) scanLine() ([]Lexeme, error) {

	sc.line++
	sc.col = 0
	sc.curr = []rune(sc.data[0])
	sc.data = sc.data[1:]

	sc.discardSpace()

	switch {
	case sc.match("#"):
		return sc.scanHeading()

	case sc.match("."):
		return sc.scanBulPoint()

	case sc.match(`"`):
		return sc.scanQuote()

	case sc.matchDigit(0):
		i := 1
		for ; sc.matchDigit(i); i++ {
		}
		if sc.matchAt(i, ".") {
			return sc.scanNumPoint()
		}
		fallthrough

	default:
		return sc.scanPhrases()
	}
}

func (sc *scanner) empty() bool {
	return len(sc.data) == 0
}

func (sc *scanner) discardSpace() {
	sc.sliceBy(UNDEFINED, func(i int, ru rune) bool {
		return unicode.IsSpace(ru)
	})
}

func (sc *scanner) scanTextLine(stopAtAny ...string) (Lexeme, error) {

	sc.discardSpace()

	i := 0
	for ; sc.inRange(i); i++ {
		if sc.match(`\`) {
			i++
			continue
		}
		if sc.matchAny(stopAtAny...) {
			break
		}
	}

	lx, e := sc.slice(TEXT, i)
	if e != nil {
		return Lexeme{}, e
	}

	v := []rune(lx.Val)
	if len(v) != 0 {
		for i := len(v) - 1; unicode.IsSpace(v[i]); i = len(v) - 1 {
			v = v[:i]
			lx.Pos.End--
		}
	}

	s := string(v)
	lx.Val = strings.ReplaceAll(s, `\`, "")
	return lx, nil
}

func (sc *scanner) scanHeading() ([]Lexeme, error) {

	var hashCount int
	var tk Token

	switch {
	case sc.match("###"):
		tk, hashCount = SUB_TOPIC, 3
	case sc.match("##"):
		tk, hashCount = TOPIC, 2
	case sc.match("#"):
		tk, hashCount = TITLE, 1
	default:
		panic("SANITY CHECK!")
	}

	var r [2]Lexeme
	var e error

	r[0], e = sc.slice(tk, hashCount)
	if e != nil {
		return nil, e
	}

	r[1], e = sc.scanTextLine("+", "-", "*", "`", "{", "&", "@")
	if e != nil {
		return nil, e
	}

	return r[:], nil
}

func (sc *scanner) scanBulPoint() ([]Lexeme, error) {
	panic("Not implemented yet")
}

func (sc *scanner) scanNumPoint() ([]Lexeme, error) {
	panic("Not implemented yet")
}

func (sc *scanner) scanQuote() ([]Lexeme, error) {
	panic("Not implemented yet")
}

func (sc *scanner) scanPhrases() ([]Lexeme, error) {
	panic("Not implemented yet")
	/*
			r := []Lexeme{}

		for len(sc.curr) != 0 {
			l, e := sc.scanNext(first)
			if e != nil {
				return e
			}
			r = append(r, l)
			return nil
		}

		return r, nil
	*/
}

func (sc *scanner) inRange(i int) bool {
	return i < len(sc.curr)
}

func (sc *scanner) at(i int) rune {
	panic("Not implemented yet: scanner.at")
}

func (sc *scanner) match(needle string) bool {
	return sc.matchAt(0, needle)
}

func (sc *scanner) matchAt(start int, needle string) bool {

	size := len(needle)
	if !sc.inRange(start + size) {
		return false
	}

	for i, ru := range needle {
		if ru != sc.curr[start+i] {
			return false
		}
	}

	return true
}

func (sc *scanner) matchDigit(i int) bool {
	// sc.matchAny("1", "2", "3", "4", "5", "6", "7", "8", "9")
	panic("Not implemented yet: scanner.matchDigit")
}

func (sc *scanner) matchAny(needles ...string) bool {
	for _, n := range needles {
		if sc.match(n) {
			return true
		}
	}
	return false
}

func (sc *scanner) slice(tk Token, n int) (Lexeme, error) {

	pos := Pos{
		Line:  sc.line,
		Start: sc.col,
		End:   sc.col + n,
	}

	val := sc.curr[:n]
	sc.curr = sc.curr[n:]
	sc.col += n

	lx := Lexeme{
		Token: tk,
		Pos:   pos,
		Val:   string(val),
	}

	return lx, nil
}

func (sc *scanner) sliceBy(tk Token, f func(int, rune) bool) (Lexeme, error) {
	i := 0
	for ; sc.inRange(i) && f(i, sc.curr[i]); i++ {
	}
	return sc.slice(tk, i)
}
