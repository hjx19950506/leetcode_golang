package myPow

import (
	"fmt"
	"testing"
)

//myPow(x float64, n int) float64
func TestMyPow(t *testing.T) {
	cases := []struct{
		x		float64
		n		int
		want	float64
	}{
		{	2.0,
			10,
			1024.0,
		},
		{	2.1,
			3,
			9.26100,
		},
		{	2.0,
			-2,
			0.25000,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T){
			got := myPow(v.x, v.n)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}