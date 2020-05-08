package addTwoNumbers

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		name	string
		nums	[]int
		target	int
		want		[]int
	}{
		{
			"test1",
			[]int{2,7,11,15},
			9,
			[]int{0, 1},
		},
		{
			"test2",
			[]int{1,2},
			9,
			nil,
		},
		{
			"test3",
			[]int{3,2,4},
			6,
			[]int{1,2},
		},
	}

	for k, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := twoSum(v.nums, v.target)
			if !assertEqual(t, got, v.want) {
				t.Errorf("case %d error, nums are %v, target is %d, want is %d, but got %d",
					k, v.nums, v.target, v.want, got)
			}
		})
	}
}

func assertEqual(t *testing.T, pair1, pair2 []int) bool {
	t.Helper()
	sort.Ints(pair1)
	sort.Ints(pair2)
	if reflect.DeepEqual(pair1, pair2) {
		return true
	}
	return false
}