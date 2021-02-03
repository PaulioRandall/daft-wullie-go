package types

import (
	"fmt"
	"strings"
	"time"
)

type (
	Position interface {
		GetLine() int
		GetStart() int
		GetEnd() int
	}

	Phrase interface {
		GetPos() Pos
		GetText() string
	}

	NodePhrase interface {
		Phrase
		GetPhrases() []Phrase
	}
)

type (
	Pos struct{ Line, Start, End int }

	Text struct {
		Pos
		Text string
	}

	Node struct {
		Pos
		Phrases []Phrase
	}

	Notes     struct{ Lines []Phrase }
	EmptyLine struct{ Text }

	Title    struct{ Text }
	Topic    struct{ Text }
	SubTopic struct{ Text }

	KeyPhrase struct{ Text }
	Snippet   struct{ Text }

	BulletItem struct{ Node }
	NumberItem struct{ Node }
	Question   struct{ Node }

	Positive struct{ Node }
	Negative struct{ Node }
	Strong   struct{ Node }

	Place struct {
		Pos
		Parts []string
	}

	Time struct {
		Pos
		Text string
		Time time.Time
	}

	Quote struct {
		Pos
		Words Text
		Src   Text
	}
)

func (p Pos) GetLine() int  { return p.Line }
func (p Pos) GetStart() int { return p.Start }
func (p Pos) GetEnd() int   { return p.End }

func (t Text) GetPos() Pos     { return t.Pos }
func (t Text) GetText() string { return t.Text }

func (n Node) GetPos() Pos          { return n.Pos }
func (n Node) GetText() string      { return concat(n.Phrases, " ") }
func (n Node) GetPhrases() []Phrase { return n.Phrases }

func (n Notes) GetLines() []Phrase { return n.Lines }
func (n Notes) GetText() string    { return concat(n.Lines, "\n") }
func (n Notes) String() string     { return concat(n.Lines, "\n") }

func (p Place) GetPos() Pos     { return p.Pos }
func (p Place) GetText() string { return concat(p.Parts, ",") }

func (t Time) GetPos() Pos     { return t.Pos }
func (t Time) GetText() string { return t.Text }

func (q Quote) GetPos() Pos { return q.Pos }
func (q Quote) GetText() string {
	return fmt.Sprintf("\"%s\" - %s", q.Words.GetText(), q.Src.GetText())
}

func DebugNotesString(notes Notes) string {
	return DebugPhrasesString(0, notes.Lines)
}

func DebugPhrasesString(indent int, ps []Phrase) string {
	s := ""
	for _, p := range ps {
		s += DebugPhraseString(indent, p)
	}
	s += "\n"
	return s
}

func DebugPhraseString(indent int, p Phrase) string {

	textStr := func(name string, p Phrase) string {
		return strings.Repeat(" ", indent) + "[" + name + "] " + p.GetText()
	}

	nodeStr := func(name string, n NodePhrase) string {
		return textStr(name, n) + DebugPhrasesString(indent+2, n.GetPhrases())
	}

	switch v := p.(type) {
	case EmptyLine:
		return ""
	case Title:
		return textStr("title", v)

	case Topic:
		return textStr("topic", v)
	case SubTopic:
		return textStr("sub-topic", v)

	case KeyPhrase:
		return textStr("key-phrase", v)
	case Snippet:
		return textStr("snippet", v)

	case BulletItem:
		return nodeStr("bullet", v)
	case NumberItem:
		return nodeStr("numbered", v)
	case Question:
		return nodeStr("question", v)

	case Positive:
		return nodeStr("positive", v)
	case Negative:
		return nodeStr("negative", v)
	case Strong:
		return nodeStr("strong", v)

	case Place:
		return textStr("place", v)
	case Time:
		return textStr("time", v)
	case Quote:
		return textStr("quote", v)

	default:
		panic("Unknown phrase type")
	}
}

func concat(list interface{}, sep string) string {

	type stringer interface {
		String() string
	}

	s := ""
	for i, v := range list.([]interface{}) {
		if i != 0 {
			s += sep
		}
		s += v.(stringer).String()
	}

	return s
}

func _enforceTypes() {

	var pos Position
	var p Phrase
	var n NodePhrase

	pos = Pos{}
	pos, p = Text{}, Text{}
	pos, p, n = Node{}, Node{}, Node{}

	pos, p = EmptyLine{}, EmptyLine{}

	pos, p = Title{}, Title{}
	pos, p = Topic{}, Topic{}
	pos, p = SubTopic{}, SubTopic{}

	pos, p = KeyPhrase{}, KeyPhrase{}
	pos, p = Snippet{}, Snippet{}

	pos, p, n = BulletItem{}, BulletItem{}, BulletItem{}
	pos, p, n = NumberItem{}, NumberItem{}, NumberItem{}
	pos, p, n = Question{}, Question{}, Question{}

	pos, p, n = Positive{}, Positive{}, Positive{}
	pos, p, n = Negative{}, Negative{}, Negative{}
	pos, p, n = Strong{}, Strong{}, Strong{}

	pos, p = Place{}, Place{}
	pos, p = Time{}, Time{}
	pos, p = Quote{}, Quote{}

	_, _, _ = pos, p, n
}
