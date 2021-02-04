package reader

type runeReader struct {
	runes []rune
}

func NewReader(s string) *runeReader {
	return &runeReader{runes: []rune(s)}
}

func (rr *runeReader) More() bool {
	return len(rr.runes) > 0
}

func (rr *runeReader) Match(start int, s string) bool {

	needle := []rune(s)
	if len(rr.runes)+start < len(needle) {
		return false
	}

	for i, ru := range needle {
		if rr.runes[i+start] != ru {
			return false
		}
	}

	return true
}

func (rr *runeReader) Accept(s string) bool {
	if rr.Match(0, s) {
		rr.ReadMany(len([]rune(s)))
		return true
	}
	return false
}

func (rr *runeReader) MatchNewline() bool {
	return rr.Match(0, "\n") || rr.Match(0, "\r\n")
}

func (rr *runeReader) AcceptNewline() bool {
	return rr.Accept("\n") || rr.Accept("\r\n")
}

func (rr *runeReader) Drain() []rune {
	return rr.ReadMany(len(rr.runes))
}

func (rr *runeReader) Read() rune {

	if len(rr.runes) == 0 {
		panic("Index out of range, reading too many runes")
	}

	r := rr.runes[0]
	rr.runes = rr.runes[1:]
	return r
}

func (rr *runeReader) ReadMany(n int) []rune {

	if len(rr.runes) < n {
		panic("Index out of range, reading too many runes")
	}

	if n == 0 {
		return []rune{}
	}

	r := rr.runes[:n]
	rr.runes = rr.runes[n:]
	return r
}

func (rr *runeReader) ReadLine() []rune {
	for i := 0; i < len(rr.runes); i++ {
		if rr.Match(i, "\n") || rr.Match(i, "\r\n") {
			return rr.ReadMany(i)
		}
	}
	return rr.Drain()
}
