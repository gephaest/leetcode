package p122_buy_sell_stock

func maxProfit(prices []int) int {
	sum := 0

	for i := 0; i < len(prices); i++ {
		if i+1 >= len(prices) {
			break
		}

		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}

	return sum
}
