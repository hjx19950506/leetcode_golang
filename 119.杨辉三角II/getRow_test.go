package getRow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	cases := []struct{
		rowIndex	int
		want		[]int
	} {
		{
			0,
			[]int{1},
		},
		{
			3,
			[]int{1,3,3,1},
		},
		{
			4,
			[]int{1,4,6,4,1},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := getRow(v.rowIndex)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got , v.want)
			}
		})
	}
}
