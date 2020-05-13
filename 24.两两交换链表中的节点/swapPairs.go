package swapPairs
/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/

//快慢指针
//但注意 两个指针并非指向需要交换的两个节点
//而是这两个节点各自的前一个节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	res := &ListNode{-1, head}
	fastPtr, slowPtr := res, head
	for slowPtr.Next != nil {
		fastPtr.Next, slowPtr.Next = slowPtr.Next, slowPtr.Next.Next
		fastPtr.Next.Next = slowPtr
		fastPtr = slowPtr
		if slowPtr.Next != nil {
			slowPtr = slowPtr.Next
		} else {
			break
		}
	}

	return res.Next
}