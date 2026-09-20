func countBits(n int) []int {
	ks:=make([]int,n+1)
	for i:=0;i<=n;i++{
		t:=i
		k:=0
		for {
			if t==0{
				break
			}
			r:=t%2
			if r==1{
				k++
			}
			t=t/2
		}
		ks[i]=k
	}
	return ks
}
