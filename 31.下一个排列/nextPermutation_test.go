package nextPermutation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	cases := []struct{
		nums 	[]int
		want 	[]int
	} {
		{
			[]int{1,3,2},
			[]int{2,1,3},
		},
		{
			[]int{1,2,3},
			[]int{1,3,2},
		},
		{
			[]int{3,2,1},
			[]int{1,2,3},
		},
		{
			[]int{1,1,5},
			[]int{1,5,1},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			nextPermutation(v.nums)
			if !reflect.DeepEqual(v.nums, v.want) {
				t.Errorf("got %v want %v", v.nums, v.want)
			}
		})
	}
}
