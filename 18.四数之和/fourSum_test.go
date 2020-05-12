package fourSum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	cases := []struct{
		nums	[]int
		target 	int
		want 	[][]int
	} {
		{
			[]int{1,0,-1,0,-2,2},
			0,
			[][]int{{-1,0,0,1}, {-2,-1,1,2}, {-2,0,0,2}},
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := fourSum(v.nums, v.target)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %#v want %#v", got, v.want)
			}
		})
	}
}
