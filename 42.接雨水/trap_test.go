package trap

import (
	"fmt"
	"testing"
)

//trap(height []int)
func TestTrap(t *testing.T) {
	cases := []struct{
		height 	[]int
		want 	int
	} {
		{
			[]int{0,1,0,2,1,0,1,3,2,1,2,1},
			6,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := trap(v.height)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}