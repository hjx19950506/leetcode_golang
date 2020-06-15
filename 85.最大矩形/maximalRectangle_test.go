package maximalRectangle

import (
	"fmt"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	cases := []struct{
		matrix		[][]byte
		want		int
	} {
		{
			[][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'}, {'1','1','1','1','1'}, {'1','0','0','1','0'}},
			6,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := maximalRectangle(v.matrix)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}