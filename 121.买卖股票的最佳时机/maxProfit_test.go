package maxProfit

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct{
		prices	[]int
		want	int
	}{
		{
			[]int{7,1,5,3,6,4},
			5,
		},
		{
			[]int{7,1,6,3,6,4},
			5,
		},

		{
			[]int{7,6,4,3,1},
			0,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := maxProfit(v.prices)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}