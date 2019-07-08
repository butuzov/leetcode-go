package main

import (
	"fmt"
	"strings"
)

func main() {

	t := newTree(strings.Split("25|12l|10l|4l|1l|p|p|p|5r|p|p|4r|0r|31r|100r", "|"))
	fmt.Println(t)
	fmt.Println(t.PostOrder())
	fmt.Println(postorderTraversalRecursive(t))
	fmt.Println(postorderTraversal(t))

}
