package partition

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct{
		head	*ListNode
		x		int
		want	*ListNode
	} {
		{
			makeList(2,0,4,1,3,1,4,0,3),
			4,
			makeList(2,0,1,3,1,0,3,4,4),
		},
		{
			makeList(1,3,-1,5,2,1),
			3,
			makeList(1,-1,2,1,3,5),
		},
		{
			makeList(1,2,3),
			4,
			makeList(1,2,3),
		},
		{
			makeList(1,4,3,2,5,2),
			3,
			makeList(1,2,2,4,3,5),
		},
		{
			makeList(4,3,2,5,2),
			3,
			makeList(2,2,4,3,5),
		},
		{
			makeList(4,3,5),
			3,
			makeList(4,3,5),
		},
		{
			makeList(1,1,1),
			3,
			makeList(1,1,1),
		},
		{
			makeList(1),
			3,
			makeList(1),
		},
		{
			makeList(5),
			3,
			makeList(5),
		},
		{
			makeList(),
			3,
			makeList(),
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := partition(v.head, v.x)
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
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}

func (head *ListNode) String() string {
	res := ""
	if head == nil {
		return res
	}
	res += fmt.Sprintf("%v", head.Val)
	head = head.Next
	for head != nil {
		res += fmt.Sprintf(" ==> %v", head.Val)
		head = head.Next
	}
	return res
}