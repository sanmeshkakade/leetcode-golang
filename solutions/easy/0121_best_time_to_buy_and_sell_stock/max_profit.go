/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

package maxprofit

// Time Complexity: O(N)
// Space Complexity: O(1)

func maxProfit(prices []int) int {

	n := len(prices)

	if n == 0 || n == 1 {
		return 0
	}

	minPrice := prices[0]
	res := 0

	for _, price := range prices {

		if price-minPrice > res {
			res = price - minPrice
		}

		if price < minPrice {
			minPrice = price
		}
	}

	return res
}

// Time Complexity: O(N^2)
// Space Complexity: O(1)

func maxProfit1(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > res {
				res = prices[i] - prices[j]
			}
		}
	}

	return res
}
