package setZeroes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	cases := []struct{
		matrix	[][]int
		want	[][]int
	} {
		{
			[][]int{{1,1,1},{1,0,1},{1,1,1}},
			[][]int{{1,0,1},{0,0,0},{1,0,1}},
		},
		{
			[][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
			[][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}},
		},
		{
			[][]int{{}},
			[][]int{{}},
		},
		{
			[][]int{{1,1,1},{0,1,2}},
			[][]int{{0,1,1},{0,0,0}},
		},
		{
			[][]int{{1,0,3}},
			[][]int{{0,0,0}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			setZeroes(v.matrix)
			if !reflect.DeepEqual(v.matrix, v.want) {
				t.Errorf("got %v want %v", v.matrix, v.want)
			}
		})
	}
}