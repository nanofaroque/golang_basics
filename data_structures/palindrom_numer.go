package main

import "fmt"

func main() {
	isPalindrome(121)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var rem = x
	var new = 0
	for rem/10 != 0 {
		var t = rem % 10
		rem = rem / 10
		fmt.Println("%d, %d",rem,t)
		new = new*10 + t
	}
	new=new*10+rem

	if new!=x{
		return false
	}

	return true
}
