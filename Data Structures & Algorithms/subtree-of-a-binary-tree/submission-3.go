/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	f:=false
	m:=true
	rec(root,subRoot,&f,&m)
	if f{
		return m
	}
	return false
}
func rec(r,s *TreeNode,f,m *bool){
	if *f&&*m{
		return
	}
	if r==nil{
		return
	}
	if r.Val==s.Val{
		*m=true
		fmt.Println("f r.Val=",r.Val," s.Val=",s.Val)
		*f=true
		same(r,s,m)
		if *m{
			return
		} else {
			*f=false
		}
	}
	if !*f{
		rec(r.Left,s,f,m)
		rec(r.Right,s,f,m)
	}
	return
}
func same(r,s *TreeNode,m *bool){
	if !*m{
		fmt.Println("false1")
		return
	}
	if r==nil&&s!=nil||r!=nil&&s==nil{
		fmt.Println("false2")
		*m=false
		return
	}
	if r==nil&&s==nil{
		return
	}
	if r.Val!=s.Val{
		fmt.Println("um r.Val=",r.Val," s.Val=",s.Val)
		*m=false
		return
	}
	same(r.Left,s.Left,m)
	same(r.Right,s.Right,m)
	return
}

