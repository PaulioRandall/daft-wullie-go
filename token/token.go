package token

type (
	// Token represents a type of token in a text string.
	Token string

	// Lexeme couples a value from a text string with its token type.
	Lexeme struct {
		Token
		Val string
	}
)

const (
	Undefined Token = ""
	Text            = "Text"
	Topic           = "Topic"
	SubTopic        = "SubTopic"
	BulPoint        = "BulPoint"
	NumPoint        = "NumPoint"
	KeyPhrase       = "KeyPhrase"
	Positive        = "Positive"
	Negative        = "Negative"
	Strong          = "Strong"
	Snippet         = "Snippet"
	Quote           = "Quote"
	Escape          = "Escape"
)

// String returns the token's human readable string representation.
func (tk Token) String() string {
	return string(tk)
}
