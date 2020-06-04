package rotateRight
/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/

//先遍历一次链表得到长度，记录旧尾
//k对链表长度取模得到实际的k
//再遍历一次找到新的头尾即可
func rotateRight(head *ListNode, k int) *ListNode {
	listLen := 0
	var oldTail *ListNode = nil
	var newHead, newTail *ListNode = nil, head
	for ptr := head; ptr != nil; ptr = ptr.Next {
		listLen++
		if ptr.Next == nil {
			oldTail = ptr
		}
	}
	if listLen == 0 {
		return nil
	}
	k = k % listLen
	if k == 0 {
		return head
	}

	for moveStep := listLen - k - 1; moveStep > 0; moveStep-- {
		newTail = newTail.Next
	}
	newHead = newTail.Next
	newTail.Next = nil
	oldTail.Next = head
	return newHead
}