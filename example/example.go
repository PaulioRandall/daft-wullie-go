package main

import (
	"fmt"

	"github.com/PaulioRandall/daft-wullie-go/ast2"
	"github.com/PaulioRandall/daft-wullie-go/parser"
	"github.com/PaulioRandall/daft-wullie-go/scanner"
)

func main() {

	const example = `
# Heading 1
## Heading 2
### Heading 3

. Bullet point
! Numbered point
> Quote with a +positive+ point

There is only text here as the control symbol '\*' has been escaped

A sentence with a **keyword** in it
+A positive sentence
-A negative sentence
*Some strong words
` + "`Snippet or literal sentence"

	tks := scanner.ScanAll(example)
	notes := parser.ParseAll(tks)

	hCount, pCount, nCount := 0, 0, 0
	f := func(n ast2.Node, lineNum, depth, orderIdx int) {
		switch n.Type() {
		case ast2.H1:
			hCount++
		case ast2.H2:
			hCount++
		case ast2.H3:
			hCount++
		case ast2.Positive:
			pCount++
		case ast2.Negative:
			nCount++
		}
	}

	ast2.DescendNotes(notes, f)

	fmt.Println()
	fmt.Print("```")
	fmt.Println(example)
	fmt.Println("```")
	fmt.Println()
	fmt.Println("The text above contains:")
	fmt.Println(hCount, "header lines (any kind)")
	fmt.Println(pCount, "positive points")
	fmt.Println(nCount, "negative points")
}
