/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	m:=0
    rec(root,&m)
	return m
}
func rec(t *TreeNode,m *int)int{
	if t==nil{
		return 0
	}
	lh:=rec(t.Left,m)
	rh:=rec(t.Right,m)
	*m=max(*m,lh+rh)
	h:=1+max(lh,rh)
	fmt.Println("t.Val=",t.Val," lh=",lh," rh=",rh," h=",h)
	return h
}
