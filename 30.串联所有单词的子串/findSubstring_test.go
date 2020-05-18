package findSubstring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	cases := []struct{
		s 		string
		words	[]string
		want 	[]int
	} {
		{
			"barfoothefoobarman",
			[]string{"foo","bar"},
			[]int{0,9},
		},
		{
			"wordgoodgoodgoodbestword",
			[]string{"word","good","best","word"},
			[]int{},
		},
		{
			"wordgoodgoodgoodbestword",
			[]string{"word","good","best","good"},
			[]int{8},
		},
		{
			"a",
			[]string{},
			[]int{},
		},
		{
			"a",
			[]string{"a", "a"},
			[]int{},
		},
		{
			"aaaaaa",
			[]string{"aaa","aaa"},
			[]int{0},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := findSubstring(v.s, v.words)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
