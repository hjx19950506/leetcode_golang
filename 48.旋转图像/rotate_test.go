package rotate

import (
	"fmt"
	"reflect"
	"testing"
)

//rotate(matrix [][]int)
func TestRotate(t *testing.T) {
	cases := []struct{
		matrix	[][]int
		want 	[][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}},
			[][]int{{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T){
			rotate(v.matrix)
			if !reflect.DeepEqual(v.matrix, v.want) {
				t.Errorf("got %v want %v", v.matrix, v.want)
			}
		})
	}
}