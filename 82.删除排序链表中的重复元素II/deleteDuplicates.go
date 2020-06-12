package deleteDuplicates


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

//一次遍历
//一快一慢双指针
//快的去遍历相同数字的一段数字，找到这段节点的末尾节点
//慢的在这段节点之前，其Next指向这段之后
func deleteDuplicates(head *ListNode) *ListNode {
	preNode := ListNode{0, head}
	ptr := &preNode
	res := &preNode

	for ptr.Next != nil {
		tempPtr := ptr.Next
		tempVar := tempPtr.Val
		tempCount := 1
		for tempPtr.Next != nil && tempPtr.Next.Val == tempVar {
			tempPtr = tempPtr.Next
			tempCount++
		}
		if tempCount > 1 {
			ptr.Next = tempPtr.Next
		} else {
			ptr = ptr.Next
		}
	}

	return res.Next
}