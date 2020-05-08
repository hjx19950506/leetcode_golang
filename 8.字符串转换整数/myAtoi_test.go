package myAtoi

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	cases := []struct{
		name	string
		str		string
		want	int
	} {
		{
			"test1",
			"42",
			42,
		},
		{
			"test2",
			"   -42",
			-42,
		},
		{
			"test3",
			"4193 with words",
			4193,
		},
		{
			"test4",
			"words and 987",
			0,
		},
		{
			"test5",
			"-91283472332",
			-2147483648,
		},
		{
			"test6",
			"+-2",
			0,
		},
		{
			"test7",
			"+2",
			2,
		},
		{
			"test8",
			"   +0 123",
			0,
		},
		{
			"test9",
			"2147483648",
			2147483647,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := myAtoi(v.str)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
