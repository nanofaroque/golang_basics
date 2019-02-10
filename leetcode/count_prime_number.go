package main

import "fmt"

func main() {
	result := countPrimes(10)
	fmt.Println()
	fmt.Print(result)
}
func countPrimes(n int) int {
	if n == 2 {
		return 0
	}

	visited := make([]bool, n+1)
	fmt.Println(visited)
	total := 0
	for i := 2; i <= n; i++ {
		if !visited[i] {
			total++
			//fmt.Print(i,"->")
			temp := i
			mul := 1
			for temp <= n {
				visited[temp] = true
				mul += 1
				temp = i * mul

			}
			//fmt.Println(visited)
		}
	}

	return total
}
