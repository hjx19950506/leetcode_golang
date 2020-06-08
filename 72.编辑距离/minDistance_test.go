package minDistance

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	cases := []struct{
		word1	string
		word2	string
		want	int
	} {
		{
			"horse",
			"ros",
			3,
		},
		{
			"intention",
			"execution",
			5,
		},
		{
			"horse",
			"",
			5,
		},
		{
			"",
			"asd",
			3,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := minDistance(v.word1, v.word2)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}