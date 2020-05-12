package letterCombinations

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	cases := []struct{
		digits 	string
		want 	[]string
	} {
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"",
			[]string{},
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := letterCombinations(v.digits)
			sort.Strings(got)
			sort.Strings(v.want)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %#v want %#v", got, v.want)
			}
		})
	}
}
