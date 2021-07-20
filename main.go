package main

import (
	"fmt"
	"math"
)

//Assume stock price is int
func get_max_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit = int(math.Max(float64(profit), float64(prices[i]-min)))
		} else {
			min = int(math.Min(float64(min), float64(prices[i])))
		}
	}
	return profit
}

func main() {
	prices := []int{2, 3, 5, 7, 11, 13}
	result := get_max_profit(prices)
	fmt.Println(result)
}
