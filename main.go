package main

import "algos/binary_tree"

func main() {
    root := &binary_tree.TreeNode{
        Val : 1,
        Left: &binary_tree.TreeNode{
            Val : 2,
            Left : &binary_tree.TreeNode {
                Val: 4,
                Left: &binary_tree.TreeNode{
                    Val: 7,
                },
            },
            Right: &binary_tree.TreeNode{
                Val: 5,
                Right: &binary_tree.TreeNode{
                    Val: 8,
                    Left: &binary_tree.TreeNode{
                        Val: 10,
                    },
                    Right: &binary_tree.TreeNode{
                        Val: 11,
                        Right: &binary_tree.TreeNode{
                            Val: 13,
                        },
                    },
                },
            },
        },
        Right: &binary_tree.TreeNode{
            Val: 3,
            Right: &binary_tree.TreeNode{
                Val: 6,
                Right: &binary_tree.TreeNode{
                    Val: 9,
                    Right: &binary_tree.TreeNode{
                        Val: 12,
                    },
                },
            },
        },
    }
    result := binary_tree.DiameterOfBinaryTree2(root)
    print(result)
}
