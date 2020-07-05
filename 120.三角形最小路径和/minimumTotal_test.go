package minimumTotal

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	cases := []struct{
		triangle 	[][]int
		want		int
	}{
		{
			[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			11,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := minimumTotal(v.triangle)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}