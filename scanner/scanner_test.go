package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func lex(tk Token, val string, line, start, end int) Lexeme {
	return Lexeme{
		Token: tk,
		Pos: Pos{
			Line:  line,
			Start: start,
			End:   end,
		},
		Val: val,
	}
}

func doTestScan(in string) ([][]Lexeme, error) {

	var (
		act  = [][]Lexeme{}
		scan = NewScanLine(in)
		line []Lexeme
		e    error
	)

	for scan != nil {
		if line, scan, e = scan(); e != nil {
			return nil, e
		}
		act = append(act, line)
	}

	return act, nil
}

func TestTitle_1(t *testing.T) {

	exp := [][]Lexeme{
		[]Lexeme{
			lex(TITLE, "#", 1, 2, 3),
			lex(TEXT, "Title", 1, 5, 10),
		},
	}

	act, e := doTestScan(`  #  Title  `)
	require.Nil(t, e, "%+v")
	require.Equal(t, exp, act)
}

func TestTopic_1(t *testing.T) {

	exp := [][]Lexeme{
		[]Lexeme{
			lex(TOPIC, "##", 1, 0, 2),
			lex(TEXT, "Topic", 1, 3, 8),
		},
	}

	act, e := doTestScan(`## Topic`)
	require.Nil(t, e, "%+v")
	require.Equal(t, exp, act)
}

func TestSubTopic_1(t *testing.T) {

	exp := [][]Lexeme{
		[]Lexeme{
			lex(SUB_TOPIC, "###", 1, 0, 3),
			lex(TEXT, "Subtopic", 1, 4, 12),
		},
	}

	act, e := doTestScan(`### Subtopic`)
	require.Nil(t, e, "%+v")
	require.Equal(t, exp, act)
}
