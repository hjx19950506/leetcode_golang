package multiply

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	cases := []struct{
		num1 	string
		num2	string
		want	string
	}{
		{
			"0",
			"0",
			"0",
		},
		{
			"2",
			"3",
			"6",
		},
		{
			"123",
			"456",
			"56088",
		},
		{
			"123456789",
			"987654321",
			"121932631112635269",
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T){
			got := multiply(v.num1, v.num2)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
