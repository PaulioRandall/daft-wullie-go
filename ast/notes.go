package ast

import (
	"strings"
)

type (
	Notes       []LineNode
	DescendFunc func(n Node, lineNum, depth, orderIdx int)
)

func RemoveExtraLines(notes Notes) Notes {

	r := []LineNode{}
	prevEmpty := false

	for _, l := range notes {
		if prevEmpty {
			if _, ok := l.(EmptyLine); ok {
				continue
			}
		}

		r = append(r, l)
		_, prevEmpty = l.(EmptyLine)
	}

	return Notes(r)
}

func DescendNotes(notes Notes, f DescendFunc) {
	for i, n := range notes {
		descendNode(n, i+1, 0, 0, f)
	}
}

func DecendNode(n Node, f DescendFunc) {
	descendNode(n, 1, 0, 0, f)
}

func descendNode(n Node, lineNum, depth, orderIdx int, f DescendFunc) {
	type par interface {
		Children() []PhraseNode
	}
	f(n, lineNum, depth, orderIdx)
	if v, ok := n.(par); ok {
		descendPhraseNodes(v.Children(), lineNum, depth+1, orderIdx, f)
	}
}

func descendPhraseNodes(ns []PhraseNode, lineNum, depth, orderIdx int, f DescendFunc) {
	for i, n := range ns {
		descendNode(n, lineNum, depth, i, f)
	}
}

func PlainString(notes Notes) string {
	sb := strings.Builder{}
	for _, n := range notes {
		s := strings.TrimSpace(n.Text())
		sb.WriteString(s)
		sb.WriteString("\n")
	}
	return sb.String()
}

func FmtString(notes Notes) string {
	sb := &strings.Builder{}
	for _, n := range notes {
		fmtNodeString(sb, n)
		sb.WriteString("\n")
	}
	return sb.String()
}

func fmtNodeString(sb *strings.Builder, n Node) {

	writeGroup := func(prefix string, v interface{}, suffix string) {
		sb.WriteString(prefix)
		if ns, ok := v.(Parent); ok {
			for _, sub := range ns.Children() {
				fmtNodeString(sb, sub)
			}
		} else {
			s := v.(Node).Text()
			sb.WriteString(s)
		}
		sb.WriteString(suffix)
	}

	switch v := n.(type) {
	case Phrase:
		writeGroup("", v, "")

	case Quote:
		writeGroup(">", v, "")
	case Snippet:
		writeGroup("`", v, "`")

	case H1:
		writeGroup("#", v, "")
	case H2:
		writeGroup("##", v, "")
	case H3:
		writeGroup("###", v, "")

	case TextLine:
		writeGroup("", v, "")
	case BulPoint:
		writeGroup(".", v, "")
	case NumPoint:
		writeGroup("!", v, "")

	case KeyPhrase:
		writeGroup("**", v, "**")
	case Positive:
		writeGroup("+", v, "+")
	case Negative:
		writeGroup("-", v, "-")
	case Strong:
		writeGroup("*", v, "*")
	}
}
