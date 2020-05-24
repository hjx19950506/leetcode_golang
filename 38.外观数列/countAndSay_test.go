package countAndSay

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	cases := []struct{
		n 		int
		want 	string
	} {
		{
			1,
			"1",
		},
		{
			2,
			"11",
		},
		{
			3,
			"21",
		},
		{
			4,
			"1211",
		},
		{
			5,
			"111221",
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := countAndSay(v.n)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
