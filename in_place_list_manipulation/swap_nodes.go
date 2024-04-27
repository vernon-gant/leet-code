package in_place_list_manipulation

func SwapNodes(head *ListNode, k int) *ListNode {
	return swapNodesHelper(head, head, head, head, 1, k)
}

func swapNodesHelper(head, current, front, end *ListNode, counter, k int) *ListNode {
	if counter == k {
		return swapNodesHelper(head, current.Next, current, head, counter + 1, k)
	}
	if current == nil {
		frontVal := front.Val
		front.Val = end.Val
		end.Val = frontVal
		return head
	}
	return swapNodesHelper(head, current.Next, front, end.Next, counter + 1, k)
}