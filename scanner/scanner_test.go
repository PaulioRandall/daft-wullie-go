package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func lex(tk Token, val string) Lexeme {
	return Lexeme{Token: tk, Val: val}
}

func TestTopic_1(t *testing.T) {

	in := `  #  Topic  `
	exp := [][]Lexeme{
		[]Lexeme{
			lex(TOPIC, "#"),
			lex(TEXT, "  Topic  "),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestSubTopic_1(t *testing.T) {

	in := `## Sub topic`
	exp := [][]Lexeme{
		[]Lexeme{
			lex(SUB_TOPIC, "##"),
			lex(TEXT, " Sub topic"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestHeading_1(t *testing.T) {

	in := `### Heading`
	exp := [][]Lexeme{
		[]Lexeme{
			lex(HEADING, "###"),
			lex(TEXT, " Heading"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestBulletPoint_1(t *testing.T) {

	in := `. Point`
	exp := [][]Lexeme{
		[]Lexeme{
			lex(BUL_POINT, "."),
			lex(TEXT, " Point"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNumberPoint_1(t *testing.T) {

	in := `123. Point`
	exp := [][]Lexeme{
		[]Lexeme{
			lex(NUM_POINT, "123."),
			lex(TEXT, " Point"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestQuote_1(t *testing.T) {

	in := `> Fly high through apocalypse skies`
	exp := [][]Lexeme{
		[]Lexeme{
			lex(QUOTE, ">"),
			lex(TEXT, " Fly high through apocalypse skies"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNodes_1(t *testing.T) {

	in := "**+-*`?"
	exp := [][]Lexeme{
		[]Lexeme{
			lex(KEY_PHRASE, "**"),
			lex(POSITIVE, "+"),
			lex(NEGATIVE, "-"),
			lex(STRONG, "*"),
			lex(SNIPPET, "`"),
			lex(QUESTION, "?"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNodes_2(t *testing.T) {

	in := "* +positive+ and -negative- *"
	exp := [][]Lexeme{
		[]Lexeme{
			lex(STRONG, "*"),
			lex(TEXT, " "),
			lex(POSITIVE, "+"),
			lex(TEXT, "positive"),
			lex(POSITIVE, "+"),
			lex(TEXT, " and "),
			lex(NEGATIVE, "-"),
			lex(TEXT, "negative"),
			lex(NEGATIVE, "-"),
			lex(TEXT, " "),
			lex(STRONG, "*"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestLines_1(t *testing.T) {

	line := func(lxs ...Lexeme) []Lexeme { return lxs }
	emptyLine := func() []Lexeme { return []Lexeme{} }

	in := `
### Trees:


. Burnable -(Wildfires)-
. Central to *ecosystems*
.+ Fun to climb
.- Can fall over
`
	exp := [][]Lexeme{
		emptyLine(),
		line(
			lex(HEADING, "###"),
			lex(TEXT, " Trees:"),
		),
		emptyLine(),
		emptyLine(),
		line(
			lex(BUL_POINT, "."),
			lex(TEXT, " Burnable "),
			lex(NEGATIVE, "-"),
			lex(TEXT, "(Wildfires)"),
			lex(NEGATIVE, "-"),
		),
		line(
			lex(BUL_POINT, "."),
			lex(TEXT, " Central to "),
			lex(STRONG, "*"),
			lex(TEXT, "ecosystems"),
			lex(STRONG, "*"),
		),
		line(
			lex(BUL_POINT, "."),
			lex(POSITIVE, "+"),
			lex(TEXT, " Fun to climb"),
		),
		line(
			lex(BUL_POINT, "."),
			lex(NEGATIVE, "-"),
			lex(TEXT, " Can fall over"),
		),
		emptyLine(),
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestLines_2(t *testing.T) {

	line := func(lxs ...Lexeme) []Lexeme { return lxs }
	emptyLine := func() []Lexeme { return []Lexeme{} }

	in := `
> I aten't ded
A quote by whom?
`
	exp := [][]Lexeme{
		emptyLine(),
		line(
			lex(QUOTE, ">"),
			lex(TEXT, " I aten't ded"),
		),
		line(
			lex(TEXT, "A quote by whom"),
			lex(QUESTION, "?"),
		),
		emptyLine(),
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}
