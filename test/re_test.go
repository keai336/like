package test

import (
	"regexp"
	"testing"
)

func TestRe(t *testing.T) {
	a := "[^a]+abc.+?j"
	b := regexp.MustCompile(a)
	findstring := "hgkhnxabcdfjiexzjlfssdvxabcsdlfj"
	c := b.FindAllStringIndex(findstring, -1)
	t.Log(c)
	for _, v := range c {
		t.Log(findstring[v[0]:v[1]])
	}
}
