package node

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveExtraLines_1(t *testing.T) {

	//
	//
	//
	in := Notes{
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeEmptyLine(),
	}

	exp := Notes{
		MakeEmptyLine(),
	}

	act := RemoveExtraLines(in)
	require.Equal(t, exp, act)
}

func TestRemoveExtraLines_2(t *testing.T) {

	//
	//
	//
	// # H1
	//
	//
	//
	// ## H2
	//
	//
	//
	// ## H3
	//
	//
	//
	in := Notes{
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeH1("H1"),
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeH2("H2"),
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeH3("H3"),
		MakeEmptyLine(),
		MakeEmptyLine(),
		MakeEmptyLine(),
	}

	exp := Notes{
		MakeEmptyLine(),
		MakeH1("H1"),
		MakeEmptyLine(),
		MakeH2("H2"),
		MakeEmptyLine(),
		MakeH3("H3"),
		MakeEmptyLine(),
	}

	act := RemoveExtraLines(in)
	require.Equal(t, exp, act)
}
