package simplifyPath

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	cases := []struct{
		path	string
		want	string
	} {
		{
			"/home/",
			"/home",
		},
		{
			"/../",
			"/",
		},
		{
			"/home//foo/",
			"/home/foo",
		},
		{
			"/a/./b/../../c/",
			"/c",
		},
		{
			"/a/../../b/../c//.//",
			"/c",
		},
		{
			"/a//b////c/d//././/..",
			"/a/b/c",
		},
		{
			"////",
			"/",
		},
		{
			"",
			"/",
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprint("test", k+1), func(t *testing.T) {
			got := simplifyPath(v.path)
			if got != v.want {
				t.Errorf("got %v want %v", got ,v.want)
			}
		})
	}
}