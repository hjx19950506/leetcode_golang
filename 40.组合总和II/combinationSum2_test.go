package combinationSum2

import (
	"fmt"
	"reflect"
	"testing"
)

//combinationSum(candidates []int, target int) [][]int
func TestCombinationSum2(t *testing.T) {
	cases := []struct{
		candidates	[]int
		target		int
		want		[][]int
	}{
		{
			[]int{2,3,6,7},
			7,
			[][]int{{7}},
		},
		{
			[]int{2,3,5},
			8,
			[][]int{{3,5}},
		},
		{
			[]int{10,1,2,7,6,1,5},
			8,
			[][]int{{1,7},{1,2,5},{2,6},{1,1,6}},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T){
			got := combinationSum2(v.candidates, v.target)

			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}