package romanToInt

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct{
		s 		string
		want 	int
	} {
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		},
		{
			"LVIII",
			58,
		},
		{
			"MCMXCIV",
			1994,
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := romanToInt(v.s)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}