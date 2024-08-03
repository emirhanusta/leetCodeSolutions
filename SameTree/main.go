package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	p := TreeNode{1, nil, nil}
	q := TreeNode{1, nil, nil}

	result := isSameTree(&p, &q)
	println(result)
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return p == q
	}
}
