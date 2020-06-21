package numDecodings

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	cases := []struct{
		s		string
		want	int
	} {
		{
			"12",
			2,
		},
		{
			"226",
			3,
		},
		{
			"2226",
			5,
		},
		{
			"2",
			1,
		},
		{
			"10",
			1,
		},
		{
			"202",
			1,
		},
		{
			"100",
			0,
		},
		{
			"110",
			1,
		},
		{
			"301",
			0,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := numDecodings(v.s)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}