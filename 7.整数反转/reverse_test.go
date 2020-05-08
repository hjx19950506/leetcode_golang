package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct{
		name	string
		x		int
		want	int
	} {
		{
			"test1",
			123,
			321,
		},
		{
			"test2",
			-123,
			-321,
		},
		{
			"test3",
			120,
			21,
		},
		{
			"test4",
			2147483646,
			0,
		},
		{
			"test5",
			1534236469,
			0,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := reverse(v.x)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
