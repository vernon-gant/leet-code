package in_place_list_manipulation

func ReverseKGroup(head *ListNode, k int) *ListNode {
	reversed := &ListNode{}
	return reverseKGroupHelper(reversed, reversed, head, k)
}

func reverseKGroupHelper(reversed, reversedTail, head *ListNode, k int) *ListNode {
	nextGroupStart, currentGroupSize := findNextGroupStart(head, head, 1, k)
	if currentGroupSize == k {
		reversedTail.Next = reverseListHelper(nil, head)
	}
	if nextGroupStart == nil {
		return reversed.Next
	}
	head.Next = nextGroupStart
	return reverseKGroupHelper(reversed, head, nextGroupStart, k)
}

func findNextGroupStart(head, currentPointer *ListNode, currentSize, k int) (*ListNode, int) {
	if currentPointer == nil {
		return nil, -currentSize
	}
	if currentSize == k {
		nextGroupStart := currentPointer.Next
		currentPointer.Next = nil
		return nextGroupStart, k
	}
	return findNextGroupStart(head, currentPointer.Next, currentSize+1, k)
}
