/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	b:=true
    rec(p,q,&b)
	return b
}
func rec(p,q *TreeNode, b *bool){
	if !*b{
		return
	}
	if p==nil&&q!=nil||p!=nil&&q==nil{
		*b=false
		return
	}
	if p==nil&&q==nil{
		return
	}
	if p.Val!=q.Val{
		*b=false
		return
	}
	rec(p.Left,q.Left,b)
	rec(p.Right,q.Right,b)
	return
}
