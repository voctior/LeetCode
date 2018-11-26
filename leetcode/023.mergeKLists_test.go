package leetcode

import "testing"

func Test_mergeKLists(t *testing.T) {
	p := &ListNode{Val: 1}
	p0 := p
	p.Next = &ListNode{Val: 4}
	p = p.Next
	p.Next = &ListNode{Val: 5}

	p = &ListNode{Val: 1}
	p1 := p
	p.Next = &ListNode{Val: 3}
	p = p.Next
	p.Next = &ListNode{Val: 4}

	p = &ListNode{Val: 2}
	p2 := p
	p.Next = &ListNode{Val: 6}

	r := mergeKLists([]*ListNode{p0, p1, p2})
	if r == nil {
		t.Error("数据对应失败")
	}
}
