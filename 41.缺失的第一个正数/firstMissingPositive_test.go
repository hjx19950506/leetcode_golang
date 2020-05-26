package firstMissingPositive

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct{
		nums 	[]int
		want 	int
	} {
		{
			[]int{1,2,0},
			3,
		},
		{
			[]int{3,4,-1,1},
			2,
		},
		{
			[]int{7,8,9,11,12},
			1,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T){
			got := firstMissingPositive(v.nums)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}