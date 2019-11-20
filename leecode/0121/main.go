package main

func maxProfit(prices []int) int {
	var res int
	if len(prices) == 0 {
		return 0
	}
	m := prices[0]
	for i := 0; i < len(prices); i++ {
		m = min(m, prices[i])
		if (prices[i] - m - res) > 0 {
			res = prices[i] - m
		}
	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
