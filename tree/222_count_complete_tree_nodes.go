package tree

func countNodes(root *TreeNode) int {
	if root != nil {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
	return 0
}
