func lastStoneWeight(stones []int) int {
	st:=make([]int,0,0)
	for _,v:=range stones{
		st=getSorted(st,v)
	}
	for {
		if len(st)<2{
			break
		}
		d:=st[len(st)-1]-st[len(st)-2]
		if d!=0{
			st=getSorted(st[:len(st)-2],d)
		} else {
			st=st[:len(st)-2]
		}
	}
	if len(st)==1{
		return st[0]
	}
	return 0
}
func getSorted(st []int,v int) []int{
	m:=0
	l:=0
	r:=len(st)-1
	if len(st)==0{
		st=append(st,v)
		return st
	}
	for{
		m=(l+r)/2
		if m>=len(st)||r<=l{
			break
		}
		if st[m]==v{
			break
		}
		if v>st[m]{
			l=m+1
		} else {
			r=m
		}
	}
	if v>st[m]{
		if m>=len(st)-1{
			st=append(st,v)
		} else {
			st=append(st[:m+1],append([]int{v},st[m+1:]...)...)
		}
	} else {
		st=append(st[:m],append([]int{v},st[m:]...)...)
	}
	return st
}
