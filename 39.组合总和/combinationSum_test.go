package combinationSum

import (
	"fmt"
	"reflect"
	"testing"
)

//combinationSum(candidates []int, target int) [][]int
func TestCombinationSum(t *testing.T) {
	cases := []struct{
		candidates	[]int
		target		int
		want		[][]int
	}{
		{
			[]int{2,3,6,7},
			7,
			[][]int{{7},{2,2,3}},
		},
		{
			[]int{2,3,5},
			8,
			[][]int{{2,2,2,2},{2,3,3},{3,5}},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T){
			got := combinationSum(v.candidates, v.target)

			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}