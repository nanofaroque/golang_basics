package main

import "fmt"

func main() {
	data1 := [] int{7, 1, 5, 3, 6, 4}
	res1 := maxProfitSecond(data1)
	fmt.Print(res1)

	data2 := [] int{1, 2, 3, 4, 5}
	res2 := maxProfitSecond(data2)
	fmt.Print(res2)

	data3 := [] int{7,6,4,3,1}
	res3 := maxProfitSecond(data3)
	fmt.Print(res3)
}

func maxProfitSecond(prices []int) int {
	if len(prices)<2 {return 0}
	mProfit := 0
	buy := prices[0]
	var sale int

	for i := 1; i < len(prices); i++ {
		sale = prices[i]
		if sale < buy {
			buy = sale
		} else {
			if i+1 < len(prices) {
				if prices[i+1] < sale {
					mProfit += sale - buy
					buy = prices[i+1]
					i++
				}
			} else if sale > buy {
				mProfit += sale - buy
			}
		}
	}
	return mProfit
}
