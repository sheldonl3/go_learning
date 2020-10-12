package bstree

func Build_tree(s string) TreeNode { //给出层序，构造二叉树，返回root
	if s == "" || s[0] == '#' {
		return TreeNode{}
	}
	root := TreeNode{Val: int(s[0]) - 48} //string里面用uint8表示字符串（asc）
	que := make([]*TreeNode, 0)
	que = append(que, &root)
	i, j := 0, 1 //i是节点队列的index，j是string的index
	for {
		if j == len(s) {
			return root
		}
		node := que[i]
		if node == nil {
			i++
			continue
		}
		if s[j] != '#' {
			node.Left = &TreeNode{Val: int(s[j] - 48)}
		} else {
			node.Left = nil
		}
		que = append(que, node.Left)
		j++
		if j == len(s) {
			return root
		}
		if s[j] != '#' {
			node.Right = &TreeNode{Val: int(s[j] - 48)}
		} else {
			node.Right = nil
		}
		que = append(que, node.Right)
		j++
		i++
	}
}
