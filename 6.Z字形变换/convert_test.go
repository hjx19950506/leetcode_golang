package convert

import (
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []struct{
		name	string
		s		string
		numRows	int
		want	string
	} {
		{
			"test1",
			"LEETCODEISHIRING",
			3,
			"LCIRETOESIIGEDHN",
		},
		{
			"test2",
			"LEETCODEISHIRING",
			4,
			"LDREOEIIECIHNTSG",
		},
		{
			"test3",
			"1",
			4,
			"1",
		},
		{
			"test4",
			"AB",
			1,
			"AB",
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			got := convert(v.s, v.numRows)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
