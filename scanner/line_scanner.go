package scanner

import (
	"unicode"
)

type lineScanner struct{ text []rune }

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var key_symbols = []string{"\\", "+", "-", "*", "`", "?"}
var key_tokens = []struct {
	sym string
	tk  Token
}{
	{"\\", ESCAPE},
	{"**", KEY_PHRASE},
	{"+", POSITIVE},
	{"-", NEGATIVE},
	{"*", STRONG},
	{"`", SNIPPET},
	{"?", QUESTION},
}

func (ls *lineScanner) scanLine() []Lexeme {

	ls.discardSpace()

	switch {
	case ls.matchStr("###"):
		return []Lexeme{ls.slice(HEADING, 3), ls.scanTextLine()}

	case ls.matchStr("##"):
		return []Lexeme{ls.slice(SUB_TOPIC, 2), ls.scanTextLine()}

	case ls.matchStr("#"):
		return []Lexeme{ls.slice(TOPIC, 1), ls.scanTextLine()}

	case ls.matchStr(">"):
		r := []Lexeme{ls.slice(QUOTE, 1)}
		return append(r, ls.scanNodes()...)

	case ls.matchStr("."):
		r := []Lexeme{ls.slice(BUL_POINT, 1)}
		return append(r, ls.scanNodes()...)

	case ls.matchNumPoint():
		return ls.scanNumPoint()

	default:
		return ls.scanNodes()
	}
}

func (ls *lineScanner) scanNodes() []Lexeme {
	r := []Lexeme{}
	for ls.inRange(0) {
		lx := ls.scanNode()
		r = append(r, lx)
	}
	return r
}

func (ls *lineScanner) scanNode() Lexeme {
	for _, v := range key_tokens {
		if ls.matchStr(v.sym) {
			return ls.slice(v.tk, len(v.sym))
		}
	}
	return ls.scanText()
}

func (ls *lineScanner) scanTextLine() Lexeme {
	return ls.sliceBy(TEXT, anyMatcher)
}

func (ls *lineScanner) scanText() Lexeme {
	return ls.sliceBy(TEXT, nonKeyMatcher)
}

func (ls *lineScanner) scanNumPoint() []Lexeme {
	r := []Lexeme{ls.sliceBy(NUM_POINT, newNumPointMatcher())}
	return append(r, ls.scanNodes()...)
}

func (ls *lineScanner) at(i int) rune {
	if !ls.inRange(i) {
		panic("Index out of range")
	}
	return ls.text[i]
}

func (ls *lineScanner) inRange(i int) bool {
	return i < len(ls.text)
}

func (ls *lineScanner) match(ru rune) bool {
	return ls.inRange(0) && ls.at(0) == ru
}

func (ls *lineScanner) matchStr(s string) bool {
	for i, ru := range s {
		if !ls.inRange(i) || ls.at(i) != ru {
			return false
		}
	}
	return true
}

func (ls *lineScanner) matchAny(haystack ...string) bool {
	return ls.inRange(0) &&
		matchAny(string(ls.at(0)), haystack...)
}

func (ls *lineScanner) matchNumPoint() bool {
	i := 0
	for ; ls.inRange(i); i++ {
		if !matchAny(string(ls.at(i)), digits...) {
			break
		}
	}
	return i > 0 && ls.inRange(i) && ls.at(i) == '.'
}

func (ls *lineScanner) discardSpace() {
	ls.sliceBy(UNDEFINED, spaceMatcher)
}

func (ls *lineScanner) slice(tk Token, n int) Lexeme {
	val := ls.text[:n]
	ls.text = ls.text[n:]
	return Lexeme{
		Token: tk,
		Val:   string(val),
	}
}

func (ls *lineScanner) sliceBy(tk Token, f func(rune) bool) Lexeme {
	i := 0
	for ; ls.inRange(i) && f(ls.text[i]); i++ {
	}
	return ls.slice(tk, i)
}

func matchAny(needle string, haystack ...string) bool {
	for _, ru := range haystack {
		if ru == needle {
			return true
		}
	}
	return false
}

func anyMatcher(rune) bool {
	return true
}

func spaceMatcher(ru rune) bool {
	return unicode.IsSpace(ru)
}

func nonKeyMatcher(ru rune) bool {
	return !matchAny(string(ru), key_symbols...)
}

func newNumPointMatcher() func(rune) bool {
	done := false
	return func(ru rune) bool {
		if done {
			return false
		}
		if ru == '.' {
			done = true
		}
		return true
	}
}
