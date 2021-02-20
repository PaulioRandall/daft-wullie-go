// Package ast defines the node types that may appear in an abstract syntax
// tree of a line of text.
package ast

import (
	"strings"
)

type (
	// Node represents node in an AST.
	Node interface {
		Text() string
		Name() string
	}

	// LineNode represents the root of an AST.
	LineNode interface {
		Node
		lineNode()
	}

	// PhraseNode represents any node not at the root of an AST.
	PhraseNode interface {
		Node
		phraseNode()
	}

	EmptyLine struct{}
	Phrase    struct{ Txt string }
	Snippet   struct{ Txt string }

	Parent struct{ Nodes []PhraseNode }

	H1        struct{ Parent }
	H2        struct{ Parent }
	H3        struct{ Parent }
	BulPoint  struct{ Parent }
	NumPoint  struct{ Parent }
	Quote     struct{ Parent }
	TextLine  struct{ Parent }
	KeyPhrase struct{ Parent }
	Positive  struct{ Parent }
	Negative  struct{ Parent }
	Strong    struct{ Parent }
)

func MakeEmptyLine() EmptyLine        { return EmptyLine{} }
func MakePhrase(text string) Phrase   { return Phrase{text} }
func MakeSnippet(text string) Snippet { return Snippet{text} }

func MakeH1(nodes ...PhraseNode) H1             { return H1{makeParent(nodes)} }
func MakeH2(nodes ...PhraseNode) H2             { return H2{makeParent(nodes)} }
func MakeH3(nodes ...PhraseNode) H3             { return H3{makeParent(nodes)} }
func MakeTextLine(nodes ...PhraseNode) TextLine { return TextLine{makeParent(nodes)} }
func MakeBulPoint(nodes ...PhraseNode) BulPoint { return BulPoint{makeParent(nodes)} }
func MakeNumPoint(nodes ...PhraseNode) NumPoint { return NumPoint{makeParent(nodes)} }
func MakeQuote(nodes ...PhraseNode) Quote       { return Quote{makeParent(nodes)} }

func MakeKeyPhrase(nodes ...PhraseNode) KeyPhrase { return KeyPhrase{makeParent(nodes)} }
func MakePositive(nodes ...PhraseNode) Positive   { return Positive{makeParent(nodes)} }
func MakeNegative(nodes ...PhraseNode) Negative   { return Negative{makeParent(nodes)} }
func MakeStrong(nodes ...PhraseNode) Strong       { return Strong{makeParent(nodes)} }

func makeParent(nodes []PhraseNode) Parent {
	if nodes == nil {
		return Parent{Nodes: []PhraseNode{}}
	}
	return Parent{Nodes: nodes}
}

func (n EmptyLine) lineNode() {}
func (n H1) lineNode()        {}
func (n H2) lineNode()        {}
func (n H3) lineNode()        {}
func (n TextLine) lineNode()  {}
func (n BulPoint) lineNode()  {}
func (n NumPoint) lineNode()  {}
func (n Quote) lineNode()     {}

func (n Snippet) phraseNode()   {}
func (n Phrase) phraseNode()    {}
func (n KeyPhrase) phraseNode() {}
func (n Positive) phraseNode()  {}
func (n Negative) phraseNode()  {}
func (n Strong) phraseNode()    {}

func (n EmptyLine) Text() string { return "\n" }
func (n Phrase) Text() string    { return n.Txt }
func (n Snippet) Text() string   { return n.Txt }
func (n Parent) Text() string {
	sb := strings.Builder{}
	for _, c := range n.Nodes {
		sb.WriteString(c.Text())
	}
	return sb.String()
}

func (n Parent) Children() []PhraseNode { return n.Nodes }

func (n EmptyLine) Name() string { return "EmptyLine" }
func (n Phrase) Name() string    { return "Phrase" }
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

	type par interface{ Children() []PhraseNode }

	var (
		ln LineNode
		pn PhraseNode
		p  par
	)

	ln = EmptyLine{}
	pn = Phrase{}
	pn = Snippet{}

	ln, p = H1{}, H1{}
	ln, p = H2{}, H2{}
	ln, p = H3{}, H3{}

	ln, p = TextLine{}, TextLine{}
	ln, p = BulPoint{}, BulPoint{}
	ln, p = NumPoint{}, NumPoint{}
	ln, p = Quote{}, Quote{}

	pn, p = KeyPhrase{}, KeyPhrase{}
	pn, p = Positive{}, Positive{}
	pn, p = Negative{}, Negative{}
	pn, p = Strong{}, Strong{}

	_, _, _ = ln, pn, p
}
