package main

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/daft-wullie-go/node"
	"github.com/PaulioRandall/daft-wullie-go/parser"
	"github.com/PaulioRandall/daft-wullie-go/scanner"
)

const example = `
# Heading 1
## Heading 2
### Heading 3
. Bullet point
1. Numbered point
> Quote

**Key phrase
+Positive
-Negative
*Strong
` + "`Snippet" + `

*Strong*+Positive+-Negative-
`

func main() {
	tks := scanner.ScanAll(example)
	notes := parser.ParseAll(tks)

	f := func(n node.Node, lineNum, depth, orderIdx int) {
		fmt.Print(strings.Repeat("  ", depth))
		switch n.(type) {
		case node.EmptyLine:
			// Ignore
		case node.Phrase:
			fmt.Println(`"` + strings.TrimSpace(n.Text()) + `"`)
		case node.Quote:
			fmt.Println(n.Name(), `"`+strings.TrimSpace(n.Text())+`"`)
		default:
			fmt.Println(n.Name())
		}
	}

	node.DescendNotes(notes, f)
}
