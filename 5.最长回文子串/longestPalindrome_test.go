package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct{
		name	string
		s		string
		want	string
	} {
		{
			"test1",
			"babad",
			"bab",
		},
		{
			"test2",
			"cbbd",
			"bb",
		},
		{
			"test3",
			"",
			"",
		},
		{
			"test4",
			"a",
			"a",
		},
		{
			"test5",
			"zxcvdabcbadqweq",
			"dabcbad",
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := longestPalindrome(v.s)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
