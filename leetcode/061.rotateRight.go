package leetcode

func rotateRight(head *ListNode, k int) (r *ListNode) {
	if head == nil {
		return nil
	}
	temp := head
	cont := 0
	for temp != nil {
		cont++
		temp = temp.Next
	}
	k = k % cont
	if k == 0 {
		return head
	}
	temp = head
	temphead := head
	i := 0
	for temp.Next != nil && temphead.Next != nil {
		i++
		if i > k {
			temphead = temphead.Next
		}
		temp = temp.Next
	}
	if temp.Next == nil {
		temp.Next = head
		r = temphead.Next
		temphead.Next = nil

	}
	return r
}
