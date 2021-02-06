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

/*
func TestEmptyLine_1(t *testing.T) {

	in := [][]token.Lexeme{}

	exp := []node.Node{
		node.Empty{},
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
*/

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

func TestQuote_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.QUOTE, ">"),
			lex(token.TEXT, "The Turtle Moves!"),
		},
	}

	exp := []node.Node{
		node.Quote{M_Text: "The Turtle Moves!"},
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestNestableNodes_1(t *testing.T) {

	lxs := func(tk token.Token, v string) []token.Lexeme {
		return []token.Lexeme{lex(tk, v), lex(tk, v)}
	}

	in := [][]token.Lexeme{
		lxs(token.KEY_PHRASE, "**"),
		lxs(token.POSITIVE, "+"),
		lxs(token.NEGATIVE, "-"),
		lxs(token.STRONG, "*"),
		lxs(token.SNIPPET, "`"),
	}

	exp := []node.Node{
		node.FmtLine{M_Nodes: []node.Node{
			node.KeyPhrase{M_Nodes: []node.Node{}},
		}},
		node.FmtLine{M_Nodes: []node.Node{
			node.Positive{M_Nodes: []node.Node{}},
		}},
		node.FmtLine{M_Nodes: []node.Node{
			node.Negative{M_Nodes: []node.Node{}},
		}},
		node.FmtLine{M_Nodes: []node.Node{
			node.Strong{M_Nodes: []node.Node{}},
		}},
		node.FmtLine{M_Nodes: []node.Node{
			node.Snippet{M_Nodes: []node.Node{}},
		}},
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestScript_1(t *testing.T) {

	// # Cheese
	// >  Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein.
	// *Cheese is +very tasty+ but also quite -smelly-, +good on pizza+*
	//
	// ## History
	// Who knows.
	//
	// ## Types
	// . Chedder
	// . Brie
	// . Mozzarella
	// . Stilton
	// . etc
	//
	// ## Process
	// 1. Curdling
	// 2. Curd processing
	// 3. Ripening
	//
	// ## Safety
	// ### Bacteria
	// Milk used should be **pasteurized** to kill infectious diseases
	//
	// ### Heart disease
	// -Recommended that cheese consumption be minimised
	// -There isn't any *convincing* evidence that cheese lowers heart disease
	//
	// Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese

	/*
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
	*/
}
