package node

import (
	"strings"
)

type (
	Node interface {
		Text() string
		node()
	}

	Parent interface {
		Node
		Nodes() []Node
	}

	Phrase    struct{ M_Text string }
	HubNode   struct{ M_Nodes []Node }
	EmptyLine struct{}

	H1 struct{ Phrase }
	H2 struct{ Phrase }
	H3 struct{ Phrase }

	Quote struct{ Phrase }

	FmtLine  struct{ HubNode }
	BulPoint struct{ HubNode }
	NumPoint struct {
		Num string
		HubNode
	}

	KeyPhrase struct{ HubNode }
	Positive  struct{ HubNode }
	Negative  struct{ HubNode }
	Strong    struct{ HubNode }
	Snippet   struct{ HubNode }
)

func neverNil(nodes []Node) []Node {
	if nodes == nil {
		return []Node{}
	}
	return nodes
}

func MakePhrase(text string) Phrase    { return Phrase{M_Text: text} }
func MakeHubNode(nodes []Node) HubNode { return HubNode{M_Nodes: neverNil(nodes)} }
func MakeEmptyLine() EmptyLine         { return EmptyLine{} }

func MakeH1(text string) H1 { return H1{Phrase: MakePhrase(text)} }
func MakeH2(text string) H2 { return H2{Phrase: MakePhrase(text)} }
func MakeH3(text string) H3 { return H3{Phrase: MakePhrase(text)} }

func MakeQuote(text string) Quote { return Quote{Phrase: MakePhrase(text)} }

func MakeFmtLine(nodes ...Node) FmtLine   { return FmtLine{HubNode: MakeHubNode(nodes)} }
func MakeBulPoint(nodes ...Node) BulPoint { return BulPoint{HubNode: MakeHubNode(nodes)} }
func MakeNumPoint(num string, nodes ...Node) NumPoint {
	return NumPoint{Num: num, HubNode: MakeHubNode(nodes)}
}

func MakeKeyPhrase(nodes ...Node) KeyPhrase { return KeyPhrase{HubNode: MakeHubNode(nodes)} }
func MakePositive(nodes ...Node) Positive   { return Positive{HubNode: MakeHubNode(nodes)} }
func MakeNegative(nodes ...Node) Negative   { return Negative{HubNode: MakeHubNode(nodes)} }
func MakeStrong(nodes ...Node) Strong       { return Strong{HubNode: MakeHubNode(nodes)} }
func MakeSnippet(nodes ...Node) Snippet     { return Snippet{HubNode: MakeHubNode(nodes)} }

func (p Phrase) node()    {}
func (p HubNode) node()   {}
func (p EmptyLine) node() {}

func (p Phrase) Text() string    { return p.M_Text }
func (p EmptyLine) Text() string { return "" }
func (p HubNode) Text() string {
	sb := strings.Builder{}
	for _, n := range p.M_Nodes {
		sb.WriteString(n.Text())
	}
	return sb.String()
}

func (p HubNode) Nodes() []Node { return p.M_Nodes }

func _enforceTypes() {

	var (
		n Node
		p Parent
	)

	n = Phrase{}
	n = EmptyLine{}

	n = H1{}
	n = H2{}
	n = H3{}

	n = Quote{}

	n, p = FmtLine{}, FmtLine{}
	n, p = BulPoint{}, BulPoint{}
	n, p = NumPoint{}, NumPoint{}

	n, p = KeyPhrase{}, KeyPhrase{}
	n, p = Positive{}, Positive{}
	n, p = Negative{}, Negative{}
	n, p = Strong{}, Strong{}
	n, p = Snippet{}, Snippet{}

	_, _ = n, p
}
