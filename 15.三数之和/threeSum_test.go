package threeSum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct{
		nums	[]int
		want 	[][]int
	} {
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1,-1,2}, {-1,0,1}},
		},
		{
			[]int{3,-1,-2},
			[][]int{{-2,-1,3}},
		},
		{
			[]int{-2,0,0,2,2},
			[][]int{{-2,0,2}},
		},
		{
			[]int{0,0,0},
			[][]int{{0,0,0}},
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := threeSum(v.nums)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
