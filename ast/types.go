package ast

type NodeType int

const (
	Undefined NodeType = iota
	_lineNodeTypes
	Topic
	SubTopic
	BulPoint
	NumPoint
	Quote
	Snippet
	TextLine
	EmptyLine
	_phraseNodeTypes
	Text
	KeyPhrase
	Positive
	Negative
	Strong
)

func (nt NodeType) IsLineNode() bool {
	return nt > _lineNodeTypes && nt < _phraseNodeTypes
}

func (nt NodeType) IsPhraseNode() bool {
	return nt > _phraseNodeTypes
}

func (nt NodeType) String() string {
	switch nt {
	case Topic:
		return "Topic"
	case SubTopic:
		return "SubTopic"
	case BulPoint:
		return "BulPoint"
	case NumPoint:
		return "NumPoint"
	case Quote:
		return "Quote"
	case Snippet:
		return "Snippet"
	case TextLine:
		return "TextLine"
	case EmptyLine:
		return "EmptyLine"
	case Text:
		return "Text"
	case KeyPhrase:
		return "KeyPhrase"
	case Positive:
		return "Positive"
	case Negative:
		return "Negative"
	case Strong:
		return "Strong"
	default:
		return "Undefined"
	}
}
