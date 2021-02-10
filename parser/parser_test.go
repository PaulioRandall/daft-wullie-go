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
		node.MakeH1(""),
		node.MakeH2(""),
		node.MakeH3(""),
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
		node.MakeQuote("The Turtle Moves!"),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestBulPoint_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.BUL_POINT, "."),
			lex(token.TEXT, "The Turtle Moves!"),
		},
	}

	exp := []node.Node{
		node.MakeBulPoint(
			node.MakePhrase("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestNumPoint_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.NUM_POINT, "9."),
			lex(token.TEXT, "The Turtle Moves!"),
		},
	}

	exp := []node.Node{
		node.MakeNumPoint(
			"9",
			node.MakePhrase("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestNestableNodes_1(t *testing.T) {

	lxs := func(tk token.Token, v string) []token.Lexeme {
		return []token.Lexeme{lex(tk, v), lex(tk, v)}
	}

	doTest := func(in []token.Lexeme, exp node.Node) {
		input := [][]token.Lexeme{in}
		expect := []node.Node{exp}
		act := ParseAll(input)
		require.Equal(t, expect, act)
	}

	doTest(lxs(token.KEY_PHRASE, "**"), node.MakeFmtLine(node.MakeKeyPhrase()))
	doTest(lxs(token.POSITIVE, "+"), node.MakeFmtLine(node.MakePositive()))
	doTest(lxs(token.NEGATIVE, "-"), node.MakeFmtLine(node.MakeNegative()))
	doTest(lxs(token.STRONG, "*"), node.MakeFmtLine(node.MakeStrong()))
	doTest(lxs(token.SNIPPET, "`"), node.MakeFmtLine(node.MakeSnippet("")))
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
		node.MakeH1("Cheese"),
		node.MakeQuote("Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		node.MakeFmtLine(
			node.MakeStrong(
				node.MakePhrase("Cheese is "),
				node.MakePositive(
					node.MakePhrase("very tasty"),
				),
				node.MakePhrase(" but also quite "),
				node.MakeNegative(
					node.MakePhrase("smelly"),
				),
				node.MakePhrase(", "),
				node.MakePositive(
					node.MakePhrase("good on pizza"),
				),
			),
		),
		node.MakeEmptyLine(),

		node.MakeH2("History"),
		node.MakeFmtLine(node.MakePhrase("Who knows.")),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),

		node.MakeH2("Types"),
		node.MakeBulPoint(node.MakePhrase("Chedder")),
		node.MakeBulPoint(node.MakePhrase("Brie")),
		node.MakeBulPoint(node.MakePhrase("Mozzarella")),
		node.MakeBulPoint(node.MakePhrase("Stilton")),
		node.MakeBulPoint(node.MakePhrase("etc")),
		node.MakeEmptyLine(),

		node.MakeH2("Process"),
		node.MakeNumPoint("1", node.MakePhrase("Curdling")),
		node.MakeNumPoint("2", node.MakePhrase("Curd processing")),
		node.MakeNumPoint("3", node.MakePhrase("Ripening")),
		node.MakeEmptyLine(),

		node.MakeH2("Safety"),
		node.MakeH3("Bacteria"),
		node.MakeFmtLine(
			node.MakePhrase("Milk used should be "),
			node.MakeKeyPhrase(node.MakePhrase("pasteurized")),
			node.MakePhrase(" to kill infectious diseases"),
		),
		node.MakeEmptyLine(),

		node.MakeH3("Heart disease"),
		node.MakeFmtLine(
			node.MakeNegative(
				node.MakePhrase("Recommended that cheese consumption be minimised"),
			),
		),
		node.MakeFmtLine(
			node.MakeNegative(
				node.MakePhrase("There isn't any "),
				node.MakeStrong(node.MakePhrase("convincing")),
				node.MakePhrase(" evidence that cheese lowers heart disease"),
			),
		),
		node.MakeEmptyLine(),

		node.MakeFmtLine(
			node.MakePhrase("Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese"),
		),
		node.MakeEmptyLine(),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
