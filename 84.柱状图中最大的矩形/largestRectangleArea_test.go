package largestRectangleArea

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	cases := []struct{
		heights		[]int
		want		int
	} {
		{
			[]int{2,1,5,6,2,3},
			10,
		},
		{
			[]int{2,1,5,6,5,2,3},
			15,
		},
		{
			[]int{2,1,5,5,3,5,2,3},
			12,
		},
		{
			[]int{2,1,5,5,2,3},
			10,
		},
		{
			[]int{},
			0,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := largestRectangleArea(v.heights)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}