package isPalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct{
		name	string
		x		int
		want	bool
	} {
		{
			"test1",
			121,
			true,
		},
		{
			"test2",
			-121,
			false,
		},
		{
			"test3",
			10,
			false,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := isPalindrome(v.x)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
