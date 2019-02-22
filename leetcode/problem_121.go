package main

import (
	"fmt"
)

func main() {
	data1:=[] int{7,6,4,3,1}
	res:=maxProfit(data1)
	fmt.Print(res)

}

func maxProfit(prices []int) int {
	if len(prices)<2 {return 0}
	maxProfit:=0
	leftMin:=prices[0]
	for i:=1; i< len(prices);i++  {
		selPrice:=prices[i]
		profit:=selPrice-leftMin
		if profit>maxProfit{maxProfit=profit}
		if selPrice<leftMin{leftMin=selPrice}
	}
	return maxProfit

}