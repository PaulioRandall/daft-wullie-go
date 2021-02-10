package main

import (
	"fmt"

	"github.com/PaulioRandall/daft-wullie-go/node"
	"github.com/PaulioRandall/daft-wullie-go/parser"
	"github.com/PaulioRandall/daft-wullie-go/scanner"
)

const example = `
# H1
## H2
### H3
. Bullet point
1. Numbered point
> Quote

**Key phrase**
+Positive+
-Negative-
*Strong*
` + "`Snippet`"

func main() {
	tks := scanner.ScanAll(example)
	notes := parser.ParseAll(tks)
	s := node.FmtString(notes)
	fmt.Println(s)
}
