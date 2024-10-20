package bestbuysellstock

func maxProfit1(prices []int) int {
	profit, minPrice := 0, int(1e9)

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if profit < prices[i]-minPrice {
				profit = prices[i] - minPrice
			}
		}
	}

	return profit
}

func maxProfit(prices []int) int {
	profit, mini := 0, int(1e9)

	for i := 0; i < len(prices); i++ {
		profit = max(profit, prices[i]-mini)
		mini = min(mini, prices[i])
	}
	return profit
}
