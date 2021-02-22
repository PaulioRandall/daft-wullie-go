package ast

type NodeType string

const (
	Undefined        NodeType = ""
	_lineNodeTypes            = "_lineNodeTypes"
	Topic                     = "Topic"
	SubTopic                  = "SubTopic"
	BulPoint                  = "BulPoint"
	SubBulPoint               = "SubBulPoint"
	NumPoint                  = "NumPoint"
	SubNumPoint               = "SubNumPoint"
	TextLine                  = "TextLine"
	EmptyLine                 = "EmptyLine"
	_phraseNodeTypes          = "_phraseNodeTypes"
	Text                      = "Text"
	KeyPhrase                 = "KeyPhrase"
	Positive                  = "Positive"
	Negative                  = "Negative"
	Strong                    = "Strong"
	Quote                     = "Quote"
	Artifact                  = "Artifact"
	Snippet                   = "Snippet"
)

func (nt NodeType) IsLineNode() bool {
	return nt > _lineNodeTypes && nt < _phraseNodeTypes
}

func (nt NodeType) IsPhraseNode() bool {
	return nt > _phraseNodeTypes
}

func (nt NodeType) String() string {
	return string(nt)
}
