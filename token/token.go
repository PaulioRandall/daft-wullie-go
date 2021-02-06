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
	H1
	H2
	H3
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

	case H1:
		return "H1"
	case H2:
		return "H2"
	case H3:
		return "H3"

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
