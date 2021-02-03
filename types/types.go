package types

import (
	"strings"
	"time"
)

type PhraseType int

const (
	UNDEFINED PhraseType = iota

	TITLE     // #
	TOPIC     // ##
	SUB_TOPIC // ###

	LIST_ITEM     // .
	NUM_LIST_ITEM // 1.

	QUOTE    // "" -
	QUESTION // ?

	POS_PHRASE  // +
	NEG_PHRASE  // -
	STR_PHRASE  // *
	KEY_PRASE   // {}
	CODE_PHRASE // ``

	PLACE    // &
	DATETIME // @
)

func (t PhraseType) String() string {
	switch t {
	case TITLE:
		return "TITLE"
	case TOPIC:
		return "TOPIC"
	case SUB_TOPIC:
		return "SUB_TOPIC"

	case LIST_ITEM:
		return "LIST_ITEM"
	case NUM_LIST_ITEM:
		return "NUM_LIST_ITEM"

	case QUOTE:
		return "QUOTE"
	case QUESTION:
		return "QUESTION"

	case POS_PHRASE:
		return "POS_PHRASE"
	case NEG_PHRASE:
		return "NEG_PHRASE"
	case STR_PHRASE:
		return "STR_PHRASE"
	case KEY_PRASE:
		return "KEY_PRASE"
	case CODE_PHRASE:
		return "CODE_PHRASE"

	case PLACE:
		return "PLACE"
	case DATETIME:
		return "DATETIME"
	}

	return "UNDEFINED"
}

type (
	Phrase interface {
		Type() PhraseType
		String() string
		phrase()
	}

	Notes struct {
		lines []Phrase
	}

	NodePhrase struct {
		PhraseType
		phrases []Phrase
	}

	TextPhrase struct {
		PhraseType
		text string
	}

	TimePhrase struct {
		PhraseType
		text string
		time time.Time
	}
)

func MakeNotes(lines []Phrase) Notes { return Notes{lines: lines} }
func (n Notes) Lines() []Phrase      { return n.lines }
func (n Notes) String() string       { return concat(n.lines, "\n") }

func MakeNode(phrases []Phrase) NodePhrase { return NodePhrase{phrases: phrases} }
func (n NodePhrase) phrase()               {}
func (p NodePhrase) Type() PhraseType      { return p.PhraseType }
func (p NodePhrase) Phrases() []Phrase     { return p.phrases }
func (p NodePhrase) String() string        { return concat(p.phrases, " ") }

func MakeText(text string) TextPhrase { return TextPhrase{text: text} }
func (n TextPhrase) phrase()          {}
func (p TextPhrase) Type() PhraseType { return p.PhraseType }
func (p TextPhrase) Text() string     { return p.text }
func (p TextPhrase) String() string   { return p.text }

func MakeTime(text string, time time.Time) TimePhrase {
	return TimePhrase{text: text, time: time}
}
func (n TimePhrase) phrase()          {}
func (p TimePhrase) Type() PhraseType { return p.PhraseType }
func (p TimePhrase) Text() string     { return p.text }
func (p TimePhrase) Time() time.Time  { return p.time }
func (p TimePhrase) String() string   { return p.text }

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

func NotesString(notes Notes) string {
	return PhrasesString(0, notes.lines)
}

func PhrasesString(indent int, phrases []Phrase) string {
	s := ""
	for _, p := range phrases {
		s += PhraseString(indent, p)
	}
	s += "\n"
	return s
}

func PhraseString(indent int, node interface{}) string {

	s := strings.Repeat(" ", indent)

	switch v := node.(type) {
	case NodePhrase:
		return s + v.Type().String() +
			PhrasesString(indent+2, v.phrases)
	case TextPhrase:
		return s + v.Type().String()
	case TimePhrase:
		return s + v.Type().String()
	default:
		panic("Unknown phrase type")
	}
}
