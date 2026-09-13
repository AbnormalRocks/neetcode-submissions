func maxProfit(prices []int) int {
	min:=prices[0]
	max:=0
	for i,v:=range prices{
		if v<min{
			min=prices[i]
		}
		if v-min>max{
			max=v-min
		}
	}
	return max
}
