package totalNQueens

import (
	"fmt"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	cases := []struct{
		n		int
		want	int
	}{
		{
			4,
			2,
		},
	}
	for k, v := range cases{
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := totalNQueens(v.n)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}