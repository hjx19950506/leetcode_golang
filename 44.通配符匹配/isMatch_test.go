package isMatch

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	cases := []struct{
		s		string
		p		string
		want	bool
	} {
		{
			"aa",
			"a",
			false,
		},
		{
			"aa",
			"*",
			true,
		},
		{
			"cb",
			"?a",
			false,
		},
		{
			"adceb",
			"*a*b",
			true,
		},
		{
			"acdcb",
			"a*c?b",
			false,
		},
		{
			"",
			"*",
			true,
		},
		{
			"asfd",
			"",
			false,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T){
			got := isMatch(v.s, v.p)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}