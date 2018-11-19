package leetcode

import "testing"

func Test_letterCombinations(t *testing.T) {
	r := letterCombinations("345")
	if r == nil {
		t.Error(r)
	}
}
