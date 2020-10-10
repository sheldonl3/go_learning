package list

//Definition for singly-linked list.
//反转链表：递归+迭代

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new_head := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return new_head
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var cur *ListNode
	for {
		if head == nil {
			break
		}
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}
	return cur
}
