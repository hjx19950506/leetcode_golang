package partition

type ListNode struct {
	Val int
	Next *ListNode
}

//首先确保链表头为小于x的节点
//然后遍历一次链表，将小于x的节点前移即可
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	//确保链表头为小于x的节点
	if head.Val >= x {
		ptr := head
		for ptr.Next != nil && ptr.Next.Val >= x {
			ptr = ptr.Next
		}
		//全链表都大于等于x则无需改变
		if ptr.Next == nil {
			return head
		}
		moveNode(preHead, ptr)
	}
	lastLittleNode, ptr := preHead.Next, preHead.Next
	//找到第一个大于等于x节点的前一个节点，排除全链表都小于x的特殊情况
	for ptr.Next != nil && ptr.Next.Val < x {
		ptr = ptr.Next
		lastLittleNode = lastLittleNode.Next
	}
	if ptr.Next == nil {
		return head
	}
	//遍历一次链表，将小于x的节点前移即可
	for ptr != nil && ptr.Next != nil {
		if ptr.Next.Val < x {
			moveNode(lastLittleNode, ptr)
			lastLittleNode = lastLittleNode.Next
		} else {
			ptr = ptr.Next
		}
	}
	return preHead.Next
}

func moveNode(dst *ListNode, src *ListNode) {
	temp := src.Next
	src.Next = src.Next.Next
	temp.Next = dst.Next
	dst.Next = temp
}