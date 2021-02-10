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

	writeLine := func(prefix string, n Node) {
		sb.WriteString(prefix)
		sb.WriteString(n.Text())
	}

	writeGroup := func(prefix string, ns []Node, suffix string) {
		sb.WriteString(prefix)
		hubString(ns)
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
		hubString(v.M_Nodes)
	case BulPoint:
		writeGroup(".", v.M_Nodes, "")
	case NumPoint:
		writeGroup(v.Num+".", v.M_Nodes, "")

	case KeyPhrase:
		writeGroup("**", v.M_Nodes, "**")
	case Positive:
		writeGroup("+", v.M_Nodes, "+")
	case Negative:
		writeGroup("-", v.M_Nodes, "-")
	case Strong:
		writeGroup("*", v.M_Nodes, "*")
	case Snippet:
		writeGroup("`", v.M_Nodes, "`")
	}
}
