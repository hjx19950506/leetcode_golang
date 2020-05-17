package divide

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		dividend 	int
		divisor		int
		want 		int
	} {
		{
			10,
			3,
			3,
		},
		{
			7,
			-3,
			-2,
		},
		{
			0,
			-3,
			0,
		},
		{
			1,
			1,
			1,
		},
		{
			-1,
			1,
			-1,
		},
		{
			2147483647,
			1,
			2147483647,
		},
		{
			-2147483648,
			-2,
			1073741824,
		},
		{
			-2147483648,
			-2147483648,
			1,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := divide(v.dividend, v.divisor)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
