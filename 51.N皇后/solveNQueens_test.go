package solveNQueens

import (
	"fmt"
	"reflect"
	"testing"
)

//solveNQueens(n int) [][]string
func TestSolveNQueens(t *testing.T) {
	cases := []struct{
		n		int
		want	[][]string
	}{
		{
			4,
			[][]string{
				{".Q..","...Q","Q...","..Q."},
				{"..Q.","Q...","...Q",".Q.."},
			},
		},
	}
	for k, v := range cases{
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := solveNQueens(v.n)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}