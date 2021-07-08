package week03

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// https://leetcode-cn.com/submissions/detail/182616563/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	mid := 0
	for k, v := range inorder {
		if root.Val == v {
			mid = k
			break
		}
	}
	root.Left = buildTree(inorder[:mid], postorder[:mid])
	root.Right = buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])
	return root
}
