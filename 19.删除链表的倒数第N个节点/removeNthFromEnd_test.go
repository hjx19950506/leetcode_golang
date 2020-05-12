package removeNthFromEnd

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct{
		head 	*ListNode
		n 		int
		want 	*ListNode
	} {
		{
			makeList([]int{1,2,3,4,5}),
			2,
			makeList([]int{1,2,3,5}),
		},
		{
			makeList([]int{1,2}),
			2,
			makeList([]int{2}),
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := removeNthFromEnd(v.head, v.n)
			if !assertListEqual(got, v.want) {
				t.Errorf("got %#v want %#v", got, v.want)
			}
		})
	}
}

func makeList(nums []int) *ListNode {
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