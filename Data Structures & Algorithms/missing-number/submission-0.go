func missingNumber(nums []int) int {
	s:=make([]int,len(nums)+1)
	for _,v:=range nums{
		s[v]=1
	}
	for i,v:=range s{
		if v==0{
			return i
		}
	}
	return 0
}
