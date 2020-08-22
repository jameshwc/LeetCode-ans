
/* Slow but easy-to-understand version */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA, lenB int
	tempA, tempB := headA, headB
	for tempA != nil {
		lenA += 1
		tempA = tempA.Next
	}
	for tempB != nil {
		lenB += 1
		tempB = tempB.Next
	}
	tempA, tempB = headA, headB
	if diff := lenA - lenB; diff > 0 {
		for diff != 0 {
			tempA = tempA.Next
			diff--
		}
	} else if diff < 0 {
		for diff != 0 {
			tempB = tempB.Next
			diff++
		}
	}
	for tempA != nil && tempB != nil {
		if tempA == tempB {
			return tempA
		}
		tempA = tempA.Next
		tempB = tempB.Next
	}
	return nil
}

/* Faster */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n1, n2 := headA, headB
	if n1 == nil || n2 == nil {
		return nil
	}
	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
		if n1 == n2 {
			return n1
		}
		if n1 == nil {
			n1 = headB
		}
		if n2 == nil {
			n2 = headA
		}
	}
	return n1
}