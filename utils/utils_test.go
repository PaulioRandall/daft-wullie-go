package utils

import (
	"testing"

	"github.com/PaulioRandall/daft-wullie-go/node"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicateLines_1(t *testing.T) {

	//
	//
	//
	in := []node.Node{
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
	}

	exp := []node.Node{
		node.MakeEmptyLine(),
	}

	act := RemoveDuplicateLines(in)
	require.Equal(t, exp, act)
}

func TestRemoveDuplicateLines_2(t *testing.T) {

	//
	//
	// # Tree
	//
	//
	in := []node.Node{
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeH1("Tree"),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
	}

	exp := []node.Node{
		node.MakeEmptyLine(),
		node.MakeH1("Tree"),
		node.MakeEmptyLine(),
	}

	act := RemoveDuplicateLines(in)
	require.Equal(t, exp, act)
}

func TestRemoveDuplicateLines_3(t *testing.T) {

	// # H1
	//
	//
	//
	// ## H2
	//
	//
	//
	// ## H3
	in := []node.Node{
		node.MakeH1("H1"),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeH2("H2"),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeEmptyLine(),
		node.MakeH3("H3"),
	}

	exp := []node.Node{
		node.MakeH1("H1"),
		node.MakeEmptyLine(),
		node.MakeH2("H2"),
		node.MakeEmptyLine(),
		node.MakeH3("H3"),
	}

	act := RemoveDuplicateLines(in)
	require.Equal(t, exp, act)
}

/*
func TestTrimSpaces_1(t *testing.T) {

	// # H1
	// #H2
	// # H3
	in := []node.Node{
		node.MakeH1(" H1"),
		node.MakeH2("H2 "),
		node.MakeH3(" H3 "),
	}

	exp := []node.Node{
		node.MakeH1("H1"),
		node.MakeH2("H2"),
		node.MakeH3("H3"),
	}

	act := TrimSpaces(in)
	require.Equal(t, exp, act)
}
*/
