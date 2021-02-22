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
			lex(token.Text, "#daft"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestEscape_2(t *testing.T) {

	in := `\da\ft`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.Text, `daft`),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestEscape_3(t *testing.T) {

	in := `da\\ft`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.Text, `da\ft`),
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
			lex(token.Text, "  Topic  "),
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
			lex(token.Text, " Sub topic"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestBulletPoint_1(t *testing.T) {

	in := `. Point`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.BulPoint, "."),
			lex(token.Text, " Point"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNumberPoint_1(t *testing.T) {

	in := `! Point`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.NumPoint, "!"),
			lex(token.Text, " Point"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestQuote_1(t *testing.T) {

	in := `> Fly high through apocalypse skies`
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.Quote, ">"),
			lex(token.Text, " Fly high through apocalypse skies"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNodes_1(t *testing.T) {

	in := "**+-*`"
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.KeyPhrase, "**"),
			lex(token.Positive, "+"),
			lex(token.Negative, "-"),
			lex(token.Strong, "*"),
			lex(token.Snippet, "`"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestNodes_2(t *testing.T) {

	in := "* +positive+ and -negative- *"
	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.Strong, "*"),
			lex(token.Text, " "),
			lex(token.Positive, "+"),
			lex(token.Text, "positive"),
			lex(token.Positive, "+"),
			lex(token.Text, " and "),
			lex(token.Negative, "-"),
			lex(token.Text, "negative"),
			lex(token.Negative, "-"),
			lex(token.Text, " "),
			lex(token.Strong, "*"),
		},
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}

func TestLines_1(t *testing.T) {

	line := func(lxs ...token.Lexeme) []token.Lexeme { return lxs }
	emptyLine := func() []token.Lexeme { return []token.Lexeme{} }

	in := `
## Trees:


. Burnable -(Wildfires)-
. Central to *ecosystems*
.+ Fun to climb
.- Can fall over
`
	exp := [][]token.Lexeme{
		emptyLine(),
		line(
			lex(token.H2, "##"),
			lex(token.Text, " Trees:"),
		),
		emptyLine(),
		emptyLine(),
		line(
			lex(token.BulPoint, "."),
			lex(token.Text, " Burnable "),
			lex(token.Negative, "-"),
			lex(token.Text, "(Wildfires)"),
			lex(token.Negative, "-"),
		),
		line(
			lex(token.BulPoint, "."),
			lex(token.Text, " Central to "),
			lex(token.Strong, "*"),
			lex(token.Text, "ecosystems"),
			lex(token.Strong, "*"),
		),
		line(
			lex(token.BulPoint, "."),
			lex(token.Positive, "+"),
			lex(token.Text, " Fun to climb"),
		),
		line(
			lex(token.BulPoint, "."),
			lex(token.Negative, "-"),
			lex(token.Text, " Can fall over"),
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
			lex(token.Quote, ">"),
			lex(token.Text, " I aten't ded"),
		),
		line(
			lex(token.Text, "A quote by whom?"),
		),
		emptyLine(),
	}

	act := ScanAll(in)
	require.Equal(t, exp, act)
}
