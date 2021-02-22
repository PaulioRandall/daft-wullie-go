package main

import (
	"fmt"

	"github.com/PaulioRandall/daft-wullie-go/ast"
	"github.com/PaulioRandall/daft-wullie-go/parser"
	"github.com/PaulioRandall/daft-wullie-go/scanner"
)

func main() {

	const example = `
# Topic
## Sub Topic

. Bullet point
.. Sub bullet point
! Numbered point
!! Sub numbered point
"Quote with a +positive+ point"

There is only text here as the control symbol '\*' has been escaped

A sentence with a **keyword** in it
+A positive sentence
-A negative sentence
*Some strong words
"Quote
` + "`Snippet or literal sentence"

	tks := scanner.ScanAll(example)
	notes := parser.ParseAll(tks)

	hCount, pCount, nCount := 0, 0, 0
	f := func(n ast.Node, lineNum, depth, orderIdx int) {
		switch n.Type() {
		case ast.Topic, ast.SubTopic:
			hCount++
		case ast.Positive:
			pCount++
		case ast.Negative:
			nCount++
		}
	}

	ast.DescendNotes(notes, f)

	fmt.Println()
	fmt.Print("```")
	fmt.Println(example)
	fmt.Println("```")
	fmt.Println()
	fmt.Println("The text above contains:")
	fmt.Println(hCount, "Topic & sub topic lines")
	fmt.Println(pCount, "positive points")
	fmt.Println(nCount, "negative points")
}
