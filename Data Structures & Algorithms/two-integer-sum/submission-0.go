func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
	for i,v:=range nums{
		_,ok:=m[v]
		if ok{
			return []int{m[v],i}
		}
		m[target-v]=i
	}
	return []int{}
}
