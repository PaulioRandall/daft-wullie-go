package parser

import (
	"testing"

	"github.com/PaulioRandall/daft-wullie-go/ast"
	"github.com/PaulioRandall/daft-wullie-go/token"

	"github.com/stretchr/testify/require"
)

func lex(tk token.Token, val string) token.Lexeme {
	return token.Lexeme{Token: tk, Val: val}
}

func TestHeadings_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{lex(token.Topic, "#"), lex(token.Text, "1")},
		[]token.Lexeme{lex(token.SubTopic, "##"), lex(token.Text, "2")},
	}

	exp := []ast.Node{
		ast.MakeTopic(ast.MakeText("1")),
		ast.MakeSubTopic(ast.MakeText("2")),
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

	exp := []ast.Node{
		ast.MakeBulPoint(
			ast.MakeText("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestSubBulPoint_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.SubBulPoint, ".."),
			lex(token.Text, "The Turtle Moves!"),
		},
	}

	exp := []ast.Node{
		ast.MakeSubBulPoint(
			ast.MakeText("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestNumPoint_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.NumPoint, "!"),
			lex(token.Text, "The Turtle Moves!"),
		},
	}

	exp := []ast.Node{
		ast.MakeNumPoint(
			ast.MakeText("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestSubNumPoint_1(t *testing.T) {

	in := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.SubNumPoint, "!!"),
			lex(token.Text, "The Turtle Moves!"),
		},
	}

	exp := []ast.Node{
		ast.MakeSubNumPoint(
			ast.MakeText("The Turtle Moves!"),
		),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}

func TestNestableNodes_1(t *testing.T) {

	lxs := func(tk token.Token, v string) []token.Lexeme {
		return []token.Lexeme{lex(tk, v), lex(tk, v)}
	}

	doTest := func(in []token.Lexeme, exp ast.Node) {
		input := [][]token.Lexeme{in}
		expect := []ast.Node{exp}
		act := ParseAll(input)
		require.Equal(t, expect, act)
	}

	doTest(lxs(token.KeyPhrase, "**"), ast.MakeTextLine(ast.MakeKeyPhrase()))
	doTest(lxs(token.Positive, "+"), ast.MakeTextLine(ast.MakePositive()))
	doTest(lxs(token.Negative, "-"), ast.MakeTextLine(ast.MakeNegative()))
	doTest(lxs(token.Strong, "*"), ast.MakeTextLine(ast.MakeStrong()))
	doTest(lxs(token.Quote, `"`), ast.MakeTextLine(ast.MakeQuote()))
	doTest(lxs(token.Artifact, "$"), ast.MakeTextLine(ast.MakeArtifact()))
	doTest(lxs(token.Snippet, "`"), ast.MakeTextLine(ast.MakeSnippet("")))
}

func TestScript_1(t *testing.T) {

	// # Cheese
	// "Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein.
	// *Cheese is +very tasty+ but also quite -smelly-, +good on pizza+
	//
	// ## History
	// Who knows.
	//
	//
	//
	// ## Types
	// . Chedder, always from $Chedder,Somerset,England
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
	// ## Bacteria
	// Milk used should be **pasteurized** to kill infectious diseases
	//
	// ## Heart disease
	// -Recommended that cheese consumption be minimised
	// -There isn't any *convincing* evidence that cheese lowers heart disease
	//
	// Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese

	in := [][]token.Lexeme{ // Lines
		[]token.Lexeme{
			lex(token.Topic, "#"),
			lex(token.Text, "Cheese"),
		},
		[]token.Lexeme{
			lex(token.Quote, `"`),
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

		[]token.Lexeme{lex(token.SubTopic, "##"), lex(token.Text, "History")},
		[]token.Lexeme{lex(token.Text, "Who knows.")},
		[]token.Lexeme{},
		[]token.Lexeme{},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.SubTopic, "##"), lex(token.Text, "Types")},
		[]token.Lexeme{lex(token.BulPoint, "."),
			lex(token.Text, "Chedder, always from "),
			lex(token.Artifact, "$"),
			lex(token.Text, "Chedder,Somerset,England"),
		},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Brie")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Mozzarella")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "Stilton")},
		[]token.Lexeme{lex(token.BulPoint, "."), lex(token.Text, "etc")},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.SubTopic, "##"), lex(token.Text, "Process")},
		[]token.Lexeme{lex(token.NumPoint, "!"), lex(token.Text, "Curdling")},
		[]token.Lexeme{lex(token.NumPoint, "!"), lex(token.Text, "Curd processing")},
		[]token.Lexeme{lex(token.NumPoint, "!"), lex(token.Text, "Ripening")},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.SubTopic, "##"), lex(token.Text, "Bacteria")},
		[]token.Lexeme{
			lex(token.Text, "Milk used should be "),
			lex(token.KeyPhrase, "**"),
			lex(token.Text, "pasteurized"),
			lex(token.KeyPhrase, "**"),
			lex(token.Text, " to kill infectious diseases"),
		},
		[]token.Lexeme{},

		[]token.Lexeme{lex(token.SubTopic, "##"), lex(token.Text, "Heart disease")},
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

	exp := []ast.Node{ // Lines
		ast.MakeTopic(
			ast.MakeText("Cheese"),
		),
		ast.MakeTextLine(
			ast.MakeQuote(
				ast.MakeText("Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein."),
			),
		),
		ast.MakeTextLine(
			ast.MakeStrong(
				ast.MakeText("Cheese is "),
				ast.MakePositive(
					ast.MakeText("very tasty"),
				),
				ast.MakeText(" but also quite "),
				ast.MakeNegative(
					ast.MakeText("smelly"),
				),
				ast.MakeText(", "),
				ast.MakePositive(
					ast.MakeText("good on pizza"),
				),
			),
		),
		ast.MakeEmptyLine(),

		ast.MakeSubTopic(
			ast.MakeText("History"),
		),
		ast.MakeTextLine(ast.MakeText("Who knows.")),
		ast.MakeEmptyLine(),
		ast.MakeEmptyLine(),
		ast.MakeEmptyLine(),

		ast.MakeSubTopic(
			ast.MakeText("Types"),
		),
		ast.MakeBulPoint(
			ast.MakeText("Chedder, always from "),
			ast.MakeArtifact(ast.MakeText("Chedder,Somerset,England")),
		),
		ast.MakeBulPoint(ast.MakeText("Brie")),
		ast.MakeBulPoint(ast.MakeText("Mozzarella")),
		ast.MakeBulPoint(ast.MakeText("Stilton")),
		ast.MakeBulPoint(ast.MakeText("etc")),
		ast.MakeEmptyLine(),

		ast.MakeSubTopic(
			ast.MakeText("Process"),
		),
		ast.MakeNumPoint(ast.MakeText("Curdling")),
		ast.MakeNumPoint(ast.MakeText("Curd processing")),
		ast.MakeNumPoint(ast.MakeText("Ripening")),
		ast.MakeEmptyLine(),

		ast.MakeSubTopic(
			ast.MakeText("Bacteria"),
		),
		ast.MakeTextLine(
			ast.MakeText("Milk used should be "),
			ast.MakeKeyPhrase(ast.MakeText("pasteurized")),
			ast.MakeText(" to kill infectious diseases"),
		),
		ast.MakeEmptyLine(),

		ast.MakeSubTopic(
			ast.MakeText("Heart disease"),
		),
		ast.MakeTextLine(
			ast.MakeNegative(
				ast.MakeText("Recommended that cheese consumption be minimised"),
			),
		),
		ast.MakeTextLine(
			ast.MakeNegative(
				ast.MakeText("There isn't any "),
				ast.MakeStrong(ast.MakeText("convincing")),
				ast.MakeText(" evidence that cheese lowers heart disease"),
			),
		),
		ast.MakeEmptyLine(),

		ast.MakeTextLine(
			ast.MakeText("Source [2021-02-06]: https://en.wikipedia.org/wiki/Cheese"),
		),
		ast.MakeEmptyLine(),
	}

	act := ParseAll(in)
	require.Equal(t, exp, act)
}
