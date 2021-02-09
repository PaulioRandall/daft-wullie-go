package node

type Notes []Node

func RemoveDuplicateLines(n Notes) Notes {

	r := []Node{}
	prevEmpty := false

	for _, l := range n {
		if prevEmpty {
			if _, ok := l.(Empty); ok {
				continue
			}
		}

		r = append(r, l)
		_, prevEmpty = l.(Empty)
	}

	return Notes(r)
}
