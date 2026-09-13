/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	m:=make(map[int]bool)
	for {
		if head==nil{
			break
		}
		_,ok:=m[head.Val]
		if ok{
			if head.Next==nil{
				return false
			} else {
				return true
			}
		}
		m[head.Val]=true
		head=head.Next
	}
	return false
}
