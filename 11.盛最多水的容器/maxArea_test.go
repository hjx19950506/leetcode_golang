package maxArea

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	cases := []struct{
		height	[]int
		want	int
	} {
		{
			[]int{1,2,1},
			2,
		},
		{
			[]int{1,1},
			1,
		},

		{
		[]int{1,8,6,2,5,4,8,3,7},
			49,
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := maxArea(v.height)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
