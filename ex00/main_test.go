package main

import (
	"testing"
)

type Test struct {
	arr           []int
	idx, expected int
	err           bool
}

var tests []Test = []Test{
	{arr: []int{1, 2, 3, 4, 5}, idx: 1, expected: 2},
	{arr: []int{1}, idx: 0, expected: 1},
	{arr: []int{1, 2, 3, 4, 5}, idx: -1, expected: 0, err: true},
	{arr: []int{1, 2, 3, 4, 5}, idx: 10, expected: 0, err: true},
	{arr: []int{}, idx: 0, expected: 0, err: true},
}

func TestGetElement(t *testing.T) {
	for n, test := range tests {
		if res, err := getElement(test.arr, test.idx); res != test.expected ||
			(test.err == true && err == nil) {
			if test.err == true && err == nil {
				t.Errorf("test %d failed, expected error", n+1)
			} else {
				t.Errorf("test %d failed, expected output %d, got %d\n", n+1, test.expected, res)
			}
		}
	}
}
