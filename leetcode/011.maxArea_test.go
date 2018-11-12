package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	r := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if r != 49 {
		t.Error(49, r)
	}
}
