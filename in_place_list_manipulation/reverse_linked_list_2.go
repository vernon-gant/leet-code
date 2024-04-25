package in_place_list_manipulation

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	return reverseBetweenHelper(dummy, dummy, head, left, 1, right-left+1)
}

func reverseBetweenHelper(result, beforeReversed, head *ListNode, left, listPointer, reversedListLength int) *ListNode {
	if listPointer < left {
		return reverseBetweenHelper(result, head, head.Next, left, listPointer+1, reversedListLength)
	}
	reversed, listRest := reverseBetweenReversing(nil, head, 0, reversedListLength)
	beforeReversed.Next = reversed
	head.Next = listRest
	return result.Next
}

func reverseBetweenReversing(reversed, head *ListNode, currentCounter, reversedListLength int) (*ListNode, *ListNode) {
	if currentCounter == reversedListLength {
		return reversed, head
	}
	tail := head.Next
	head.Next = reversed
	return reverseBetweenReversing(head, tail, currentCounter+1, reversedListLength)
}