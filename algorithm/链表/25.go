package list

func reverseKGroup(head *ListNode, k int) *ListNode {
	var count = 0
	cur := head
	for cur != nil && count < k { //for 可以直接当成 while 用
		cur = cur.Next
		count++
	}
	if count == k { //正好k个才反转，否则不转
		cur = reverseKGroup(cur, k)
		for count > 0 {
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
			count--
		}
		head = cur
	}
	return head
}
