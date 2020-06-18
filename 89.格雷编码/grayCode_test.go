package grayCode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	cases := []struct{
		n		int
		want	[]int
	} {
		{
			2,
			[]int{0,1,3,2},
		},
		{
			0,
			[]int{0},
		},
		{
			1,
			[]int{0,1},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := grayCode(v.n)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}