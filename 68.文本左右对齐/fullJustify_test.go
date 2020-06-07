package fullJustify

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	cases := []struct {
		words    	[]string
		maxWidth	int
		want		[]string
	} {
		{
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			[]string{"What","must","be","acknowledgment","shall","be"},
			16,
			[]string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			[]string{"Science","is","what","we","understand","well","enough","to","explain", "to","a","computer.","Art","is","everything","else","we","do"},
			20,
			[]string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := fullJustify(v.words, v.maxWidth)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %#v want %#v", got, v.want)
			}
		})
	}
}