package _24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	link := head
	var parent *ListNode

	for link != nil && link.Next != nil {
		tmp := link.Next
		link.Next = tmp.Next
		tmp.Next = link

		if parent != nil {
			parent.Next = tmp
		}

		parent = link
		link = link.Next
	}

	return newHead
}
