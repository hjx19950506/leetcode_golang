package generateParenthesis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct{
		n 		int
		want	[]string
	} {
		{
			3,
			[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := generateParenthesis(v.n)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %#v want %#v", got, v.want)
			}
		})
	}
}

