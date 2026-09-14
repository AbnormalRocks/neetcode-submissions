/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    d:=0
	m:=0
	rec(root,&d,&m)
	return m
}

func rec(r *TreeNode,d,m *int)*TreeNode{
	if r==nil{
		return r
	}
	*d++
	if *d>*m{
		*m=*d
	}
	rec(r.Left,d,m)
	rec(r.Right,d,m)
	*d--
	return r
}
