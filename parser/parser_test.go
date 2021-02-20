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
		[]token.Lexeme{lex(token.H1, "#"), lex(token.Text, "1")},
		[]token.Lexeme{lex(token.H2, "##"), lex(token.Text, "2")},
		[]token.Lexeme{lex(token.H3, "###"), lex(token.Text, "3")},
	}

	exp := []node.Node{
		node.MakeH1(node.MakePhrase("1")),
		node.MakeH2(node.MakePhrase("2")),
		node.MakeH3(node.MakePhrase("3")),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestQuote_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.Quote, ">"),
			lex(token.Text, "The Turtle Moves!"),
		},
	}

	exp := []node.Node{
		node.MakeQuote(
			node.MakePhrase("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestBulPoint_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.BulPoint, "."),
			lex(token.Text, "The Turtle Moves!"),
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
			lex(token.NumPoint, "9."),
			lex(token.Text, "The Turtle Moves!"),
		},
	}

	exp := []node.Node{
		node.MakeNumPoint(
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

	doTest(lxs(token.KeyPhrase, "**"), node.MakeFmtLine(node.MakeKeyPhrase()))
	doTest(lxs(token.Positive, "+"), node.MakeFmtLine(node.MakePositive()))
	doTest(lxs(token.Negative, "-"), node.MakeFmtLine(node.MakeNegative()))
	doTest(lxs(token.Strong, "*"), node.MakeFmtLine(node.MakeStrong()))
	doTest(lxs(token.Snippet, "`"), node.MakeFmtLine(node.MakeSnippet("")))
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
			lex(token.Text, "Cheese"),
		},
		[]token.Lexeme{
			lex(token.Quote, ">"),
			lex(token.Text, "Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		},
		[]token.Lexeme{
			lex(token.Strong, "*"),
			lex(token.Text, "Cheese is "),
			lex(token.Positive, "+"),
			lex(token.Text, "very tasty"),
			lex(token.Positive, "+"),
			lex(token.Text, " but also quite "),
			lex(token.Negative, "-"),
			lex(token.Text, "smelly"),
			lex(token.Negative, "-"),
			lex(token.Text, ", "),
			lex(token.Positive, "+"),
			lex(token.Text, "good on pizza"),
			lex(token.Positive, "+"),
		},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.Text, "History")},
		[]token.Lexeme{lex(token.Text, "Who knows.")},
		[]token.Lexeme{},
		[]token.Lexeme{},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.Text, "Types")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Chedder")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Brie")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Mozzarella")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Stilton")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "etc")},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.Text, "Process")},
		[]token.Lexeme{lex(token.NumPoint, "!"), lex(token.Text, "Curdling")},
		[]token.Lexeme{lex(token.NumPoint, "!"), lex(token.Text, "Curd processing")},
		[]token.Lexeme{lex(token.NumPoint, "!"), lex(token.Text, "Ripening")},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H2, "##"), lex(token.Text, "Safety")},
		[]token.Lexeme{lex(token.H3, "###"), lex(token.Text, "Bacteria")},
		[]token.Lexeme{
			lex(token.Text, "Milk used should be "),
			lex(token.KeyPhrase, "**"),
			lex(token.Text, "pasteurized"),
			lex(token.KeyPhrase, "**"),
			lex(token.Text, " to kill infectious diseases"),
		},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.H3, "###"), lex(token.Text, "Heart disease")},
		[]token.Lexeme{
			lex(token.Negative, "-"),
			lex(token.Text, "Recommended that cheese consumption be minimised"),
		},
		[]token.Lexeme{
			lex(token.Negative, "-"),
			lex(token.Text, "There isn't any "),
			lex(token.Strong, "*"),
			lex(token.Text, "convincing"),
			lex(token.Strong, "*"),
			lex(token.Text, " evidence that cheese lowers heart disease"),
		},
		[]token.Lexeme{},

		[]token.Lexeme{
			lex(token.Text, "Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese"),
		},
		[]token.Lexeme{},
	}

	exp := []node.Node{ // Lines
		node.MakeH1(
			node.MakePhrase("Cheese"),
		),
		node.MakeQuote(
			node.MakePhrase("Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		),
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

		node.MakeH2(
			node.MakePhrase("History"),
		),
		node.MakeFmtLine(node.MakePhrase("Who knows.")),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),

		node.MakeH2(
			node.MakePhrase("Types"),
		),
		node.MakeBulPoint(node.MakePhrase("Chedder")),
		node.MakeBulPoint(node.MakePhrase("Brie")),
		node.MakeBulPoint(node.MakePhrase("Mozzarella")),
		node.MakeBulPoint(node.MakePhrase("Stilton")),
		node.MakeBulPoint(node.MakePhrase("etc")),
		node.MakeEmptyLine(),

		node.MakeH2(
			node.MakePhrase("Process"),
		),
		node.MakeNumPoint(node.MakePhrase("Curdling")),
		node.MakeNumPoint(node.MakePhrase("Curd processing")),
		node.MakeNumPoint(node.MakePhrase("Ripening")),
		node.MakeEmptyLine(),

		node.MakeH2(
			node.MakePhrase("Safety"),
		),
		node.MakeH3(
			node.MakePhrase("Bacteria"),
		),
		node.MakeFmtLine(
			node.MakePhrase("Milk used should be "),
			node.MakeKeyPhrase(node.MakePhrase("pasteurized")),
			node.MakePhrase(" to kill infectious diseases"),
		),
		node.MakeEmptyLine(),

		node.MakeH3(
			node.MakePhrase("Heart disease"),
		),
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
