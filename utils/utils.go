package utils

import (
	"github.com/PaulioRandall/daft-wullie-go/node"
)

func RemoveDuplicateLines(lines []node.Node) []node.Node {

	r := []node.Node{}
	prevEmpty := false

	for _, l := range lines {
		if prevEmpty {
			if _, ok := l.(node.Empty); ok {
				continue
			}
		}

		r = append(r, l)
		_, prevEmpty = l.(node.Empty)
	}

	return r
}

/*
func TrimSpaces(lines []node.Node) []node.Node {

	r := make([]node.Node, 0, len(lines))

		for _, l := range lines {
			if prevEmpty {
				if _, ok := l.(node.Empty); ok {
					continue
				}
			}

			r = append(r, l)
			_, prevEmpty = l.(node.Empty)
		}

	return r
}
*/
