package subsets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	cases := []struct{
		nums	[]int
		want	[][]int
	} {
		{
			[]int{1,2,3},
			[][]int{{},{1},{2},{3},{1,2},{1,3},{2,3},{1,2,3}},
		},
		{
			[]int{1,2,3,6},
			[][]int{{},{1},{2},{3},{1,2},{1,3},{2,3},{1,2,3}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := subsets(v.nums)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}