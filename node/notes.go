package node

import (
	"strings"
	//"unicode"
)

type Notes []Node

func RemoveExtraLines(n Notes) Notes {

	r := []Node{}
	prevEmpty := false

	for _, l := range n {
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

	hubString := func(ns []Node) {
		for _, sub := range ns {
			fmtNodeString(sb, sub)
		}
	}

	writeGroup := func(prefix string, v interface{}, suffix string) {
		sb.WriteString(prefix)
		if ns, ok := v.(Parent); ok {
			hubString(ns.Nodes())
		} else {
			s := v.(Node).Text()
			sb.WriteString(s)
		}
		sb.WriteString(suffix)
	}

	switch v := n.(type) {
	case Phrase:
		writeGroup("", v, "")
	case H1:
		writeGroup("#", v, "")
	case H2:
		writeGroup("##", v, "")
	case H3:
		writeGroup("###", v, "")
	case Quote:
		writeGroup(">", v, "")

	case FmtLine:
		writeGroup("", v, "")
	case BulPoint:
		writeGroup(".", v, "")
	case NumPoint:
		writeGroup(v.Num+".", v, "")

	case KeyPhrase:
		writeGroup("**", v, "**")
	case Positive:
		writeGroup("+", v, "+")
	case Negative:
		writeGroup("-", v, "-")
	case Strong:
		writeGroup("*", v, "*")
	case Snippet:
		writeGroup("`", v, "`")
	}
}
