package uniquePaths

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		m    int
		n    int
		want int
	} {
		{
			3,
			2,
			3,
		},
		{
			7,
			3,
			28,
		},
		{
			1,
			2,
			1,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := uniquePaths(v.m, v.n)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}