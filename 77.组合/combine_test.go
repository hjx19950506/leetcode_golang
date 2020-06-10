package combine

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	cases := []struct{
		n		int
		k		int
		want	[][]int
	} {
		{4,
			2,
			[][]int{{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}},
		},
		{4,
			3,
			[][]int{{1,2,3},{1,2,4},{1,3,4},{2,3,4}},
		},
		{5,
			1,
			[][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{5,
			0,
			[][]int{},
		},
		{0,
			3,
			[][]int{},
		},
		{6,
			6,
			[][]int{{1,2,3,4,5,6}},
		},
		{5,
			3,
			[][]int{{1,2,3},{1,2,4},{1,2,5},{1,3,4},{1,3,5},{1,4,5},{2,3,4},{2,3,5},{2,4,5},{3,4,5}},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := combine(v.n, v.k)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}