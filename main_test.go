package main

import (
	"reflect"
	"testing"
)

func TestReorder(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		l := fromSlice(c.in)
		Reorder(l)
		got := toSlice(l)
		if len(c.in) == 0 {
			if l != nil {
				t.Error("nil list should stay nil")
			}
			continue
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Reorder(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
