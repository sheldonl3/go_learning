package bintree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//先序遍历
func preorder(p *TreeNode, res *[]int) {
	if p != nil {
		*res = append(*res, p.Val)
		preorder(p.Left, res)
		preorder(p.Right, res)
	}
}

//中序遍历
func inorder(p *TreeNode, res *[]int) {
	if p != nil {
		inorder(p.Left, res)
		*res = append(*res, p.Val)
		inorder(p.Right, res)
	}
}

//后序遍历
func postorder(p *TreeNode, res *[]int) {
	if p != nil {
		postorder(p.Left, res)
		postorder(p.Right, res)
		*res = append(*res, p.Val)
	}
}
