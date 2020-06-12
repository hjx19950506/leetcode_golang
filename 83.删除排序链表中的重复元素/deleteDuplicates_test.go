package deleteDuplicates

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	cases := []struct{
		head		*ListNode
		want		*ListNode
	} {
		{
			makeList(1,2,3,3,4,4,5),
			makeList(1,2,3,4,5),
		},
		{
			makeList(1,1,1,2,3),
			makeList(1,2,3),
		},
		{
			makeList(1,1,1,2,3,3,3,4,5,5),
			makeList(1,2,3,4,5),
		},
		{
			makeList(1),
			makeList(1),
		},
		{
			makeList(),
			makeList(),
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := deleteDuplicates(v.head)
			if !assertListEqual(got, v.want) {
/*				fmt.Printf("got:")
				printList(got)
				fmt.Printf("want:")
				printList(v.want)*/
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