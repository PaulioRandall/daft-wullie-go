package scanner

import (
	"unicode"

	"github.com/PaulioRandall/daft-wullie-go/token"
)

type lineScanner struct{ text []rune }

func (ls *lineScanner) scanLine() []token.Lexeme {

	ls.discardSpace()

	switch {
	case ls.matchStr("##"):
		return []token.Lexeme{ls.slice(token.SubTopic, 2), ls.scanTextLine()}

	case ls.matchStr("#"):
		return []token.Lexeme{ls.slice(token.Topic, 1), ls.scanTextLine()}

	case ls.matchStr(".."):
		r := []token.Lexeme{ls.slice(token.SubBulPoint, 2)}
		return append(r, ls.scanNodes()...)

	case ls.matchStr("."):
		r := []token.Lexeme{ls.slice(token.BulPoint, 1)}
		return append(r, ls.scanNodes()...)

	case ls.matchStr("!!"):
		r := []token.Lexeme{ls.slice(token.SubNumPoint, 2)}
		return append(r, ls.scanNodes()...)

	case ls.matchStr("!"):
		r := []token.Lexeme{ls.slice(token.NumPoint, 1)}
		return append(r, ls.scanNodes()...)

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

	keyTokens := []struct {
		sym string
		tk  token.Token
	}{
		{"\\", token.Escape},
		{"**", token.KeyPhrase},
		{"+", token.Positive},
		{"-", token.Negative},
		{"*", token.Strong},
		{`"`, token.Quote},
		{"`", token.Snippet},
	}

	for _, v := range keyTokens {
		if ls.matchStr(v.sym) {
			return ls.slice(v.tk, len(v.sym))
		}
	}

	return ls.scanText()
}

func (ls *lineScanner) scanTextLine() token.Lexeme {
	return ls.sliceBy(token.Text, anyMatcher)
}

func (ls *lineScanner) scanText() token.Lexeme {
	return ls.sliceBy(token.Text, nonKeyMatcher)
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

func (ls *lineScanner) discardSpace() {
	ls.sliceBy(token.Undefined, spaceMatcher)
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
	return !matchAny(string(ru), "\\", "+", "-", "*", "`", `"`)
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
// Descriptive definition of behaviour:
// - input must not be empty or nil
// - the symbol immediately after an escape token is always converted to text
// - a '\\' will be converted to the text '\'
// - all escape symbols are discarded except escaped escape symbols
// - a trailing '\' in the input will be discarded
//
// Axiomatic definition of behaviour:
// - ANY := non-ESCAPE token
// - ESCAPE ANY      -> Text(ANY)
// - ESCAPE1 ESCAPE2 -> Text(ESCAPE2)
// - ESCAPE EOF      -> EOF
func applyEscaping(in []token.Lexeme) []token.Lexeme {

	size := len(in)
	out := make([]token.Lexeme, 0, size)

	for i := 0; i < size; i++ {
		tk := in[i]

		if tk.Token != token.Escape {
			out = append(out, tk)
			continue
		}

		i++
		if i < size {
			tk = in[i]
			tk.Token = token.Text
			out = append(out, tk)
		}
	}

	return out
}

// mergeLexemes merges lexemes where possible, i.e. merging sections of text
// that appear next to each other in the input.
//
// The following are some experimental documentation formats:
//
// Descriptive definition of behaviour:
// - input must not be empty or nil
// - all text tokens in series are merged into one
//
// Axiomatic definition of behaviour:
// - TEXT1 TEXT2 -> TEXT(TEXT1 + TEXT2)
func mergeLexemes(in []token.Lexeme) []token.Lexeme {

	size := len(in)
	out := make([]token.Lexeme, 0, size)

	out = append(out, in[0])
	last := 0

	for i := 1; i < size; i++ {
		lx := in[i]

		if lx.Token == token.Text && out[last].Token == token.Text {
			out[last].Val += lx.Val // Merge
			continue
		}

		out = append(out, lx)
		last++
	}

	return out
}
