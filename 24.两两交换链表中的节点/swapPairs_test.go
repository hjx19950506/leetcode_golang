package swapPairs

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func TestSwapPairs(t *testing.T) {
	cases := []struct{
		head 	*ListNode
		want 	*ListNode
	} {
		{
			makeList(1,2,3,4,9,10),
			makeList(2,1,4,3,10,9),
		},
		{
			makeList(1,2,3,4),
			makeList(2,1,4,3),
		},
		{
			makeList(1,2,3),
			makeList(2,1,3),
		},
		{
			makeList(1),
			makeList(1),
		},
		{
			nil,
			nil,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := swapPairs(v.head)
			if !assertListEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}


func makeList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var head, ptr *ListNode
	ptr = &ListNode{nums[0], nil}
	head = ptr
	for i := 1; i < len(nums); i++ {
		ptr.Next = &ListNode{nums[i], nil}
		ptr = ptr.Next
	}
	return head
}

func assertListEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val == l2.Val {
			l1 = l1.Next
			l2 = l2.Next
		} else {
			return false
		}
	}
	return true
}