package insert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct{
		intervals	[][]int
		newInterval	[]int
		want		[][]int
	} {
		{
			[][]int{{1,3},{2,6},{8,10},{15,18}},
			[]int{9,10},
			[][]int{{1,6},{8,10},{15,18}},
		},
		{
			[][]int{{1,4},{4,5}},
			[]int{6,7},
			[][]int{{1,5},{6,7}},
		},
		{
			[][]int{},
			[]int{1,2},
			[][]int{{1,2}},
		},
		{
			[][]int{{1,4},{2,3}},
			[]int{3,4},
			[][]int{{1,4}},
		},
		{
			[][]int{{1,5}},
			[]int{0,3},
			[][]int{{0,5}},
		},
		{
			[][]int{{0,5},{9,12}},
			[]int{7,16},
			[][]int{{0,5},{7,16}},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := insert(v.intervals, v.newInterval)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}