// Package ast defines node and its types that may appear in an abstract syntax
// tree of a line of text.
package ast2

import (
	"strings"
)

type (
	Node interface {
		Type() NodeType
		Text() string
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

func MakeTextNode(nt NodeType, s string) TextNode {
	return TextNode{NodeType: nt, Txt: s}
}

func MakeParentNode(nt NodeType, ns ...Node) ParentNode {
	if ns == nil {
		return ParentNode{NodeType: nt, Children: []Node{}}
	}
	return ParentNode{NodeType: nt, Children: ns}
}
