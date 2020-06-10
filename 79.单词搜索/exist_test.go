package exist

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	cases := []struct{
		board		[][]byte
		word		string
		want		bool
	} {
		{
			[][]byte{{'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}},
			"ABCCED",
			true,
		},
		{
			[][]byte{{'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}},
			"SEE",
			true,
		},
		{
			[][]byte{{'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}},
			"ABCB",
			false,
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := exist(v.board, v.word)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}