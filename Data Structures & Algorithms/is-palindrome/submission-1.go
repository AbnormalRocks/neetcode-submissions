func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	p1:=0
	for {
		//fmt.Println("p1",p1," s[p1]",s[p1])
			if p1>=len(s){
				break
			}
			asc:=s[p1]
			if (asc>=97 && asc<=120) || (asc>=48 && asc<=57){
				//fmt.Println("s1")
				break
			} 
			//fmt.Println("s1b")
			p1++
		}
	p2:=len(s)-1
	for {
		//fmt.Println("p2",p2," s[p2]",s[p2])
			if p2<0{
				break
			}
			asc:=s[p2]
			if (asc>=97 && asc<=120) || (asc>=48 && asc<=57){
				//fmt.Println("s2")
				break
			} 
			//fmt.Println("s2b")
			p2--
		}
	isP:=true
	for {
		//fmt.Println("p1=",p1," p2=",p2, " s[p1]",s[p1]," s[p2]",s[p2])
		if p2<p1 {
			break
		}
		if s[p1]!=s[p2]{
			isP=false
			break
		}
		p1++
		for {
			if p1>=len(s){
				break
			}
			asc:=s[p1]
			if (asc>=97 && asc<=120) || (asc>=48 && asc<=57){
				break
			} 
			p1++
		}
		p2--
		for {
			if p2<0{
				break
			}
			asc:=s[p2]
			if (asc>=48 && asc<=57) || (asc>=97 && asc<=120){
				break
			} 
			p2--
		}
	}
	return isP
}
