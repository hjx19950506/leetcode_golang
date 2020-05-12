package mergeTwoLists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head

	for l1 != nil || l2 != nil {
		p.Next = new(ListNode)
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			p.Next.Val = l1.Val
			l1 = l1.Next
		} else {
			p.Next.Val = l2.Val
			l2 = l2.Next
		}
		p = p.Next
	}

	return head.Next
}