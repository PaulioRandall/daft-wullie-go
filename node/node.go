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

	Phrase struct{ M_Text string }
	Empty  struct{}

	H1 struct{ M_Text string }
	H2 struct{ M_Text string }
	H3 struct{ M_Text string }

	Quote struct{ M_Text string }

	BulItem struct{ M_Nodes []Node }
	NumItem struct{ M_Nodes []Node }

	KeyPhrase struct{ M_Nodes []Node }
	Positive  struct{ M_Nodes []Node }
	Negative  struct{ M_Nodes []Node }
	Strong    struct{ M_Nodes []Node }
	Snippet   struct{ M_Nodes []Node }

	Question struct{ M_Nodes []Node }
)

func (p Phrase) node()    {}
func (p Empty) node()     {}
func (p H1) node()        {}
func (p H2) node()        {}
func (p H3) node()        {}
func (p Quote) node()     {}
func (p BulItem) node()   {}
func (p NumItem) node()   {}
func (p KeyPhrase) node() {}
func (p Positive) node()  {}
func (p Negative) node()  {}
func (p Strong) node()    {}
func (p Snippet) node()   {}
func (p Question) node()  {}

func (p Phrase) Text() string     { return p.M_Text }
func (p Empty) Text() string      { return "" }
func (p H1) Text() string         { return p.M_Text }
func (p H2) Text() string         { return p.M_Text }
func (p H3) Text() string         { return p.M_Text }
func (p Quote) Text() string      { return p.M_Text }
func (p BulItem) Nodes() []Node   { return p.M_Nodes }
func (p NumItem) Nodes() []Node   { return p.M_Nodes }
func (p KeyPhrase) Nodes() []Node { return p.M_Nodes }
func (p Positive) Nodes() []Node  { return p.M_Nodes }
func (p Negative) Nodes() []Node  { return p.M_Nodes }
func (p Strong) Nodes() []Node    { return p.M_Nodes }
func (p Snippet) Nodes() []Node   { return p.M_Nodes }
func (p Question) Nodes() []Node  { return p.M_Nodes }

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

	n, pn = BulItem{}, BulItem{}
	n, pn = NumItem{}, NumItem{}

	n, pn = KeyPhrase{}, KeyPhrase{}
	n, pn = Positive{}, Positive{}
	n, pn = Negative{}, Negative{}
	n, pn = Strong{}, Strong{}
	n, pn = Snippet{}, Snippet{}

	n, pn = Question{}, Question{}

	_, _, _ = n, tn, pn
}
