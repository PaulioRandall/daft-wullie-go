package scanner

import (
	"testing"

	"github.com/PaulioRandall/daft-wullie-go/token"

	"github.com/stretchr/testify/require"
)

func lex(tk token.Token, val string) token.Lexeme {
	return token.Lexeme{Token: tk, Val: val}
}

func TestEscape_1(t *testing.T) {

	in := `\#daft`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.TEXT, "#daft"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestEscape_2(t *testing.T) {

	in := `\da\ft`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.TEXT, `daft`),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestEscape_3(t *testing.T) {

	in := `da\\ft`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.TEXT, `da\ft`),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestTopic_1(t *testing.T) {

	in := `  #  Topic  `
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.H1, "#"),
			lex(token.TEXT, "  Topic  "),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestSubTopic_1(t *testing.T) {

	in := `## Sub topic`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.H2, "##"),
			lex(token.TEXT, " Sub topic"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestHeading_1(t *testing.T) {

	in := `### Heading`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.H3, "###"),
			lex(token.TEXT, " Heading"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestBulletPoint_1(t *testing.T) {

	in := `. Point`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.BUL_POINT, "."),
			lex(token.TEXT, " Point"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNumberPoint_1(t *testing.T) {

	in := `123. Point`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.NUM_POINT, "123."),
			lex(token.TEXT, " Point"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestQuote_1(t *testing.T) {

	in := `> Fly high through apocalypse skies`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.QUOTE, ">"),
			lex(token.TEXT, " Fly high through apocalypse skies"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNodes_1(t *testing.T) {

	in := "**+-*`"
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.KEY_PHRASE, "**"),
			lex(token.POSITIVE, "+"),
			lex(token.NEGATIVE, "-"),
			lex(token.STRONG, "*"),
			lex(token.SNIPPET, "`"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNodes_2(t *testing.T) {

	in := "* +positive+ and -negative- *"
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.STRONG, "*"),
			lex(token.TEXT, " "),
			lex(token.POSITIVE, "+"),
			lex(token.TEXT, "positive"),
			lex(token.POSITIVE, "+"),
			lex(token.TEXT, " and "),
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, "negative"),
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, " "),
			lex(token.STRONG, "*"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestLines_1(t *testing.T) {

	line := func(lxs ...token.Lexeme) []token.Lexeme { return lxs }
	emptyLine := func() []token.Lexeme { return []token.Lexeme{} }

	in := `
### Trees:


. Burnable -(Wildfires)-
. Central to *ecosystems*
.+ Fun to climb
.- Can fall over
`
	exp := [][]token.Lexeme{
		emptyLine(),
		line(
			lex(token.H3, "###"),
			lex(token.TEXT, " Trees:"),
		),
		emptyLine(),
		emptyLine(),
		line(
			lex(token.BUL_POINT, "."),
			lex(token.TEXT, " Burnable "),
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, "(Wildfires)"),
			lex(token.NEGATIVE, "-"),
		),
		line(
			lex(token.BUL_POINT, "."),
			lex(token.TEXT, " Central to "),
			lex(token.STRONG, "*"),
			lex(token.TEXT, "ecosystems"),
			lex(token.STRONG, "*"),
		),
		line(
			lex(token.BUL_POINT, "."),
			lex(token.POSITIVE, "+"),
			lex(token.TEXT, " Fun to climb"),
		),
		line(
			lex(token.BUL_POINT, "."),
			lex(token.NEGATIVE, "-"),
			lex(token.TEXT, " Can fall over"),
		),
		emptyLine(),
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestLines_2(t *testing.T) {

	line := func(lxs ...token.Lexeme) []token.Lexeme { return lxs }
	emptyLine := func() []token.Lexeme { return []token.Lexeme{} }

	in := `
> I aten't ded
A quote by whom?
`
	exp := [][]token.Lexeme{
		emptyLine(),
		line(
			lex(token.QUOTE, ">"),
			lex(token.TEXT, " I aten't ded"),
		),
		line(
			lex(token.TEXT, "A quote by whom?"),
		),
		emptyLine(),
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}
