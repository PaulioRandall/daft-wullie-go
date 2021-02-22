// Package ast defines node and its types that may appear in an abstract syntax
// tree of a line of text.
package ast

import (
	"strings"
)

type (
	Node interface {
		Type() NodeType
		Text() string
	}

	Parent interface {
		Node
		Nodes() []Node
	}

	TextNode struct {
		NodeType
		Txt string
	}

	ParentNode struct {
		NodeType
		Children []Node
	}
)

func (n TextNode) Type() NodeType   { return n.NodeType }
func (n ParentNode) Type() NodeType { return n.NodeType }

func (n TextNode) Text() string { return n.Txt }
func (n ParentNode) Text() string {
	sb := strings.Builder{}
	for _, c := range n.Children {
		sb.WriteString(c.Text())
	}
	return sb.String()
}

func (n ParentNode) Nodes() []Node { return n.Children }

func MakeEmptyLine() TextNode       { return makeTextNode(EmptyLine, "") }
func MakeText(s string) TextNode    { return makeTextNode(Text, s) }
func MakeSnippet(s string) TextNode { return makeTextNode(Snippet, s) }

func MakeH1(ns ...Node) ParentNode       { return makeParentNode(H1, ns) }
func MakeH2(ns ...Node) ParentNode       { return makeParentNode(H2, ns) }
func MakeBulPoint(ns ...Node) ParentNode { return makeParentNode(BulPoint, ns) }
func MakeNumPoint(ns ...Node) ParentNode { return makeParentNode(NumPoint, ns) }
func MakeQuote(ns ...Node) ParentNode    { return makeParentNode(Quote, ns) }
func MakeTextLine(ns ...Node) ParentNode { return makeParentNode(TextLine, ns) }

func MakeKeyPhrase(ns ...Node) ParentNode { return makeParentNode(KeyPhrase, ns) }
func MakePositive(ns ...Node) ParentNode  { return makeParentNode(Positive, ns) }
func MakeNegative(ns ...Node) ParentNode  { return makeParentNode(Negative, ns) }
func MakeStrong(ns ...Node) ParentNode    { return makeParentNode(Strong, ns) }

func makeTextNode(nt NodeType, s string) TextNode {
	return TextNode{NodeType: nt, Txt: s}
}

func makeParentNode(nt NodeType, ns []Node) ParentNode {
	if ns == nil {
		return ParentNode{NodeType: nt, Children: []Node{}}
	}
	return ParentNode{NodeType: nt, Children: ns}
}
