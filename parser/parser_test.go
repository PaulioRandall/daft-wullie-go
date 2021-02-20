package parser

import (
	"testing"

	"github.com/PaulioRandall/daft-wullie-go/ast2"
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

	exp := []ast2.Node{
		ast2.MakeH1(ast2.MakeText("1")),
		ast2.MakeH2(ast2.MakeText("2")),
		ast2.MakeH3(ast2.MakeText("3")),
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

	exp := []ast2.Node{
		ast2.MakeQuote(
			ast2.MakeText("The Turtle Moves!"),
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

	exp := []ast2.Node{
		ast2.MakeBulPoint(
			ast2.MakeText("The Turtle Moves!"),
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

	exp := []ast2.Node{
		ast2.MakeNumPoint(
			ast2.MakeText("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestNestableNodes_1(t *testing.T) {

	lxs := func(tk token.Token, v string) []token.Lexeme {
		return []token.Lexeme{lex(tk, v), lex(tk, v)}
	}

	doTest := func(in []token.Lexeme, exp ast2.Node) {
		input := [][]token.Lexeme{in}
		expect := []ast2.Node{exp}
		act := ParseAll(input)
		require.Equal(t, expect, act)
	}

	doTest(lxs(token.KeyPhrase, "**"), ast2.MakeTextLine(ast2.MakeKeyPhrase()))
	doTest(lxs(token.Positive, "+"), ast2.MakeTextLine(ast2.MakePositive()))
	doTest(lxs(token.Negative, "-"), ast2.MakeTextLine(ast2.MakeNegative()))
	doTest(lxs(token.Strong, "*"), ast2.MakeTextLine(ast2.MakeStrong()))
	doTest(lxs(token.Snippet, "`"), ast2.MakeTextLine(ast2.MakeSnippet("")))
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
	// ! Curdling
	// ! Curd processing
	// ! Ripening
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

	exp := []ast2.Node{ // Lines
		ast2.MakeH1(
			ast2.MakeText("Cheese"),
		),
		ast2.MakeQuote(
			ast2.MakeText("Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
		),
		ast2.MakeTextLine(
			ast2.MakeStrong(
				ast2.MakeText("Cheese is "),
				ast2.MakePositive(
					ast2.MakeText("very tasty"),
				),
				ast2.MakeText(" but also quite "),
				ast2.MakeNegative(
					ast2.MakeText("smelly"),
				),
				ast2.MakeText(", "),
				ast2.MakePositive(
					ast2.MakeText("good on pizza"),
				),
			),
		),
		ast2.MakeEmptyLine(),

		ast2.MakeH2(
			ast2.MakeText("History"),
		),
		ast2.MakeTextLine(ast2.MakeText("Who knows.")),
		ast2.MakeEmptyLine(),
		ast2.MakeEmptyLine(),
		ast2.MakeEmptyLine(),

		ast2.MakeH2(
			ast2.MakeText("Types"),
		),
		ast2.MakeBulPoint(ast2.MakeText("Chedder")),
		ast2.MakeBulPoint(ast2.MakeText("Brie")),
		ast2.MakeBulPoint(ast2.MakeText("Mozzarella")),
		ast2.MakeBulPoint(ast2.MakeText("Stilton")),
		ast2.MakeBulPoint(ast2.MakeText("etc")),
		ast2.MakeEmptyLine(),

		ast2.MakeH2(
			ast2.MakeText("Process"),
		),
		ast2.MakeNumPoint(ast2.MakeText("Curdling")),
		ast2.MakeNumPoint(ast2.MakeText("Curd processing")),
		ast2.MakeNumPoint(ast2.MakeText("Ripening")),
		ast2.MakeEmptyLine(),

		ast2.MakeH2(
			ast2.MakeText("Safety"),
		),
		ast2.MakeH3(
			ast2.MakeText("Bacteria"),
		),
		ast2.MakeTextLine(
			ast2.MakeText("Milk used should be "),
			ast2.MakeKeyPhrase(ast2.MakeText("pasteurized")),
			ast2.MakeText(" to kill infectious diseases"),
		),
		ast2.MakeEmptyLine(),

		ast2.MakeH3(
			ast2.MakeText("Heart disease"),
		),
		ast2.MakeTextLine(
			ast2.MakeNegative(
				ast2.MakeText("Recommended that cheese consumption be minimised"),
			),
		),
		ast2.MakeTextLine(
			ast2.MakeNegative(
				ast2.MakeText("There isn't any "),
				ast2.MakeStrong(ast2.MakeText("convincing")),
				ast2.MakeText(" evidence that cheese lowers heart disease"),
			),
		),
		ast2.MakeEmptyLine(),

		ast2.MakeTextLine(
			ast2.MakeText("Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese"),
		),
		ast2.MakeEmptyLine(),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
