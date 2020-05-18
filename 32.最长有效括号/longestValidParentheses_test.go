package longestValidParentheses

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	cases := []struct{
		s 		string
		want 	int
	} {
		{
			"(()",
			2,
		},
		{
			")()())",
			4,
		},
		{
			"",
			0,
		},
		{
			")",
			0,
		},
		{
			"(()()(()())())(",
			14,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := longestValidParentheses(v.s)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}