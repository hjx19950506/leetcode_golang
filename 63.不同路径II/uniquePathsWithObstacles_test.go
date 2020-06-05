package uniquePathsWithObstacles

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	cases := []struct {
		obstacleGrid	[][]int
		want			int
	} {
		{
			[][]int{{0,0,0},{0,1,0},{0,0,0}},
			2,
		},
		{
			[][]int{{1,0}},
			0,
		},
		{
			[][]int{{1},{0}},
			0,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := uniquePathsWithObstacles(v.obstacleGrid)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}