package groupAnagrams

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	cases := []struct{
		strs 	[]string
		want 	[][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"ate","eat","tea"},{"nat","tan"},{"bat"}},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T){
			got := groupAnagrams(v.strs)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v wan %v", got, v.want)
			}
		})
	}
}