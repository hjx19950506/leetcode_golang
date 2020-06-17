package isScramble

import (
	"fmt"
	"testing"
)

func TestIsScramble(t *testing.T) {
	cases := []struct{
		s1		string
		s2		string
		want	bool
	} {
		{
			"great",
			"rgeat",
			true,
		},
		{
			"abcde",
			"caebd",
			false,
		},
		{
			"abb",
			"bba",
			true,
		},
		{
			"ac",
			"ca",
			true,
		},
		{
			"",
			"",
			true,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := isScramble(v.s1, v.s2)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}