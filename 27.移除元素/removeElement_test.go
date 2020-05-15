package removeElement

import (
	"fmt"
	"testing"
)

func TestRemoveEement(t *testing.T) {
	cases := []struct{
		nums	[]int
		val 	int
		want 	int
	} {
		{
			[]int{},
			3,
			0,
		},
		{
			[]int{3,2,2,3},
			3,
			2,
		},
		{
			[]int{2,2,2,3},
			2,
			1,
		},
		{
			[]int{0,1,2,2,3,0,4,2},
			2,
			5,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := removeElement(v.nums, v.val)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}