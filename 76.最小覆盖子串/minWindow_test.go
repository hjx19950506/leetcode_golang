package minWindow

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	cases := []struct{
		s		string
		t		string
		want	string
	} {
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"ADOBECODEBANC",
			"CBBEEDO",
			"BECODEB",
		},
		{
			"ADOBECODEBANC",
			"DAZ",
			"",
		},
		{
			"",
			"DAZ",
			"",
		},
		{
			"ADOBECODEBANC",
			"",
			"",
		},
		{
			"ADOBECODEBANC",
			"ADO",
			"ADO",
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := minWindow(v.s, v.t)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}