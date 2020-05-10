package longestCommonPrefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct{
		strs	[]string
		want	string
	} {
		{
			[]string{"flower","flow","flight"},
			"fl",
		},
		{
			[]string{"dog","racecar","car"},
			"",
		},
		{
			[]string{""},
			"",
		},
		{
			[]string{"a"},
			"a",
		},
		{
			[]string{"c", "c"},
			"c",
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := longestCommonPrefix(v.strs)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
