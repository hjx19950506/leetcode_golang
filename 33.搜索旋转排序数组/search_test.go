package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
		0,
		4,
		},
		{
			[]int{4,5,6,7,0,1,2},
			0,
			4,
		},
		{
			[]int{4,5,6,7,0,1,2},
			4,
			0,
		},
		{
			[]int{4,5,6,7,0,1,2},
			2,
			6,
		},
		{
			[]int{},
			3,
			-1,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := search(v.nums, v.target)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
