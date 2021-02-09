package utils

import (
	"testing"

	"github.com/PaulioRandall/daft-wullie-go/node"

	"github.com/stretchr/testify/require"
)

func neverNil(nodes []node.Node) []node.Node {
	if nodes == nil {
		return []node.Node{}
	}
	return nodes
}

func emptyLine() node.Empty { return node.Empty{} }

func h1(text string) node.H1 { return node.H1{M_Text: text} }
func h2(text string) node.H2 { return node.H2{M_Text: text} }
func h3(text string) node.H3 { return node.H3{M_Text: text} }

func quote(text string) node.Quote              { return node.Quote{M_Text: text} }
func fmtLine(nodes ...node.Node) node.FmtLine   { return node.FmtLine{M_Nodes: neverNil(nodes)} }
func bulPoint(nodes ...node.Node) node.BulPoint { return node.BulPoint{M_Nodes: neverNil(nodes)} }
func numPoint(nodes ...node.Node) node.NumPoint { return node.NumPoint{M_Nodes: neverNil(nodes)} }

func keyPhrase(nodes ...node.Node) node.KeyPhrase { return node.KeyPhrase{M_Nodes: neverNil(nodes)} }
func positive(nodes ...node.Node) node.Positive   { return node.Positive{M_Nodes: neverNil(nodes)} }
func negative(nodes ...node.Node) node.Negative   { return node.Negative{M_Nodes: neverNil(nodes)} }
func strong(nodes ...node.Node) node.Strong       { return node.Strong{M_Nodes: neverNil(nodes)} }
func snippet(nodes ...node.Node) node.Snippet     { return node.Snippet{M_Nodes: neverNil(nodes)} }
func phrase(text string) node.Phrase              { return node.Phrase{M_Text: text} }

func TestRemoveDuplicateLines_1(t *testing.T) {

	//
	//
	//
	in := []node.Node{
		emptyLine(),
		emptyLine(),
		emptyLine(),
	}

	exp := []node.Node{
		emptyLine(),
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
		emptyLine(),
		emptyLine(),
		h1("Tree"),
		emptyLine(),
		emptyLine(),
	}

	exp := []node.Node{
		emptyLine(),
		h1("Tree"),
		emptyLine(),
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
		h1("H1"),
		emptyLine(),
		emptyLine(),
		emptyLine(),
		h2("H2"),
		emptyLine(),
		emptyLine(),
		emptyLine(),
		h3("H3"),
	}

	exp := []node.Node{
		h1("H1"),
		emptyLine(),
		h2("H2"),
		emptyLine(),
		h3("H3"),
	}

	act := RemoveDuplicateLines(in)
	require.Equal(t, exp, act)
}
