package travel

//层序遍历
func cengxu(p *TreeNode) []int {
	res := make([]int, 0)
	if p == nil {
		return res
	}
	queue := []*TreeNode{p}
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res = append(res, queue[0].Val)
			queue = queue[1:]
		}
	}
	return res
}

//先序遍历
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)

	for len(queue) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1].Right
		queue = queue[:len(queue)-1]
	}
	return result
}

//中序遍历
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)

	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}

		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	var lastVisited *TreeNode

	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		n := queue[len(queue)-1]
		if n.Right == nil || n.Right == lastVisited {
			result = append(result, n.Val)
			queue = queue[:len(queue)-1]
			lastVisited = n
		} else {
			root = n.Right
		}
	}

	return result
}
