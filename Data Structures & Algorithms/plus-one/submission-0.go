func plusOne(digits []int) []int {
	nd:=make([]int,0)
	s:=0
    for i,v:=range digits{
		s+=v*int(math.Pow(float64(10),float64(len(digits)-i-1)))
	}
	s+=1
	fmt.Println("s=",s)
	for {
		if s==0{
			break
		}
		r:=s%10
		nd=append([]int{r},nd...)
		s=s/10
	}
	return nd
}
