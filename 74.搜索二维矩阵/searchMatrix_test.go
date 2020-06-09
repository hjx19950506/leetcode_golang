package searchMatrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	cases := []struct {
		matrix	[][]int
		target	int
		want	bool
	} {
		{
			[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,50}},
			3,
			true,
		},
		{
			[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,50}},
			13,
			false,
		},
		{
			[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,50}},
			7,
			true,
		},
		{
			[][]int{{1,3,5,7}},
			13,
			false,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := searchMatrix(v.matrix, v.target)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}