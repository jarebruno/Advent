package main

import (
	"fmt"
	"strings"
)

var (
	empty = "_"
)

func addTreeBase(tree *[]string) {
	baseBorder := strings.Repeat(empty, len(*tree)-2)
	*tree = append(*tree, fmt.Sprintf("%s#%s", baseBorder, baseBorder))
	*tree = append(*tree, fmt.Sprintf("%s#%s", baseBorder, baseBorder))
}

func CreateXmasTree(height int, ornament string) string {
	tree := []string{"/*"}

	for i := 1; i <= height; i++ {
		border := strings.Repeat(empty, height-i)
		treeLevel := strings.Repeat(ornament, i*2-1)
		tree = append(tree, fmt.Sprintf("%s%s%s", border, treeLevel, border))
	}

	addTreeBase(&tree)

	tree = append(tree, "*/")
	return strings.Join(tree, "\n")
}
