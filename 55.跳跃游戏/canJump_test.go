package canJump

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	cases := []struct{
		nums	[]int
		want	bool
	} {
		{
			[]int{2,3,1,1,4},
			true,
		},
		{
			[]int{3,2,1,0,4},
			false,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := canJump(v.nums)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}