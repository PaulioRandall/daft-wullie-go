package parser

import (
	"github.com/PaulioRandall/daft-wullie-go/token"
)

type (
	lineReader struct {
		lines [][]token.Lexeme
		idx   int
	}

	tokenReader struct {
		tks []token.Lexeme
		idx int
	}
)

func (r *lineReader) more() bool {
	return r.idx < len(r.lines)
}

func (r *lineReader) nextLine() *tokenReader {
	if !r.more() {
		panic("Line out of range, check for EOF first")
	}
	rr := &tokenReader{tks: r.lines[r.idx]}
	r.idx++
	return rr
}

func (r *tokenReader) _curr() token.Token {
	return r.tks[r.idx].Token
}

func (r *tokenReader) more() bool {
	return r.idx < len(r.tks)
}

func (r *tokenReader) match(tk token.Token) bool {
	return r.more() && r._curr() == tk
}

func (r *tokenReader) read() token.Lexeme {
	if !r.more() {
		panic("Lexeme out of range, check for EOF first")
	}
	lx := r.tks[r.idx]
	r.idx++
	return lx
}

func (r *tokenReader) accept(tk token.Token) bool {
	if r.match(tk) {
		r.read()
		return true
	}
	return false
}

func (r *tokenReader) backup() {
	if r.idx == 0 {
		panic("Colunm index out of range")
	}
	r.idx--
}
