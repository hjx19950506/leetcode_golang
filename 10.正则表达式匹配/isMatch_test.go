package isMatch

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	cases := []struct{
		s 		string
		p 		string
		want	bool
	} {
		{
			"",
			"",
			true,
		},
		{
			"",
			".",
			false,
		},
		{
			"a",
			"",
			false,
		},
		{
			"aa",
			"a",
			false,
		},
		{
			"aa",
			"a*",
			true,
		},
		{
			"ab",
			".*",
			true,
		},
		{
			"aab",
			"c*a*b",
			true,
		},
		{
			"mississippi",
			"mis*is*p*.",
			false,
		},
	}

	for _, v := range cases {
		t.Run(fmt.Sprintf("s:%v,p:%v", v.s, v.p), func(t *testing.T) {
			got := isMatch(v.s, v.p)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
