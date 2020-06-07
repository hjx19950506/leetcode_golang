package mySqrt

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	cases := []struct {
		x		int
		want	int
	} {
		{
			4,
			2,
		},
		{
			8,
			2,
		},
		{
			2401,
			49,
		},
		{
			0,
			0,
		},
		{
			2147395599,
			46339,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := mySqrt(v.x)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}