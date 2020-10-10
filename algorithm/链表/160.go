package list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cura, curb := headA, headB
	a, b := 0, 0
	for cura != nil {
		a++
		cura = cura.Next
	}
	for curb != nil {
		b++
		curb = curb.Next
	}
	cura, curb = headA, headB
	if a > b {
		for i := 0; i < a-b; i++ {
			cura = cura.Next
		}
	} else if b > a {
		for i := 0; i < b-a; i++ {
			curb = curb.Next
		}
	}
	for {
		if cura == curb {
			return cura
		}
		if cura == nil || curb == nil {
			break
		}
		cura = cura.Next
		curb = curb.Next
	}
	return nil
}
