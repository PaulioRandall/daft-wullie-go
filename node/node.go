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
	Parent interface {
		Node
		Children() []Node
	}

	Phrase    struct{ Txt string }
	EmptyLine struct{}

	Snippet struct{ Txt string }

	H1       struct{ Nodes []Node }
	H2       struct{ Nodes []Node }
	H3       struct{ Nodes []Node }
	FmtLine  struct{ Nodes []Node }
	BulPoint struct{ Nodes []Node }
	NumPoint struct {
		Num   string
		Nodes []Node
	}
	Quote struct{ Nodes []Node }

	KeyPhrase struct{ Nodes []Node }
	Positive  struct{ Nodes []Node }
	Negative  struct{ Nodes []Node }
	Strong    struct{ Nodes []Node }
)

func orEmpty(nodes []Node) []Node {
	if nodes == nil {
		return []Node{}
	}
	return nodes
}

func MakePhrase(text string) Phrase { return Phrase{Txt: text} }
func MakeEmptyLine() EmptyLine      { return EmptyLine{} }

func MakeSnippet(text string) Snippet { return Snippet{Txt: text} }

func MakeH1(nodes ...Node) H1             { return H1{Nodes: orEmpty(nodes)} }
func MakeH2(nodes ...Node) H2             { return H2{Nodes: orEmpty(nodes)} }
func MakeH3(nodes ...Node) H3             { return H3{Nodes: orEmpty(nodes)} }
func MakeFmtLine(nodes ...Node) FmtLine   { return FmtLine{Nodes: orEmpty(nodes)} }
func MakeBulPoint(nodes ...Node) BulPoint { return BulPoint{Nodes: orEmpty(nodes)} }
func MakeNumPoint(num string, nodes ...Node) NumPoint {
	return NumPoint{Num: num, Nodes: orEmpty(nodes)}
}
func MakeQuote(nodes ...Node) Quote { return Quote{Nodes: nodes} }

func MakeKeyPhrase(nodes ...Node) KeyPhrase { return KeyPhrase{Nodes: orEmpty(nodes)} }
func MakePositive(nodes ...Node) Positive   { return Positive{Nodes: orEmpty(nodes)} }
func MakeNegative(nodes ...Node) Negative   { return Negative{Nodes: orEmpty(nodes)} }
func MakeStrong(nodes ...Node) Strong       { return Strong{Nodes: orEmpty(nodes)} }

func (n Phrase) Text() string    { return n.Txt }
func (n EmptyLine) Text() string { return "\n" }
func (n Snippet) Text() string   { return n.Txt }
func (n H1) Text() string        { return joinTexts(n.Nodes) }
func (n H2) Text() string        { return joinTexts(n.Nodes) }
func (n H3) Text() string        { return joinTexts(n.Nodes) }
func (n FmtLine) Text() string   { return joinTexts(n.Nodes) }
func (n BulPoint) Text() string  { return joinTexts(n.Nodes) }
func (n NumPoint) Text() string  { return joinTexts(n.Nodes) }
func (n Quote) Text() string     { return joinTexts(n.Nodes) }
func (n KeyPhrase) Text() string { return joinTexts(n.Nodes) }
func (n Positive) Text() string  { return joinTexts(n.Nodes) }
func (n Negative) Text() string  { return joinTexts(n.Nodes) }
func (n Strong) Text() string    { return joinTexts(n.Nodes) }

func (n Phrase) Name() string    { return "Phrase" }
func (n EmptyLine) Name() string { return "EmptyLine" }
func (n Snippet) Name() string   { return "Snippet" }
func (n H1) Name() string        { return "H1" }
func (n H2) Name() string        { return "H2" }
func (n H3) Name() string        { return "H3" }
func (n FmtLine) Name() string   { return "FmtLine" }
func (n BulPoint) Name() string  { return "BulPoint" }
func (n NumPoint) Name() string  { return "NumPoint" }
func (n Quote) Name() string     { return "Quote" }
func (n KeyPhrase) Name() string { return "KeyPhrase" }
func (n Positive) Name() string  { return "Positive" }
func (n Negative) Name() string  { return "Negative" }
func (n Strong) Name() string    { return "Strong" }

func (n H1) Children() []Node        { return n.Nodes }
func (n H2) Children() []Node        { return n.Nodes }
func (n H3) Children() []Node        { return n.Nodes }
func (n FmtLine) Children() []Node   { return n.Nodes }
func (n BulPoint) Children() []Node  { return n.Nodes }
func (n NumPoint) Children() []Node  { return n.Nodes }
func (n Quote) Children() []Node     { return n.Nodes }
func (n KeyPhrase) Children() []Node { return n.Nodes }
func (n Positive) Children() []Node  { return n.Nodes }
func (n Negative) Children() []Node  { return n.Nodes }
func (n Strong) Children() []Node    { return n.Nodes }

func joinTexts(nodes []Node) string {
	sb := strings.Builder{}
	for _, n := range nodes {
		sb.WriteString(n.Text())
	}
	return sb.String()
}

func _enforceTypes() {

	var (
		n Node
		p Parent
	)

	n = Phrase{}
	n = EmptyLine{}

	n, p = H1{}, H1{}
	n, p = H2{}, H2{}
	n, p = H3{}, H3{}

	n = Snippet{}

	n, p = FmtLine{}, FmtLine{}
	n, p = BulPoint{}, BulPoint{}
	n, p = NumPoint{}, NumPoint{}
	n, p = Quote{}, Quote{}

	n, p = KeyPhrase{}, KeyPhrase{}
	n, p = Positive{}, Positive{}
	n, p = Negative{}, Negative{}
	n, p = Strong{}, Strong{}

	_, _ = n, p
}
