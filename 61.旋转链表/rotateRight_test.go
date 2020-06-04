package rotateRight

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T){
	cases := []struct{
		head 	*ListNode
		k		int
		want	*ListNode
	} {
		{
			makeList(1),
			4,
			makeList(1),
		},
		{
			makeList(1,2,3,4,5),
			2,
			makeList(4,5,1,2,3),
		},
		{
			makeList(0,1,2),
			4,
			makeList(2,0,1),
		},
		{
			makeList(0),
			4,
			makeList(0),
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := rotateRight(v.head, v.k)
			if !assertListEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}

type ListNode struct {
	Val int
	Next *ListNode
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
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}