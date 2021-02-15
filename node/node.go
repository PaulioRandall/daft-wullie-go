package node

import (
	"strings"
)

type (
	Node interface {
		Text() string
		Name() string
	}

	Parent interface {
		Node
		Nodes() []Node
	}

	Phrase    struct{ M_Text string }
	EmptyLine struct{}

	Quote   struct{ M_Text string }
	Snippet struct{ M_Text string }

	H1       struct{ M_Nodes []Node }
	H2       struct{ M_Nodes []Node }
	H3       struct{ M_Nodes []Node }
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
)

func orEmpty(nodes []Node) []Node {
	if nodes == nil {
		return []Node{}
	}
	return nodes
}

func MakePhrase(text string) Phrase { return Phrase{M_Text: text} }
func MakeEmptyLine() EmptyLine      { return EmptyLine{} }

func MakeQuote(text string) Quote     { return Quote{M_Text: text} }
func MakeSnippet(text string) Snippet { return Snippet{M_Text: text} }

func MakeH1(nodes ...Node) H1             { return H1{M_Nodes: orEmpty(nodes)} }
func MakeH2(nodes ...Node) H2             { return H2{M_Nodes: orEmpty(nodes)} }
func MakeH3(nodes ...Node) H3             { return H3{M_Nodes: orEmpty(nodes)} }
func MakeFmtLine(nodes ...Node) FmtLine   { return FmtLine{M_Nodes: orEmpty(nodes)} }
func MakeBulPoint(nodes ...Node) BulPoint { return BulPoint{M_Nodes: orEmpty(nodes)} }
func MakeNumPoint(num string, nodes ...Node) NumPoint {
	return NumPoint{Num: num, M_Nodes: orEmpty(nodes)}
}

func MakeKeyPhrase(nodes ...Node) KeyPhrase { return KeyPhrase{M_Nodes: orEmpty(nodes)} }
func MakePositive(nodes ...Node) Positive   { return Positive{M_Nodes: orEmpty(nodes)} }
func MakeNegative(nodes ...Node) Negative   { return Negative{M_Nodes: orEmpty(nodes)} }
func MakeStrong(nodes ...Node) Strong       { return Strong{M_Nodes: orEmpty(nodes)} }

func (n Phrase) Text() string    { return n.M_Text }
func (n EmptyLine) Text() string { return "\n" }
func (n Quote) Text() string     { return n.M_Text }
func (n Snippet) Text() string   { return n.M_Text }
func (n H1) Text() string        { return joinTexts(n.M_Nodes) }
func (n H2) Text() string        { return joinTexts(n.M_Nodes) }
func (n H3) Text() string        { return joinTexts(n.M_Nodes) }
func (n FmtLine) Text() string   { return joinTexts(n.M_Nodes) }
func (n BulPoint) Text() string  { return joinTexts(n.M_Nodes) }
func (n NumPoint) Text() string  { return joinTexts(n.M_Nodes) }
func (n KeyPhrase) Text() string { return joinTexts(n.M_Nodes) }
func (n Positive) Text() string  { return joinTexts(n.M_Nodes) }
func (n Negative) Text() string  { return joinTexts(n.M_Nodes) }
func (n Strong) Text() string    { return joinTexts(n.M_Nodes) }

func (n Phrase) Name() string    { return "Phrase" }
func (n EmptyLine) Name() string { return "EmptyLine" }
func (n Quote) Name() string     { return "Quote" }
func (n Snippet) Name() string   { return "Snippet" }
func (n H1) Name() string        { return "H1" }
func (n H2) Name() string        { return "H2" }
func (n H3) Name() string        { return "H3" }
func (n FmtLine) Name() string   { return "FmtLine" }
func (n BulPoint) Name() string  { return "BulPoint" }
func (n NumPoint) Name() string  { return "NumPoint" }
func (n KeyPhrase) Name() string { return "KeyPhrase" }
func (n Positive) Name() string  { return "Positive" }
func (n Negative) Name() string  { return "Negative" }
func (n Strong) Name() string    { return "Strong" }

func (n H1) Nodes() []Node        { return n.M_Nodes }
func (n H2) Nodes() []Node        { return n.M_Nodes }
func (n H3) Nodes() []Node        { return n.M_Nodes }
func (n FmtLine) Nodes() []Node   { return n.M_Nodes }
func (n BulPoint) Nodes() []Node  { return n.M_Nodes }
func (n NumPoint) Nodes() []Node  { return n.M_Nodes }
func (n KeyPhrase) Nodes() []Node { return n.M_Nodes }
func (n Positive) Nodes() []Node  { return n.M_Nodes }
func (n Negative) Nodes() []Node  { return n.M_Nodes }
func (n Strong) Nodes() []Node    { return n.M_Nodes }

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

	n = Quote{}
	n = Snippet{}

	n, p = FmtLine{}, FmtLine{}
	n, p = BulPoint{}, BulPoint{}
	n, p = NumPoint{}, NumPoint{}

	n, p = KeyPhrase{}, KeyPhrase{}
	n, p = Positive{}, Positive{}
	n, p = Negative{}, Negative{}
	n, p = Strong{}, Strong{}

	_, _ = n, p
}
