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

	hubString := func(h HubNode) {
		for _, sub := range h.M_Nodes {
			fmtNodeString(sb, sub)
		}
	}

	writeLine := func(prefix string, n Node) {
		sb.WriteString(prefix)
		sb.WriteString(n.Text())
	}

	writeGroup := func(prefix string, h HubNode, suffix string) {
		sb.WriteString(prefix)
		hubString(h)
		sb.WriteString(suffix)
	}

	switch v := n.(type) {
	case Phrase:
		writeLine("", n)
	case H1:
		writeLine("#", n)
	case H2:
		writeLine("##", n)
	case H3:
		writeLine("###", n)
	case Quote:
		writeLine(">", n)

	case FmtLine:
		hubString(v.HubNode)
	case BulPoint:
		writeGroup(".", v.HubNode, "")
	case NumPoint:
		writeGroup(v.Num+".", v.HubNode, "")

	case KeyPhrase:
		writeGroup("**", v.HubNode, "**")
	case Positive:
		writeGroup("+", v.HubNode, "+")
	case Negative:
		writeGroup("-", v.HubNode, "-")
	case Strong:
		writeGroup("*", v.HubNode, "*")
	case Snippet:
		writeGroup("`", v.HubNode, "`")
	}
}
