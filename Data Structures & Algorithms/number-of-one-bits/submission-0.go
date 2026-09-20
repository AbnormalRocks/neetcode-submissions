func hammingWeight(n int) int {
	k:=0
	for {
		if n==0{
			break
		}
		r:=n%2
		if r==1{
			k++
		}
		n=n/2
	}
	return k
}
