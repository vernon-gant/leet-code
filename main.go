package main

import (
    "algos/sliding_window"
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

    result := sliding_window.MinSubArrayLen(7, []int{2,3,1,2,4,3})
    fmt.Println(result)
}
