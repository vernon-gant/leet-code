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

    /*first := &fast_slow.ListNode{Val: 81}
    second := &fast_slow.ListNode{Val: 144}
    third := &fast_slow.ListNode{Val: 64}
    fourth := &fast_slow.ListNode{Val: 121}
    fifth := &fast_slow.ListNode{Val: 25}
    sixth := &fast_slow.ListNode{Val: 49}

    first.Next = second
    second.Next = third
    third.Next = fourth
    fourth.Next = fifth
    fifth.Next = sixth*/

    result := sliding_window.FindMaxAverage([]int{1,12,-5,-6,50,3}, 4)
    fmt.Println(result)
}
