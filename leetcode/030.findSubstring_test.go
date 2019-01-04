package leetcode

import "testing"

func Test_findSubstring(t *testing.T) {
	r := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	if len(r) != 2 {
		t.Error(r, "[0,9]")
	}
	r = findSubstring("wordgoodstudentgoodword", []string{"word", "student"})
	if len(r) != 0 {
		t.Error("[]")
	}
}
