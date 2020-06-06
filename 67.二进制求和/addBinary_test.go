package addBinary

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	cases := []struct {
		a    string
		b    string
		want string
	}{
		{
			"11",
			"1",
			"100",
		},
		{
			"1010",
			"1011",
			"10101",
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := addBinary(v.a, v.b)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}