package parser

import (
	"testing"

	"github.com/PaulioRandall/daft-wullie-go/node"
	"github.com/PaulioRandall/daft-wullie-go/token"

	"github.com/stretchr/testify/require"
)

func lex(tk token.Token, val string) token.Lexeme {
	return token.Lexeme{Token: tk, Val: val}
}

func TestHeadings_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{lex(token.H1, "#"), lex(token.TEXT, "")},
		[]token.Lexeme{lex(token.H2, "##"), lex(token.TEXT, "")},
		[]token.Lexeme{lex(token.H3, "###"), lex(token.TEXT, "")},
	}

	exp := []node.Node{
		node.H1{M_Text: ""},
		node.H2{M_Text: ""},
		node.H3{M_Text: ""},
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
