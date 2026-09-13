/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	prev=nil
	var tmp *ListNode
	tmp=head
    for {
		if tmp==nil{
			break
		}
		next:=tmp.Next
		tmp.Next=prev
		prev=tmp
		tmp=next
	}
	return prev
}

func print(n *ListNode){
	for {
		if n==nil{
			break
		}
		fmt.Print("Val:",n.Val," Next:",n.Next," ")
		n=n.Next
	}
}
