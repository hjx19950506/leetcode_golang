package permuteUnique

import (
	"fmt"
	"reflect"
	"testing"
)

//permute(nums []int) [][]int
func TestPermuteUnique(t *testing.T) {
	cases := []struct{
		nums	[]int
		want 	[][]int
	} {
		{
			[]int{1,2,3},
			[][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}},
		},
		{
			[]int{1,1,2},
			[][]int{{1,1,2},{1,2,1},{2,1,1}},
		},
		{
			[]int{1,1,2,2},
			[][]int{{1,1,2,2},{1,2,1,2},{1,2,2,1},{2,1,1,2},{2,1,2,1},{2,2,1,1}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T){
			got := permuteUnique(v.nums)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}