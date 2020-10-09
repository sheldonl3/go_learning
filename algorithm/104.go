package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type item struct {
	idx int
	*TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	queue := []item{{0, root}}
	for len(queue) > 0 {
		lenght := queue[len(queue)-1].idx - queue[0].idx + 1
		if lenght > res {
			res = lenght
		}
		tmp := make([]item, 0) //tmp =  append([]item{}) 切片清除0
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, item{2 * node.idx, node.Left})
			}
			if node.Right != nil {
				tmp = append(tmp, item{2*node.idx + 1, node.Right})
			}
		}
		queue = tmp
	}
	return res
}
