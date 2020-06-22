package restoreIpAddresses

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	cases := []struct{
		s		string
		want	[]string
	} {
		{
			"25525511135",
			[]string{"255.255.11.135", "255.255.111.35"},
		},
		{
			"0000",
			[]string{"0.0.0.0"},
		},
		{
			"010010",
			[]string{"0.10.0.10", "0.100.1.0"},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := restoreIpAddresses(v.s)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}