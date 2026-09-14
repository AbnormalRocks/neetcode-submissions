/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	b:=true
    rec(root, &b)
	return b
}
func rec(t *TreeNode,b *bool)int{
	if !*b{
		return 0
	}
	if t==nil{
		return 0
	}
	lh:=rec(t.Left,b)
	rh:=rec(t.Right,b)
	diff:=0
	if lh>rh{
		diff=lh-rh
	} else {
		diff=rh-lh
	}
	if diff>1{
		*b=false
	}
	h:=1+max(lh,rh)
	return h
}
