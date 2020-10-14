package bstree

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res := make([]*TreeNode, 0)
	m := make(map[string]int, 0)
	houxu(root, m, &res)
	return res

}

func houxu(root *TreeNode, m map[string]int, res *[]*TreeNode) string { //slice也是值传递，append的时候必须使用指针
	if root == nil {
		return ""
	}
	left := houxu(root.Left, m, res)
	right := houxu(root.Right, m, res)
	tmp := left + "," + right + "," + strconv.Itoa(root.Val)
	val, ok := m[tmp]
	if ok == false {
		m[tmp] = 1
	} else if val == 1 {
		*res = append(*res, root)
		m[tmp]++
	} else {
		m[tmp]++
	}
	return tmp
}
