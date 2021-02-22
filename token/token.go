package token

type (
	// Token represents a type of token in a text string.
	Token int

	// Lexeme couples a value from a text string with its token type.
	Lexeme struct {
		Token
		Val string
	}
)

const (
	Undefined Token = iota
	Text
	Topic
	SubTopic
	BulPoint
	NumPoint
	Quote
	Escape
	KeyPhrase
	Positive
	Negative
	Strong
	Snippet
)

// String returns the token's human readable string representation.
func (tk Token) String() string {
	switch tk {
	case Text:
		return "Text"

	case Topic:
		return "Topic"
	case SubTopic:
		return "SubTopic"

	case BulPoint:
		return "BulPoint"
	case NumPoint:
		return "NumPoint"

	case Snippet:
		return "Snippet"
	case Escape:
		return "Escape"

	case KeyPhrase:
		return "KeyPhrase"
	case Positive:
		return "Positive"
	case Negative:
		return "Negative"
	case Strong:
		return "Strong"
	case Quote:
		return "Quote"

	default:
		return "Undefined"
	}
}
