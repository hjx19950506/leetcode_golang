package jump

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	cases := []struct{
		nums	[]int
		want 	int
	} {
		{
			[]int{2,3,1,1,4},
			2,
		},
		{
			[]int{2,1},
			1,
		},
		{
			[]int{3,2,1},
			1,
		},
		{
			[]int{1,2,3},
			2,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T){
			got := jump(v.nums)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}