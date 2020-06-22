package reverseBetween


type ListNode struct {
	Val int
	Next *ListNode
}

//分别两个指针指向第m和n个节点
//然后依次将原链表的m+1至n的节点加到m-1节点的之后
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	pre := &ListNode{Val:  -1, Next: head}
	mPrePtr := pre

	for i := m - 1; i != 0; i--{
		mPrePtr = mPrePtr.Next
	}
	mPtr := mPrePtr.Next
	for i := n - m; i > 0; i-- {
		moveNode(mPrePtr, mPtr)
	}
	return pre.Next
}

func moveNode(dst *ListNode, src *ListNode) {
	temp := src.Next
	src.Next = src.Next.Next
	temp.Next = dst.Next
	dst.Next = temp
}