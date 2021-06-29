package week02

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// https://leetcode-cn.com/submissions/detail/190793947/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	a := mergeKLists(lists[:mid])
	b := mergeKLists(lists[mid:])
	return merge(a, b)
}

func merge(a, b *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for a != nil || b != nil {
		if a == nil {
			p.Next = b
			break
		}
		if b == nil {
			p.Next = a
			break
		}
		if a.Val < b.Val {
			p.Next = a
			a = a.Next
		} else {
			p.Next = b
			b = b.Next
		}
		p = p.Next
		p.Next = nil
	}
	return head.Next
}
