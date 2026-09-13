/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var h *ListNode
	
	if list1==nil||list2==nil{
		if list1==nil{
		if list2==nil{
			return nil
		}
		h=list2
		list2=list2.Next
	} else if list2==nil{
		if list1==nil{
			return nil
		}
		h=list1
		list1=list1.Next
	}} else if list1.Val<list2.Val{
		h=list1
		list1=list1.Next
	} else {
		h=list2
		list2=list2.Next
	}
	//h=h.Next
	l:=h
	//k:=0
	//fmt.Println("h.Val=",h.Val)
	for {
		if list1==nil{
			for {
				if list2==nil{
					return h
				}
				l.Next=list2
				list2=list2.Next
				l=l.Next
			}
			return h
		}
		if list2==nil{
			for {
				if list1==nil{
					return h
				}
				l.Next=list1
				list1=list1.Next
				l=l.Next
			}
			return h
		}
		if list1.Val<list2.Val{
			l.Next=list1
			list1=list1.Next
		} else {
			l.Next=list2
			list2=list2.Next
		}
		l=l.Next
		//fmt.Println("l.Val=",l.Val," list1.Val=",list1.Val," list2.Val",list2.Val)

	}
	return h
}
