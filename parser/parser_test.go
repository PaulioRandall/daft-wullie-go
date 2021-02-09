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

func emptyLine() node.Empty { return node.Empty{} }

func h1(text string) node.H1 { return node.H1{M_Text: text} }
func h2(text string) node.H2 { return node.H2{M_Text: text} }
func h3(text string) node.H3 { return node.H3{M_Text: text} }

func quote(text string) node.Quote              { return node.Quote{M_Text: text} }
func fmtLine(nodes ...node.Node) node.FmtLine   { return node.FmtLine{M_Nodes: neverNil(nodes)} }
func bulPoint(nodes ...node.Node) node.BulPoint { return node.BulPoint{M_Nodes: neverNil(nodes)} }
func numPoint(nodes ...node.Node) node.NumPoint { return node.NumPoint{M_Nodes: neverNil(nodes)} }

func keyPhrase(nodes ...node.Node) node.KeyPhrase { return node.KeyPhrase{M_Nodes: neverNil(nodes)} }
func positive(nodes ...node.Node) node.Positive   { return node.Positive{M_Nodes: neverNil(nodes)} }
func negative(nodes ...node.Node) node.Negative   { return node.Negative{M_Nodes: neverNil(nodes)} }
func strong(nodes ...node.Node) node.Strong       { return node.Strong{M_Nodes: neverNil(nodes)} }
func snippet(nodes ...node.Node) node.Snippet     { return node.Snippet{M_Nodes: neverNil(nodes)} }
func phrase(text string) node.Phrase              { return node.Phrase{M_Text: text} }

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
		quote("The Turtle Moves!"),
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
		fmtLine(keyPhrase()),
		fmtLine(positive()),
		fmtLine(negative()),
		fmtLine(strong()),
		fmtLine(snippet()),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestScript_1(t *testing.T) {

	// # Cheese
	// > Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein.
	// *Cheese is +very tasty+ but also quite -smelly-, +good on pizza+
	//
	// ## History
	// Who knows.
	//
	//
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
		[]token.Lexeme{
			lex(token.H1, "#"),
			lex(token.TEXT, "Cheese"),
		},
		[]token.Lexeme{
			lex(token.QUOTE, ">"),
			lex(token.TEXT, "Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		},
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
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.TEXT, "History")},
		[]token.Lexeme{lex(token.TEXT, "Who knows.")},
		[]token.Lexeme{},
		[]token.Lexeme{},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.TEXT, "Types")},
		[]token.Lexeme{lex(token.BUL_POINT, "."), lex(token.TEXT, "Chedder")},
		[]token.Lexeme{lex(token.BUL_POINT, "."), lex(token.TEXT, "Brie")},
		[]token.Lexeme{lex(token.BUL_POINT, "."), lex(token.TEXT, "Mozzarella")},
		[]token.Lexeme{lex(token.BUL_POINT, "."), lex(token.TEXT, "Stilton")},
		[]token.Lexeme{lex(token.BUL_POINT, "."), lex(token.TEXT, "etc")},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.TEXT, "Process")},
		[]token.Lexeme{lex(token.NUM_POINT, "1."), lex(token.TEXT, "Curdling")},
		[]token.Lexeme{lex(token.NUM_POINT, "2."), lex(token.TEXT, "Curd processing")},
		[]token.Lexeme{lex(token.NUM_POINT, "3."), lex(token.TEXT, "Ripening")},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.TEXT, "Safety")},
		[]token.Lexeme{lex(token.H3, "###"), lex(token.TEXT, "Bacteria")},
		[]token.Lexeme{
			lex(token.TEXT, "Milk used should be "),
			lex(token.KEY_PHRASE, "**"),
			lex(token.TEXT, "pasteurized"),
			lex(token.KEY_PHRASE, "**"),
			lex(token.TEXT, " to kill infectious diseases"),
		},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H3, "###"), lex(token.TEXT, "Heart disease")},
		[]token.Lexeme{
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, "Recommended that cheese consumption be minimised"),
		},
		[]token.Lexeme{
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, "There isn't any "),
			lex(token.STRONG, "*"),
			lex(token.TEXT, "convincing"),
			lex(token.STRONG, "*"),
			lex(token.TEXT, " evidence that cheese lowers heart disease"),
		},
		[]token.Lexeme{},

		[]token.Lexeme{
			lex(token.TEXT, "Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese"),
		},
		[]token.Lexeme{},
	}

	exp := []node.Node{ // Lines
		h1("Cheese"),
		quote("Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		fmtLine(
			strong(
				phrase("Cheese is "),
				positive(
					phrase("very tasty"),
				),
				phrase(" but also quite "),
				negative(
					phrase("smelly"),
				),
				phrase(", "),
				positive(
					phrase("good on pizza"),
				),
			),
		),
		emptyLine(),

		h2("History"),
		fmtLine(phrase("Who knows.")),
		emptyLine(),
		emptyLine(),
		emptyLine(),

		h2("Types"),
		bulPoint(phrase("Chedder")),
		bulPoint(phrase("Brie")),
		bulPoint(phrase("Mozzarella")),
		bulPoint(phrase("Stilton")),
		bulPoint(phrase("etc")),
		emptyLine(),

		h2("Process"),
		numPoint(phrase("Curdling")),
		numPoint(phrase("Curd processing")),
		numPoint(phrase("Ripening")),
		emptyLine(),

		h2("Safety"),
		h3("Bacteria"),
		fmtLine(
			phrase("Milk used should be "),
			keyPhrase(phrase("pasteurized")),
			phrase(" to kill infectious diseases"),
		),
		emptyLine(),

		h3("Heart disease"),
		fmtLine(
			negative(
				phrase("Recommended that cheese consumption be minimised"),
			),
		),
		fmtLine(
			negative(
				phrase("There isn't any "),
				strong(phrase("convincing")),
				phrase(" evidence that cheese lowers heart disease"),
			),
		),
		emptyLine(),

		fmtLine(
			phrase("Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese"),
		),
		emptyLine(),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
