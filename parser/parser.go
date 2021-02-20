package parser

import (
	"strings"

	"github.com/PaulioRandall/daft-wullie-go/node"
	"github.com/PaulioRandall/daft-wullie-go/token"
)

// ParseLine is function for recursively parsing scanned text lines,
// represented by sets of lexemes, into ASTs.
type ParseLine func() (node.LineNode, ParseLine)

// NewParser creates an initial ParseLine function for parsing 'lines'.
func NewParser(lines [][]token.Lexeme) ParseLine {
	r := &lineReader{lines: lines}
	if !r.more() {
		return nil
	}
	return parser(r)
}

// ParseAll scans all 'lines' into a slice of ASTs, each representing a line of
// annotated text.
func ParseAll(lines [][]token.Lexeme) []node.LineNode {
	var (
		f = NewParser(lines)
		r = []node.LineNode{}
		n node.LineNode
	)
	for f != nil {
		n, f = f()
		r = append(r, n)
	}
	return r
}

func parser(r *lineReader) ParseLine {
	return func() (node.LineNode, ParseLine) {
		lr := r.nextLine()
		ns := parseLine(lr)
		if r.more() {
			return ns, parser(r)
		}
		return ns, nil
	}
}

// LINE := *Nothing/empty*
// LINE := (H1 | H2 | H3 | QUOTE) TEXT_LINE
// LINE := [BUL_POINT | NUM_POINT] NODE_LINE
func parseLine(r *tokenReader) node.LineNode {
	switch {
	case !r.more():
		return node.MakeEmptyLine()

	case r.accept(token.H1):
		return node.MakeH1(parseNodeLine(r)...)

	case r.accept(token.H2):
		return node.MakeH2(parseNodeLine(r)...)

	case r.accept(token.H3):
		return node.MakeH3(parseNodeLine(r)...)

	case r.accept(token.Quote):
		return node.MakeQuote(parseNodeLine(r)...)

	case r.accept(token.BulPoint):
		return node.MakeBulPoint(parseNodeLine(r)...)

	case r.accept(token.NumPoint):
		return node.MakeNumPoint(parseNodeLine(r)...)

	default:
		return node.MakeTextLine(parseNodeLine(r)...)
	}
}

// NODE_LINE := {NODE} *EOF*
func parseNodeLine(r *tokenReader) []node.Node {
	ns := []node.Node{}
	for r.more() {
		n := parseNode(r)
		ns = append(ns, n)
	}
	return ns
}

// NODE := KEY_PHRASE {NODE} [KEY_PHRASE]
// NODE := POSITIVE   {NODE} [POSITIVE]
// NODE := NEGATIVE   {NODE} [NEGATIVE]
// NODE := STRONG     {NODE} [STRONG]
// NODE := SNIPPET    {NODE} [SNIPPET]
// NODE := TEXT_PHRASE
func parseNode(r *tokenReader) node.Node {
	switch {
	case r.accept(token.KeyPhrase):
		return node.MakeKeyPhrase(parseNodesUntil(r, token.KeyPhrase)...)

	case r.accept(token.Positive):
		return node.MakePositive(parseNodesUntil(r, token.Positive)...)

	case r.accept(token.Negative):
		return node.MakeNegative(parseNodesUntil(r, token.Negative)...)

	case r.accept(token.Strong):
		return node.MakeStrong(parseNodesUntil(r, token.Strong)...)

	case r.accept(token.Snippet):
		return node.MakeSnippet(parseTextUntil(r, token.Snippet))

	default:
		return node.Phrase{Txt: parseText(r)}
	}
}

// parseNodesUntil parses child nodes until the end of the line or the
// specified 'delim' is encountered. Upon which, the delim is read and
// discarded before the children are returned.
//
// Note: nesting may occur but only when the parent and child nodes are of
// different types. E.g. no point having strong text decoration within strong
// text decoration unless some intermidiate node negates the affect.
func parseNodesUntil(r *tokenReader, delim token.Token) []node.Node {
	ns := []node.Node{}
	for r.more() && !r.accept(delim) {
		n := parseNode(r)
		ns = append(ns, n)
	}
	return ns
}

// parseTextUntil parses text until the end of the line or the specified
// 'delim' is encountered. Upon which, the delim is read and discarded before
// the text is returned.
func parseTextUntil(r *tokenReader, delim token.Token) string {
	sb := strings.Builder{}
	for r.more() && !r.accept(delim) {
		s := r.read().Val
		sb.WriteString(s)
	}
	return sb.String()
}

// TEXT_LINE := {TEXT} *EOF*
func parseTextLine(r *tokenReader) string {
	sb := strings.Builder{}
	for r.more() {
		s := r.read().Val
		sb.WriteString(s)
	}
	return sb.String()
}

// TEXT_PHRASE := {TEXT}
func parseText(r *tokenReader) string {
	sb := strings.Builder{}
	for r.more() && r.match(token.Text) {
		s := r.read().Val
		sb.WriteString(s)
	}
	return sb.String()
}
