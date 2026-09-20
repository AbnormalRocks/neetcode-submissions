func climbStairs(n int) int {
	w:=0
	m:=make(map[int]int)
	m[1]=1
	m[2]=2
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
    rec(n,&w,m)
	return m[n]
}
func rec(n int,w *int,m map[int]int)int{
	if p,ok:=m[n];ok{
		return p
	}
	m[n]=rec(n-1,w,m)+rec(n-2,w,m)
	return m[n]
}
