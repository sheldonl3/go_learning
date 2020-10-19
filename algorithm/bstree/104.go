package bstree

import "algorithm/bstree/travel"

type TreeNode = travel.TreeNode

func maxDepth(root *TreeNode) int {
	return help(root, 0)
}

func help(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}
	if root.Right == nil && root.Left == nil {
		return dep + 1
	}
	dep_left, dep_right := dep, dep
	if root.Left != nil {
		dep++
		dep_left = help(root.Left, dep)
		dep--
	}
	if root.Right != nil {
		dep++
		dep_right = help(root.Right, dep)
		dep--
	}
	if dep_left > dep_right {
		return dep_left
	}
	return dep_right
}
