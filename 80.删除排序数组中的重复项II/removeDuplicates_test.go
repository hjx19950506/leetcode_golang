package removeDuplicates

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct{
		nums		[]int
		wantNums	[]int
		wantCnt		int
	} {
		{
			[]int{1, 1, 1, 2, 2, 3},
			[]int{1, 1, 2, 2, 3},
			5,
		},
		{
			[]int{0,0,1,1,1,1,2,3,3},
			[]int{0,0,1,1,2,3,3},
			7,
		},
		{
			[]int{0,0,1,1,1,1,2,3,3,3,4},
			[]int{0,0,1,1,2,3,3,4},
			8,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := removeDuplicates(v.nums)
			v.nums = v.nums[:got]
			if got != v.wantCnt || !reflect.DeepEqual(v.nums, v.wantNums) {
				t.Errorf("got %v want %v", v.nums, v.wantNums)
			}
		})
	}
}