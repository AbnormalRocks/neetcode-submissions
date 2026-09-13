func isValid(s string) bool {
    stack := make([]string,0,0)
	for _,v:=range s{
		ch := string(v)
		if ch=="("||ch=="{"||ch=="[" {
			stack = append(stack,ch)
		} else {
			if len(stack)==0||(ch==")"&&string(stack[len(stack)-1])!="(")||(ch=="}"&&string(stack[len(stack)-1])!="{")||(ch=="]"&&string(stack[len(stack)-1])!="[")  {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack)>0{
		return false
	}
	return true
}
