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
	EmptyLine struct{}

	H1       struct{ M_Text string }
	H2       struct{ M_Text string }
	H3       struct{ M_Text string }
	Quote    struct{ M_Text string }
	FmtLine  struct{ M_Nodes []Node }
	BulPoint struct{ M_Nodes []Node }
	NumPoint struct {
		Num     string
		M_Nodes []Node
	}
	KeyPhrase struct{ M_Nodes []Node }
	Positive  struct{ M_Nodes []Node }
	Negative  struct{ M_Nodes []Node }
	Strong    struct{ M_Nodes []Node }
	Snippet   struct{ M_Nodes []Node }
)

func orEmpty(nodes []Node) []Node {
	if nodes == nil {
		return []Node{}
	}
	return nodes
}

func MakePhrase(text string) Phrase { return Phrase{M_Text: text} }
func MakeEmptyLine() EmptyLine      { return EmptyLine{} }

func MakeH1(text string) H1       { return H1{M_Text: text} }
func MakeH2(text string) H2       { return H2{M_Text: text} }
func MakeH3(text string) H3       { return H3{M_Text: text} }
func MakeQuote(text string) Quote { return Quote{M_Text: text} }

func MakeFmtLine(nodes ...Node) FmtLine   { return FmtLine{M_Nodes: orEmpty(nodes)} }
func MakeBulPoint(nodes ...Node) BulPoint { return BulPoint{M_Nodes: orEmpty(nodes)} }
func MakeNumPoint(num string, nodes ...Node) NumPoint {
	return NumPoint{Num: num, M_Nodes: orEmpty(nodes)}
}

func MakeKeyPhrase(nodes ...Node) KeyPhrase { return KeyPhrase{M_Nodes: orEmpty(nodes)} }
func MakePositive(nodes ...Node) Positive   { return Positive{M_Nodes: orEmpty(nodes)} }
func MakeNegative(nodes ...Node) Negative   { return Negative{M_Nodes: orEmpty(nodes)} }
func MakeStrong(nodes ...Node) Strong       { return Strong{M_Nodes: orEmpty(nodes)} }
func MakeSnippet(nodes ...Node) Snippet     { return Snippet{M_Nodes: orEmpty(nodes)} }

func (n Phrase) node()    {}
func (n EmptyLine) node() {}
func (n H1) node()        {}
func (n H2) node()        {}
func (n H3) node()        {}
func (n Quote) node()     {}
func (n FmtLine) node()   {}
func (n BulPoint) node()  {}
func (n NumPoint) node()  {}
func (n KeyPhrase) node() {}
func (n Positive) node()  {}
func (n Negative) node()  {}
func (n Strong) node()    {}
func (n Snippet) node()   {}

func (n Phrase) Text() string    { return n.M_Text }
func (n EmptyLine) Text() string { return "\n" }
func (n H1) Text() string        { return n.M_Text }
func (n H2) Text() string        { return n.M_Text }
func (n H3) Text() string        { return n.M_Text }
func (n Quote) Text() string     { return n.M_Text }
func (n FmtLine) Text() string   { return joinTexts(n.M_Nodes) }
func (n BulPoint) Text() string  { return joinTexts(n.M_Nodes) }
func (n NumPoint) Text() string  { return joinTexts(n.M_Nodes) }
func (n KeyPhrase) Text() string { return joinTexts(n.M_Nodes) }
func (n Positive) Text() string  { return joinTexts(n.M_Nodes) }
func (n Negative) Text() string  { return joinTexts(n.M_Nodes) }
func (n Strong) Text() string    { return joinTexts(n.M_Nodes) }
func (n Snippet) Text() string   { return joinTexts(n.M_Nodes) }

func (n FmtLine) Nodes() []Node   { return n.M_Nodes }
func (n BulPoint) Nodes() []Node  { return n.M_Nodes }
func (n NumPoint) Nodes() []Node  { return n.M_Nodes }
func (n KeyPhrase) Nodes() []Node { return n.M_Nodes }
func (n Positive) Nodes() []Node  { return n.M_Nodes }
func (n Negative) Nodes() []Node  { return n.M_Nodes }
func (n Strong) Nodes() []Node    { return n.M_Nodes }
func (n Snippet) Nodes() []Node   { return n.M_Nodes }

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
