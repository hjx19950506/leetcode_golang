package climbStairs

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T)  {
	cases := []struct{
		n		int
		want	int
	} {
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			1,
			1,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := climbStairs(v.n)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}