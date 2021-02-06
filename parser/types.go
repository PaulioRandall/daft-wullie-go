package parser

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
	Parent struct{ M_Nodes []Node }

	H1 struct{ Phrase }
	H2 struct{ Phrase }
	H3 struct{ Phrase }

	Quote struct{ Phrase }

	BulItem struct{ Parent }
	NumItem struct{ Parent }

	KeyPhrase struct{ Parent }
	Positive  struct{ Parent }
	Negative  struct{ Parent }
	Strong    struct{ Parent }
	Snippet   struct{ Parent }

	Question struct{ Parent }
)

func (p Phrase) node() {}
func (p Empty) node()  {}
func (p Parent) node() {}

func (p Phrase) Text() string  { return p.M_Text }
func (p Empty) Text() string   { return "" }
func (p Parent) Nodes() []Node { return p.M_Nodes }

func _enforceTypes() {

	var (
		n  Node
		tn TextNode
		pn ParentNode
	)

	n, tn = Phrase{}, Phrase{}
	n, tn = Empty{}, Empty{}
	n, pn = Parent{}, Parent{}

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
