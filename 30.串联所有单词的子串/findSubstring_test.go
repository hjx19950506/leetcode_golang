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
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := findSubstring(v.s, v.words)
			if reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
