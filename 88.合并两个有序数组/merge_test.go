package merge

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct{
		nums1		[]int
		m			int
		nums2		[]int
		n			int
		want		[]int
	} {
		{
			[]int{1,2,3,0,0,0},
			3,
			[]int{2,5,6},
			3,
			[]int{1,2,2,3,5,6},
		},
		{
			[]int{1,2,3,4,0,0,0},
			4,
			[]int{2,5,6},
			3,
			[]int{1,2,2,3,4,5,6},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			merge(v.nums1, v.m, v.nums2, v.n)
			if !reflect.DeepEqual(v.nums1, v.want) {
				t.Errorf("got %v want %v", v.nums1, v.want)
			}
		})
	}
}