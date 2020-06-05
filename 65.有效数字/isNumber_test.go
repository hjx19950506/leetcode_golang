package isNumber

import (
	"fmt"
	"testing"
)

func TestIsNumber(t *testing.T) {
	cases := []struct {
		s 		string
		want	bool
	} {
		{"0",true},
		{" 0.1 ",true},
		{"abc",false},
		{"1 a",false},
		{"2e10",true},
		{" -90e3   ",true},
		{" 1e",false},
		{"e3",false},
		{" 6e-1",true},
		{" 99e2.5 ",false},
		{"53.5e93",true},
		{" --6 ",false},
		{"-+3",false},
		{"95a54e53",false},
		{".1",true},
		{"-.1",true},
		{"1.",true},
		{"005047e+6",true},
		{"46.e3",true},
		{".2e81",true},
		{".",false},
		{"+e",false},

	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := isNumber(v.s)
			if got != v.want {
				t.Errorf("got %v want %v", got ,v.want)
			}
		})
	}
}