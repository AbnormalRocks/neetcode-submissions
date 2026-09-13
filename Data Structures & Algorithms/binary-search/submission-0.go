func search(nums []int, target int) int {
	p1:=0
	p2:=len(nums)-1
	for {
		if p2<=p1{
			if nums[p1]==target{
				return p1
			}
			break
		}
		mid:=(p1+p2)/2
		if nums[mid]==target{
			return mid
		}
		if target>nums[mid]{
			p1=mid+1
		} else {
			p2=mid
		}
	}
	return -1
}
