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

func neverNil(nodes []node.Node) []node.Node {
	if nodes == nil {
		return []node.Node{}
	}
	return nodes
}

func h1(text string) node.H1                      { return node.H1{M_Text: text} }
func h2(text string) node.H2                      { return node.H2{M_Text: text} }
func h3(text string) node.H3                      { return node.H3{M_Text: text} }
func quote____(text string) node.Quote            { return node.Quote{M_Text: text} }
func fmtLine__(nodes ...node.Node) node.FmtLine   { return node.FmtLine{M_Nodes: neverNil(nodes)} }
func keyPhrase(nodes ...node.Node) node.KeyPhrase { return node.KeyPhrase{M_Nodes: neverNil(nodes)} }
func positive_(nodes ...node.Node) node.Positive  { return node.Positive{M_Nodes: neverNil(nodes)} }
func negative_(nodes ...node.Node) node.Negative  { return node.Negative{M_Nodes: neverNil(nodes)} }
func strong___(nodes ...node.Node) node.Strong    { return node.Strong{M_Nodes: neverNil(nodes)} }
func snippet__(nodes ...node.Node) node.Snippet   { return node.Snippet{M_Nodes: neverNil(nodes)} }
func phrase___(text string) node.Phrase           { return node.Phrase{M_Text: text} }

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
		h1(""),
		h2(""),
		h3(""),
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
		quote____("The Turtle Moves!"),
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
		fmtLine__(keyPhrase()),
		fmtLine__(positive_()),
		fmtLine__(negative_()),
		fmtLine__(strong___()),
		fmtLine__(snippet__()),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestScript_1(t *testing.T) {

	// # Cheese
	// >  Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein.
	// *Cheese is +very tasty+ but also quite -smelly-, +good on pizza+
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

	in := [][]token.Lexeme{ // Lines
		[]token.Lexeme{lex(token.H1, "#"), lex(token.TEXT, "Cheese")},
		[]token.Lexeme{lex(token.QUOTE, ">"), lex(token.TEXT, "Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein.")},
		[]token.Lexeme{
			lex(token.STRONG, "*"),
			lex(token.TEXT, "Cheese is "),
			lex(token.POSITIVE, "+"),
			lex(token.TEXT, "very tasty"),
			lex(token.POSITIVE, "+"),
			lex(token.TEXT, " but also quite "),
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, "smelly"),
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, ", "),
			lex(token.POSITIVE, "+"),
			lex(token.TEXT, "good on pizza"),
			lex(token.POSITIVE, "+"),
		},
	}

	exp := []node.Node{ // Lines
		h1("Cheese"),
		quote____("Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		fmtLine__(
			strong___(
				phrase___("Cheese is "),
				positive_(phrase___("very tasty")),
				phrase___(" but also quite "),
				negative_(phrase___("smelly")),
				phrase___(", "),
				positive_(phrase___("good on pizza")),
			),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
