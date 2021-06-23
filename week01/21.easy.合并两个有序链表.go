package week01

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// https://leetcode-cn.com/submissions/detail/185192240/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -int(^uint(0)>>1) - 1,
	}
	curr := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			break
		}
		if l2 == nil {
			curr.Next = l1
			break
		}
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
		curr.Next = nil
	}
	return head.Next
}
