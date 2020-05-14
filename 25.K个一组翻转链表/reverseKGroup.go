package reverseKGroup
/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/
//以间隔为k的方式遍历整个链表
//编写handleKListNode函数用于处理k个节点的一截链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	res := &ListNode{-1, head}
	head = res
	for head.Next != nil {
		kListNodeHead, kListNodeTail := handleKListNode(head, k)
		head.Next = kListNodeHead
		head = kListNodeTail
	}
	return res.Next
}

//入参是这截链表的前一个节点
//出参是逆排后这截链表的头和尾
func handleKListNode(ahead *ListNode, k int) (newHead, newTail *ListNode) {
	//遍历一次确定链表接下来的节点数量不小于k个
	//若小于则不做后续逆排工作
	testK := k
	ptr := ahead
	for ; testK > 0 && ptr.Next != nil; testK--{
		ptr = ptr.Next
	}
	if testK != 0 {
		return ahead.Next, ptr
	}
	nextKListNodeHead := ptr.Next

	tailFlag := 0
	for head, tmp := ahead.Next, ahead.Next.Next; k > 0; k-- {
		head.Next = newHead
		newHead = head
		head = tmp
		if tmp.Next != nil {
			tmp = head.Next
		}
		if tailFlag == 0 {
			newTail = newHead
			newTail.Next = nextKListNodeHead
			tailFlag = 1
		}
	}

	return
}