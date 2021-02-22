package ast

import (
	"strings"
)

type (
	Notes       []Node
	DescendFunc func(n Node, lineNum, depth, orderIdx int)
)

func DescendNotes(n Notes, f DescendFunc) {
	descendNodes(n, 1, 0, f)
}

func DecendNode(n Node, f DescendFunc) {
	descendNode(n, 1, 0, 0, f)
}

func descendNodes(ns []Node, lineNum, depth int, f DescendFunc) {
	for i, n := range ns {
		descendNode(n, lineNum, depth, i, f)
	}
}

func descendNode(n Node, lineNum, depth, orderIdx int, f DescendFunc) {
	f(n, lineNum, depth, orderIdx)
	if v, ok := n.(Parent); ok {
		descendNodes(v.Nodes(), lineNum, depth+1, f)
	}
}

func RemoveExtraLines(notes Notes) Notes {

	r := []Node{}
	prevEmpty := false

	for _, n := range notes {
		currEmpty := n.Type() == EmptyLine

		if !prevEmpty || !currEmpty {
			r = append(r, n)
			prevEmpty = currEmpty
		}
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

	writeGroup := func(prefix string, n Node, suffix string) {
		sb.WriteString(prefix)

		if p, ok := n.(ParentNode); ok {
			for _, sub := range p.Nodes() {
				fmtNodeString(sb, sub)
			}
		} else {
			s := n.Text()
			sb.WriteString(s)
		}

		sb.WriteString(suffix)
	}

	switch n.Type() {
	case Text:
		writeGroup("", n, "")

	case Topic:
		writeGroup("#", n, "")
	case SubTopic:
		writeGroup("##", n, "")

	case BulPoint:
		writeGroup(".", n, "")
	case SubBulPoint:
		writeGroup("..", n, "")
	case NumPoint:
		writeGroup("!", n, "")
	case SubNumPoint:
		writeGroup("!!", n, "")

	case TextLine, EmptyLine:
		writeGroup("", n, "")

	case KeyPhrase:
		writeGroup("**", n, "**")
	case Positive:
		writeGroup("+", n, "+")
	case Negative:
		writeGroup("-", n, "-")
	case Strong:
		writeGroup("*", n, "*")
	case Quote:
		writeGroup(`"`, n, `"`)
	case Snippet:
		writeGroup("`", n, "`")
	}
}
