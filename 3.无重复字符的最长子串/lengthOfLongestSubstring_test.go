package lengthOfLongestSubstring

import "testing"

//func lengthOfLongestSubstring(s string) int
func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct{
		name	string
		s		string
		want	int
	} {
		{
			"test1",
			"abcabcbb",
			3,
		},
		{
			"test2",
			"a",
			1,
		},
		{
			"test3",
			"",
			0,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(v.s)

			if got != v.want {
				t.Errorf("got %d want %d", got, v.want)
			}
		})
	}
}
