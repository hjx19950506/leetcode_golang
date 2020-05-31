package spiralOrder

import (
	"fmt"
	"reflect"
	"testing"
)

//spiralOrder(matrix [][]int) []int
func TestSpiralOrder(t *testing.T) {
	cases := []struct{
		matrix	[][]int
		want	[]int
	}{
		{
				[][]int{{1,2,3},{4,5,6},{7,8,9}},
				[]int{1,2,3,6,9,8,7,4,5},
		},
		{
			[][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}},
			[]int{1,2,3,4,8,12,11,10,9,5,6,7},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := spiralOrder(v.matrix)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}