package getPermutation

import (
	"fmt"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	cases := []struct{
		n		int
		k		int
		want	string
	}{
		{
			1,
			1,
			"1",
		},
		{
			3,
			3,
			"213",
		},
		{
			3,
			2,
			"132",
		},
		{
			4,
			9,
			"2314",
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := getPermutation(v.n, v.k)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}