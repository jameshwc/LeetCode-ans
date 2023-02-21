package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, answerRoot, answerCur, cur *ListNode
	prev = head
	cur = head.Next
	isRootSet := false
	isDuplicate := false
	for ; cur != nil; cur = cur.Next {
		if prev.Val == cur.Val {
			isDuplicate = true
			continue
		}

		if isDuplicate {
			isDuplicate = false
			prev = cur
			continue
		}

		if !isRootSet {
			answerRoot = prev
			answerCur = answerRoot
			isRootSet = true
		} else {
			answerCur.Next = prev
			answerCur = prev
		}

		prev = cur
	}

	if !isDuplicate {
		if !isRootSet {
			return prev
		}

		answerCur.Next = prev

	} else if !isRootSet {
		return nil
	} else {
		answerCur.Next = nil
	}

	return answerRoot
}

func refactor(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var root ListNode
	root.Next = head
	root.Val = -101 // impossible value

	prev := root
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}

		cur = cur.Next
	}

	return root.Next
}
