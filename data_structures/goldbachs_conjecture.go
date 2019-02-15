package main

import "fmt"

func main() {
	n := 20
	primes := generatePrimes(n)
	for k := range primes {
		if primes[n-k]==true{
			fmt.Print("First: ",k,"Second: ",n-k)
		}
	}
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
