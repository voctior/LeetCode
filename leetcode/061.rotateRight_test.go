package leetcode

import "testing"

func Test_rotateRight(t *testing.T) {
	p := &ListNode{Val: 1}
	p0 := p
	p.Next = &ListNode{Val: 4}
	p = p.Next
	p.Next = &ListNode{Val: 5}
	r := rotateRight(p0, 4)
	if r.Val != 5 {
		t.Error("错误", r)
	}

}
