package removeDuplicates

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct{
		nums	[]int
		want 	int
	} {
		{
			[]int{0,0,1,1,1,2,2,3,3,4},
			5,
		},
		{
			[]int{1,1,1,1,1},
			1,
		},
		{
			[]int{1,1,2},
			2,
		},
		{
			[]int{1},
			1,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := removeDuplicates(v.nums)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
