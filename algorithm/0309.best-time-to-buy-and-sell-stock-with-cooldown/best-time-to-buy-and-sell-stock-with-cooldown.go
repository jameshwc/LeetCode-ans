func maxProfit(prices []int) int {
	sell := make([]int, len(prices)) // ith day and don't have stock
	buy := make([]int, len(prices))  // ith day and have stock
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if len(prices) <= 1 {
		return 0
	}
	buy[0] = -prices[0]
	buy[1] = max(buy[0], -prices[1])
	sell[1] = max(0, buy[0]+prices[1])
	for i := 2; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
	}
	return max(sell[len(prices)-1], buy[len(prices)-1])
}