package generateMatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	cases := []struct {
		n    int
		want [][]int
	}{
		{
			3,
			[][]int{{1,2,3},{8,9,4},{7,6,5}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := generateMatrix(v.n)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}