package _1721

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	begin, end := func(head *ListNode, k int) (begin *ListNode, end *ListNode) {
		i := 0
		for curr := head; curr != nil; curr = curr.Next {
			i++
			if i == k {
				begin = curr
			}
			if i == k+1 {
				end = head
			}
			if end != nil {
				end = end.Next
			}
		}
		if i == k {
			return head, begin
		}
		return begin, end
	}(head, k)
	begin.Val, end.Val = end.Val, begin.Val
	return head
}
