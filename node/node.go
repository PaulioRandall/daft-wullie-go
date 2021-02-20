package node

import (
	"strings"
)

type (
	// Node is satisfied by all concrete node types. It represents a single
	// annotation in a tree of text annotations.
	Node interface {
		Text() string // Text joins all of the annotated text
		Name() string // Name returns the name of the node type
	}

	// Parent is a node that contains other nodes.
	Parent struct{ Nodes []Node }

	Phrase    struct{ Txt string }
	EmptyLine struct{}

	Snippet struct{ Txt string }

	H1       struct{ Parent }
	H2       struct{ Parent }
	H3       struct{ Parent }
	TextLine struct{ Parent }
	BulPoint struct{ Parent }
	NumPoint struct{ Parent }
	Quote    struct{ Parent }

	KeyPhrase struct{ Parent }
	Positive  struct{ Parent }
	Negative  struct{ Parent }
	Strong    struct{ Parent }
)

func MakePhrase(text string) Phrase   { return Phrase{Txt: text} }
func MakeEmptyLine() EmptyLine        { return EmptyLine{} }
func MakeSnippet(text string) Snippet { return Snippet{Txt: text} }

func MakeH1(nodes ...Node) H1             { return H1{makeParent(nodes)} }
func MakeH2(nodes ...Node) H2             { return H2{makeParent(nodes)} }
func MakeH3(nodes ...Node) H3             { return H3{makeParent(nodes)} }
func MakeFmtLine(nodes ...Node) TextLine  { return TextLine{makeParent(nodes)} }
func MakeBulPoint(nodes ...Node) BulPoint { return BulPoint{makeParent(nodes)} }
func MakeNumPoint(nodes ...Node) NumPoint { return NumPoint{makeParent(nodes)} }
func MakeQuote(nodes ...Node) Quote       { return Quote{makeParent(nodes)} }

func MakeKeyPhrase(nodes ...Node) KeyPhrase { return KeyPhrase{makeParent(nodes)} }
func MakePositive(nodes ...Node) Positive   { return Positive{makeParent(nodes)} }
func MakeNegative(nodes ...Node) Negative   { return Negative{makeParent(nodes)} }
func MakeStrong(nodes ...Node) Strong       { return Strong{makeParent(nodes)} }

func makeParent(nodes []Node) Parent {
	if nodes == nil {
		return Parent{Nodes: []Node{}}
	}
	return Parent{Nodes: nodes}
}

func (n Phrase) Text() string    { return n.Txt }
func (n EmptyLine) Text() string { return "\n" }
func (n Snippet) Text() string   { return n.Txt }
func (n Parent) Text() string {
	sb := strings.Builder{}
	for _, c := range n.Nodes {
		sb.WriteString(c.Text())
	}
	return sb.String()
}

func (n Parent) Children() []Node { return n.Nodes }

func (n Phrase) Name() string    { return "Phrase" }
func (n EmptyLine) Name() string { return "EmptyLine" }
func (n Snippet) Name() string   { return "Snippet" }
func (n H1) Name() string        { return "H1" }
func (n H2) Name() string        { return "H2" }
func (n H3) Name() string        { return "H3" }
func (n TextLine) Name() string  { return "TextLine" }
func (n BulPoint) Name() string  { return "BulPoint" }
func (n NumPoint) Name() string  { return "NumPoint" }
func (n Quote) Name() string     { return "Quote" }
func (n KeyPhrase) Name() string { return "KeyPhrase" }
func (n Positive) Name() string  { return "Positive" }
func (n Negative) Name() string  { return "Negative" }
func (n Strong) Name() string    { return "Strong" }

func _enforceTypes() {

	type par interface {
		Children() []Node
	}

	var (
		n Node
		p par
	)

	n = Phrase{}
	n = EmptyLine{}

	n, p = H1{}, H1{}
	n, p = H2{}, H2{}
	n, p = H3{}, H3{}

	n = Snippet{}

	n, p = TextLine{}, TextLine{}
	n, p = BulPoint{}, BulPoint{}
	n, p = NumPoint{}, NumPoint{}
	n, p = Quote{}, Quote{}

	n, p = KeyPhrase{}, KeyPhrase{}
	n, p = Positive{}, Positive{}
	n, p = Negative{}, Negative{}
	n, p = Strong{}, Strong{}

	_, _ = n, p
}
