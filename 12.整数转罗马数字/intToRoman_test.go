package intToRoman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	cases := []struct{
		num		int
		want	string
	} {
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			9,
			"IX",
		},
		{
			58,
			"LVIII",
		},
		{
			1994,
			"MCMXCIV",
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := intToRoman(v.num)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
