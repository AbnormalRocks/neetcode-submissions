func isAnagram(s string, t string) bool {
	m1:=make(map[rune]int)
    m2:=make(map[rune]int)
    for _,r1:=range s{
        m1[r1]+=1
    }
    for _,r2:=range t{
        m2[r2]+=1
    }
    for k,v:=range m1{
        c,ok:=m2[k]
        if !ok || v!=c{
            return false
        }
    }
    for k,v:=range m2{
        c,ok:=m1[k]
        if !ok || v!=c{
            return false
        }
    }
    return true
}
