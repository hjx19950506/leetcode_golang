package isValid

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct{
		s		string
		want	bool
	} {
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"(]",
			false,
		},
		{
			"([)]",
			false,
		},
		{
			"{[]}",
			true,
		},
		{
			"",
			true,
		},
	}
	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := isValid(v.s)
			if got != v.want {
				t.Errorf("got %#v want %#v", got, v.want)
			}
		})
	}
}
