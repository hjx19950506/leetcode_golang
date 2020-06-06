package plusOne

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		digits []int
		want   []int
	}{
		{
			[]int{1,2,3},
			[]int{1,2,4},
		},
		{
			[]int{4,3,2,1},
			[]int{4,3,2,2},
		},
		{
			[]int{1,2,9},
			[]int{1,3,0},
		},
		{
			[]int{9,9,9,9},
			[]int{1,0,0,0,0},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := plusOne(v.digits)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}