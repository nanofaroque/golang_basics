package main

import "fmt"

func main() {
	n := 20
	fmt.Print(n)

	primes := generatePrimes(n)
	fmt.Println(primes)
}

func generatePrimes(n int) map[int]bool {
	var arr = make([]bool, n)
	m := make(map[int]bool)
	for i := 2; i < n; i++ {
		temp := i
		count := 2
		if arr[i] == false {
			m[i] = true
			for temp < n {
				arr[temp] = true
				temp = i * count
				count += 1
			}
		}
	}
	return m
}
