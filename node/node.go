package node

type (
	Node interface {
		node()
	}

	TextNode interface {
		Node
		Text() string
	}

	ParentNode interface {
		Node
		Nodes() []Node
	}

	Phrase  struct{ M_Text string }
	HubNode struct{ M_Nodes []Node }
	Empty   struct{}

	H1 struct{ Phrase }
	H2 struct{ Phrase }
	H3 struct{ Phrase }

	Quote struct{ Phrase }

	FmtLine  struct{ HubNode }
	BulPoint struct{ HubNode }
	NumPoint struct{ HubNode }

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
func MakeEmptyLine() Empty             { return Empty{} }

func MakeH1(text string) H1 { return H1{Phrase: MakePhrase(text)} }
func MakeH2(text string) H2 { return H2{Phrase: MakePhrase(text)} }
func MakeH3(text string) H3 { return H3{Phrase: MakePhrase(text)} }

func MakeQuote(text string) Quote { return Quote{Phrase: MakePhrase(text)} }

func MakeFmtLine(nodes ...Node) FmtLine   { return FmtLine{HubNode: MakeHubNode(nodes)} }
func MakeBulPoint(nodes ...Node) BulPoint { return BulPoint{HubNode: MakeHubNode(nodes)} }
func MakeNumPoint(nodes ...Node) NumPoint { return NumPoint{HubNode: MakeHubNode(nodes)} }

func MakeKeyPhrase(nodes ...Node) KeyPhrase { return KeyPhrase{HubNode: MakeHubNode(nodes)} }
func MakePositive(nodes ...Node) Positive   { return Positive{HubNode: MakeHubNode(nodes)} }
func MakeNegative(nodes ...Node) Negative   { return Negative{HubNode: MakeHubNode(nodes)} }
func MakeStrong(nodes ...Node) Strong       { return Strong{HubNode: MakeHubNode(nodes)} }
func MakeSnippet(nodes ...Node) Snippet     { return Snippet{HubNode: MakeHubNode(nodes)} }

func (p Phrase) node()  {}
func (p HubNode) node() {}
func (p Empty) node()   {}

func (p Phrase) Text() string   { return p.M_Text }
func (p HubNode) Nodes() []Node { return p.M_Nodes }
func (p Empty) Text() string    { return "" }

func _enforceTypes() {

	var (
		n  Node
		tn TextNode
		pn ParentNode
	)

	n, tn = Phrase{}, Phrase{}
	n, tn = Empty{}, Empty{}

	n, tn = H1{}, H1{}
	n, tn = H2{}, H2{}
	n, tn = H3{}, H3{}

	n, tn = Quote{}, Quote{}

	n, pn = FmtLine{}, FmtLine{}
	n, pn = BulPoint{}, BulPoint{}
	n, pn = NumPoint{}, NumPoint{}

	n, pn = KeyPhrase{}, KeyPhrase{}
	n, pn = Positive{}, Positive{}
	n, pn = Negative{}, Negative{}
	n, pn = Strong{}, Strong{}
	n, pn = Snippet{}, Snippet{}

	_, _, _ = n, tn, pn
}
