package main

import (
    "algos/recursion"
    "fmt"
)

func main() {
    first := recursion.ListNode{ Val : 5}
    second := recursion.ListNode{ Val : 2}
    third := recursion.ListNode{ Val : 13}
    fourth := recursion.ListNode{ Val : 3}
    fifth := recursion.ListNode{ Val : 8}
    first.Next = &second
    second.Next = &third
    third.Next = &fourth
    fourth.Next = &fifth
    result := recursion.RemoveNodesNonRec(&first)
    fmt.Println(result)
}
