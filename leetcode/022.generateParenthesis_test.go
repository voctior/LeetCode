package leetcode

import "testing"

func Test_generateParenthesis(t *testing.T) {
	r := generateParenthesis(3)
	if len(r) != 5 {
		t.Error(r)
	}
}
