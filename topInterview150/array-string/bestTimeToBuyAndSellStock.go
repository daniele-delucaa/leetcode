package main

import "math"

func maxProfit(prices []int) int {
	maxIndex, minIndex := 0, 0
	max := 0
	min := math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			minIndex = i
		}
	}
	for i := minIndex; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
			maxIndex = i
		}
	}
	if maxIndex == 0 {
		return 0
	}
	return max - min
}
