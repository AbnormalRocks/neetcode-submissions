func minCostClimbingStairs(cost []int) int {
    m:=make(map[int]int,len(cost))
	m[0]=0
	m[1]=0
	n:=2
	for {
		if n==len(cost)+1{
			break
		}
		m[n]=min(m[n-2]+cost[n-2],m[n-1]+cost[n-1])
		n++
	}
	return m[n-1]
}
