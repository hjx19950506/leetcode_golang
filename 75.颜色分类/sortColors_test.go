package sortColors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		nums	[]int
		want	[]int
	} {
		{
			[]int{2,0,2,1,1,0},
			[]int{0,0,1,1,2,2},
		},
		{
			[]int{2,0},
			[]int{0,2},
		},
		{
			[]int{2,0,1},
			[]int{0,1,2},
		},
		{
			[]int{1,2,0},
			[]int{0,1,2},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			sortColors(v.nums)
			if !reflect.DeepEqual(v.nums, v.want) {
				t.Errorf("got %v want %v", v.nums , v.want)
			}
		})
	}
}