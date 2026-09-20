func isHappy(n int) bool {
	m:=make(map[int]bool)
	m[n]=true
	c:=n
    for {
		r:=0
		s:=0
		for {
			if c==0{
				break
			}
			r=c%10
			s=s+(r*r)
			c=c/10
		}
		if s==1{
			return true
		}
		if _,ok:=m[s];ok{
			return false
		}
		m[s]=true
		c=s
	}
	return false
}
