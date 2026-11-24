package main

import (
	"reflect"
	"testing"
)

func TestPathToNode(t *testing.T) {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, nil, nil))
	cases := []struct {
		target int
		want   []int
	}{
		{5, []int{1, 2, 5}},
		{4, []int{1, 2, 4}},
		{3, []int{1, 3}},
		{1, []int{1}},
		{99, nil},
	}
	for _, c := range cases {
		if got := PathToNode(root, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("PathToNode(%d) = %v, want %v", c.target, got, c.want)
		}
	}
	if got := PathToNode(nil, 1); got != nil {
		t.Errorf("nil tree: %v", got)
	}
}
