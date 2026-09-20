type KthLargest struct {
    snums []int
    k int
}


func Constructor(k int, nums []int) KthLargest {
    snums:=make([]int,0,0)
    for _,v:=range nums{
        snums=getSorted(snums,v)
    }
    return KthLargest{snums,k}
}


func (this *KthLargest) Add(val int) int {
    this.snums=getSorted(this.snums,val)
    return this.snums[len(this.snums)-this.k]
}

func getSorted(snums []int,v int)[]int{
    m:=0
    l:=0
    r:=len(snums)-1
    if len(snums)==0{
        snums=append(snums,v)
        return snums
    }
    for {
        m=(l+r)/2
        if m>=len(snums)||r<=l{
            break
        }
        if snums[m]==v{
            break
        }
        if snums[m]<v{
            l=m+1
        } else {
            r=m
        }
    }
    if v>snums[m]{
        if m>=len(snums)-1{
            snums=append(snums,v)
        } else {
            snums=append(snums[:m+1],append([]int{v},snums[m+1:]...)...)
        }
    } else {
        snums=append(snums[:m],append([]int{v},snums[m:]...)...)
    }
    return snums
}
