package main

import (
	"fmt"
)

func main() {

	fmt.Println(isUgly(14))

	fmt.Println(isUgly(6))

	fmt.Println(isUgly(0))

	fmt.Println(isUgly(-2147483648))

}

func isUgly(num int) bool {
	if num==0{
		return false
	}
	//num= int(math.Abs(float64(num)))
	for num != 1 {
		fmt.Println(num)
		if num%2 == 0 {
			num = num / 2
		} else if num%3 == 0 {
			num = num / 3
		} else if num%5 == 0 {
			num = num / 5
		} else{
			return false
		}
	}
	return true
}
