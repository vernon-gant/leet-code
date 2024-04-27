package main

import (
	"algos/in_place_list_manipulation"
	"fmt"
)

func main() {
	one := in_place_list_manipulation.ListNode{Val: 1}
	two := in_place_list_manipulation.ListNode{Val: 2}
	three := in_place_list_manipulation.ListNode{Val: 3}
	four := in_place_list_manipulation.ListNode{Val: 4}
	five := in_place_list_manipulation.ListNode{Val: 5}
	six := in_place_list_manipulation.ListNode{Val: 6}
	seven := in_place_list_manipulation.ListNode{Val: 7}
	eight := in_place_list_manipulation.ListNode{Val: 8}
	nine := in_place_list_manipulation.ListNode{Val: 9}
	ten := in_place_list_manipulation.ListNode{Val: 10}
	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five
	five.Next = &six
	six.Next = &seven
	seven.Next = &eight
	eight.Next = &nine
	nine.Next = &ten
	in_place_list_manipulation.SwapNodes(&one,5)
	fmt.Print(one)
}
