package mergeKLists

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func TestMergeKLists(t *testing.T) {
	cases := []struct {
		lists 	[]*ListNode
		want 	*ListNode
	} {
		{
			[]*ListNode{
				makeList(1,4,5),
				makeList(1,3,4),
				makeList(2,6),
			},
			makeList(1,1,2,3,4,4,5,6),
		},
		{
			[]*ListNode{
				makeList(3),
				makeList(6,9,11),
				makeList(1),
				makeList(55,110),
			},
			makeList(1,3,6,9,11,55,110),
		},
		{
			[]*ListNode{
				nil,
			},
			nil,
		},
		{
			[]*ListNode{},
			nil,
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := mergeKLists(v.lists)
			if !assertListEqual(got, v.want) {
				t.Errorf("got %#v want %#v", got, v.want)
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