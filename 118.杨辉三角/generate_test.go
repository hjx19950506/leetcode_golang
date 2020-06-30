package generate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct{
		numRows 	int
		want		[][]int
	}{
		{
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			1,
			[][]int{{1}},
		},
		{
			0,
			[][]int{},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := generate(v.numRows)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
