package maxSubArray

import (
	"fmt"
	"testing"
)

//maxSubArray(nums []int)
func TestMaxSubArray(t *testing.T) {
	cases := []struct{
		nums	[]int
		want	int
	}{
		{
			[]int{-2,1,-3,4,-1,2,1,-5,4},
			6,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := maxSubArray(v.nums)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}