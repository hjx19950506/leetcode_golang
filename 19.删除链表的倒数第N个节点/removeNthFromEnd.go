package removeNthFromEnd
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//快慢指针
//在链表前再加一个节点
//这样要删除头节点时也可以处理
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head}
	fastPtr, slowPtr := newHead, newHead
	for n > 0 {
		fastPtr = fastPtr.Next
		n--
	}

	for fastPtr.Next != nil {
		fastPtr = fastPtr.Next
		slowPtr = slowPtr.Next
	}
	slowPtr.Next = slowPtr.Next.Next

	return newHead.Next
}