package main

import (
	"algos/trees"
	"fmt"
)

func main() {
	root := trees.TreeNode{Val: 3}
	nineNode := trees.TreeNode{Val: 9}
	twentyNode := trees.TreeNode{Val: 20}
	fifteenNode := trees.TreeNode{Val: 15}
	sevenNode := trees.TreeNode{Val: 7}
	root.Left = &nineNode
	root.Right = &twentyNode
	twentyNode.Left = &fifteenNode
	twentyNode.Right = &sevenNode
	fmt.Println(trees.IsBalanced(&root))
}
