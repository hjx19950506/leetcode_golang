package addTwoNumbers

import (
	"fmt"
)


 type ListNode struct {
      Val int
      Next *ListNode
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    preHead := &ListNode{
		0, nil,
	}
	ptr := preHead
	var tmp, carry, remainder int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmp = l2.Val + carry
		} else if l2 == nil {
			tmp = l1.Val + carry
		} else {
			tmp = l1.Val + l2.Val + carry
		}
		carry = tmp / 10
		remainder = tmp % 10
		ptr.Next = &ListNode{
			remainder, nil,
		}
		ptr = ptr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		ptr.Next = &ListNode{
			carry, nil,
		}
		ptr = ptr.Next
	}

	return preHead.Next
}

func main() {
	ln1 := ListNode{
		1, nil,
	}
	ln2 := ListNode{
		2, &ln1,
	}
	ln3 := ListNode{
		3, &ln2,
	}

	ln4 := ListNode{
		4, nil,
	}
	ln5 := ListNode{
		5, &ln4,
	}
	ln6 := ListNode{
		6, &ln5,
	}
	ln7 := ListNode{
		7, &ln6,
	}

	ln := addTwoNumbers(&ln3, &ln7)
	for ln != nil {
		fmt.Println(ln.Val)
		ln = ln.Next
	}
}