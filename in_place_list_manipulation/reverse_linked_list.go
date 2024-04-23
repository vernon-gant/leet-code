package in_place_list_manipulation

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func reverse(head *EduLinkedListNode) *EduLinkedListNode {
	return reverseHelper(nil, head);
}

func reverseHelper(reversed, currentHead *EduLinkedListNode) *EduLinkedListNode {
	if currentHead == nil {
		return reversed
	}
	tail := currentHead.next
	currentHead.next = reversed
	return reverseHelper(currentHead, tail)
}