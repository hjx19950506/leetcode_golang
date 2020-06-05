package minPathSum

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	cases := []struct{
		grid	[][]int
		want	int
	} {
		{
			[][]int{{1,3,1},{1,5,1},{4,2,1}},
			7,
		},
		{
			[][]int{{1,3,1}},
			5,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := minPathSum(v.grid)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}