package token

type (
	Token  int
	Lexeme struct {
		Token
		Val string
	}
)

const (
	UNDEFINED Token = iota
	TEXT
	TOPIC
	SUB_TOPIC
	HEADING
	BUL_POINT
	NUM_POINT
	QUOTE
	ESCAPE
	POSITIVE
	NEGATIVE
	STRONG
	KEY_PHRASE
	SNIPPET
	QUESTION
)

func (tk Token) String() string {
	switch tk {
	case TEXT:
		return "TEXT"

	case TOPIC:
		return "TOPIC"
	case SUB_TOPIC:
		return "SUB_TOPIC"
	case HEADING:
		return "HEADING"

	case BUL_POINT:
		return "BUL_POINT"
	case NUM_POINT:
		return "NUM_POINT"
	case QUOTE:
		return "QUOTE"

	case ESCAPE:
		return "ESCAPE"

	case KEY_PHRASE:
		return "KEY_PHRASE"
	case STRONG:
		return "STRONG"
	case POSITIVE:
		return "POSITIVE"
	case NEGATIVE:
		return "NEGATIVE"
	case SNIPPET:
		return "SNIPPET"

	case QUESTION:
		return "QUESTION"

	default:
		return "UNDEFINED"
	}
}
