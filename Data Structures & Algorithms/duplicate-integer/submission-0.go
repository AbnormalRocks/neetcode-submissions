func hasDuplicate(nums []int) bool {
	m:=make(map[int]int)
    for _,v:=range nums{
		_,ok:=m[v]
		if ok{
			return true
		}
		m[v]=1
	}
	return false
}
