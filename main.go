package main

import (
    "algos/two_pointers"
    "fmt"
)

func main() {
    /*root := depth_first.TreeNode{Val: 1}
    two := depth_first.TreeNode{ Val: 2 }
    three := depth_first.TreeNode{ Val : 3}
    four := depth_first.TreeNode{ Val : 4 }
    five := depth_first.TreeNode{ Val: 5 }
    six := depth_first.TreeNode{ Val : 6 }
    root.Left = &two
    root.Right = &five
    two.Left = &three
    two.Right = &four
    five.Right = &six
    depth_first.Flatten(&root)*/

    result := two_pointers.MinMovesToMakePalindrome("xyzyx")
    fmt.Println(result)
}
