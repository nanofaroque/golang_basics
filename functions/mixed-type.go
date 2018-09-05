package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	mixed := []interface{}{"foo", 10, primes }
	//fmt.Println(mixed...)

	test := map[string]interface{}{
		"omar":mixed,
	}

	fmt.Println(test["omar"])

}
