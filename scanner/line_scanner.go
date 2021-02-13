package scanner

import (
	"unicode"

	"github.com/PaulioRandall/daft-wullie-go/token"
)

type lineScanner struct{ text []rune }

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var key_symbols = []string{"\\", "+", "-", "*", "`"}
var key_tokens = []struct {
	sym string
	tk  token.Token
}{
	{"\\", token.ESCAPE},
	{"**", token.KEY_PHRASE},
	{"+", token.POSITIVE},
	{"-", token.NEGATIVE},
	{"*", token.STRONG},
	{"`", token.SNIPPET},
}

func (ls *lineScanner) scanLine() []token.Lexeme {

	ls.discardSpace()

	switch {
	case ls.matchStr("###"):
		return []token.Lexeme{ls.slice(token.H3, 3), ls.scanTextLine()}

	case ls.matchStr("##"):
		return []token.Lexeme{ls.slice(token.H2, 2), ls.scanTextLine()}

	case ls.matchStr("#"):
		return []token.Lexeme{ls.slice(token.H1, 1), ls.scanTextLine()}

	case ls.matchStr(">"):
		r := []token.Lexeme{ls.slice(token.QUOTE, 1)}
		return append(r, ls.scanNodes()...)

	case ls.matchStr("."):
		r := []token.Lexeme{ls.slice(token.BUL_POINT, 1)}
		return append(r, ls.scanNodes()...)

	case ls.matchNumPoint():
		return ls.scanNumPoint()

	default:
		return ls.scanNodes()
	}
}

func (ls *lineScanner) scanNodes() []token.Lexeme {
	r := []token.Lexeme{}
	for ls.inRange(0) {
		lx := ls.scanNode()
		r = append(r, lx)
	}
	return normalise(r)
}

func (ls *lineScanner) scanNode() token.Lexeme {
	for _, v := range key_tokens {
		if ls.matchStr(v.sym) {
			return ls.slice(v.tk, len(v.sym))
		}
	}
	return ls.scanText()
}

func (ls *lineScanner) scanTextLine() token.Lexeme {
	return ls.sliceBy(token.TEXT, anyMatcher)
}

func (ls *lineScanner) scanText() token.Lexeme {
	return ls.sliceBy(token.TEXT, nonKeyMatcher)
}

func (ls *lineScanner) scanNumPoint() []token.Lexeme {
	r := []token.Lexeme{ls.sliceBy(token.NUM_POINT, newNumPointMatcher())}
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
	ls.sliceBy(token.UNDEFINED, spaceMatcher)
}

func (ls *lineScanner) slice(tk token.Token, n int) token.Lexeme {
	val := ls.text[:n]
	ls.text = ls.text[n:]
	return token.Lexeme{
		Token: tk,
		Val:   string(val),
	}
}

func (ls *lineScanner) sliceBy(tk token.Token, f func(rune) bool) token.Lexeme {
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

func normalise(lxs []token.Lexeme) []token.Lexeme {
	if len(lxs) == 0 {
		return lxs
	}
	lxs = applyEscaping(lxs)
	return mergeLexemes(lxs)
}

// applyEscaping converts non-text tokens into text ones if they follow an
// escape token.
//
// The following are some experimental documentation formats:
//
// Descriptive list definition of behaviour:
// - input must not be empty
// - the symbol immediately after an escape token is always converted to text
// - a '\\' will be converted to the text '\'
// - all escape symbols are discarded except escaped escape symbols
// - a trailing '\' in the input will be discarded
//
// Axiomatic definition of behaviour:
// - ANY := non-ESCAPE token
// - ESCAPE ANY    -> TEXT(ANY)
// - ESCAPE ESCAPE -> TEXT(ESCAPE)
// - ESCAPE EOF    -> EOF
func applyEscaping(in []token.Lexeme) []token.Lexeme {

	size := len(in)
	out := make([]token.Lexeme, 0, size)

	for i := 0; i < size; i++ {
		tk := in[i]

		if tk.Token != token.ESCAPE {
			out = append(out, tk)
			continue
		}

		i++
		if i < size {
			tk = in[i]
			tk.Token = token.TEXT
			out = append(out, tk)
		}
	}

	return out
}

// mergeLexemes merges lexemes where possible, i.e. merging sections of text
// that appear next to each other in the input. Assumes the input is not empty.
func mergeLexemes(in []token.Lexeme) []token.Lexeme {

	size := len(in)
	out := make([]token.Lexeme, 0, size)
	tryMerge := func(lx token.Lexeme) bool {
		last := len(out) - 1
		if lx.Token == token.TEXT && out[last].Token == token.TEXT {
			out[last].Val += lx.Val
			return true
		}
		return false
	}

	out = append(out, in[0])
	for i := 1; i < size; i++ {
		lx := in[i]
		if !tryMerge(lx) {
			out = append(out, lx)
		}
	}

	return out
}
