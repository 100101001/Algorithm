package main

import "math"

//
func maxProfit(prices []int) int {
	var res = math.MinInt64
	var minPrice = prices[0]
	for i := 0; i < len(prices); i++ {
		res = maxInt(res, prices[i]-minPrice)
		minPrice = minInt(prices[i], minPrice)
	}
	return res
}
