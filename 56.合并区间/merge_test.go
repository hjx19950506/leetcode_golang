package merge

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct{
		intervals	[][]int
		want		[][]int
	} {
		{
			[][]int{{1,3},{2,6},{8,10},{15,18}},
			[][]int{{1,6},{8,10},{15,18}},
		},
		{
			[][]int{{1,4},{4,5}},
			[][]int{{1,5}},
		},
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{{1,4},{2,3}},
			[][]int{{1,4}},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := merge(v.intervals)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}