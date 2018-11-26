package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	lenth := len(lists)
	var t *ListNode
	for i := 0; i < lenth; i++ {
		t = mergeTwoLists(t, lists[i])
	}
	return t
}

func mergeTwoLists(listsOne, listsTwo *ListNode) *ListNode {
	if listsOne == nil && listsTwo == nil {
		return nil
	} else if listsOne == nil && listsTwo != nil {
		return listsTwo
	} else if listsOne != nil && listsTwo == nil {
		return listsOne
	}
	var r *ListNode
	var result *ListNode
	var val int
	for listsOne != nil && listsTwo != nil {
		if listsOne.Val >= listsTwo.Val {
			val = listsTwo.Val
			listsTwo = listsTwo.Next
		} else {
			val = listsOne.Val
			listsOne = listsOne.Next
		}
		if r != nil {
			r.Next = &ListNode{Val: val}
			r = r.Next
		} else {
			r = &ListNode{Val: val}
			result = r
		}
	}
	if listsOne == nil {
		r.Next = listsTwo
	} else {
		r.Next = listsOne
	}
	return result
}
