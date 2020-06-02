package lengthOfLastWord

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T){
	cases := []struct{
		s 		string
		want	int
	} {
		{
			"Hello World",
			5,
		},
		{
			"Hell0 Worl",
			4,
		},
		{
			"Hello World  asd a ",
			1,
		},
		{
			"    ",
			0,
		},
		{
			"    a",
			1,
		},
		{
			"a",
			1,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := lengthOfLastWord(v.s)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}