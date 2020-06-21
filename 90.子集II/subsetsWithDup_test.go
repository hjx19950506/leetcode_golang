package subsetsWithDup

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	cases := []struct{
		nums	[]int
		want	[][]int
	} {
		{
			[]int{1,2,2},
			[][]int{{},{1},{2},{1,2},{2,2},{1,2,2}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := subsetsWithDup(v.nums)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}