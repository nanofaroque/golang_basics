package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(scanner.Text())
	nIn, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < nIn; i++ {
		n, _ := strconv.Atoi(scanner.Text())
		out := findSum(n)
		fmt.Println(out)
	}
	start := time.Now()
	fmt.Print(findSum(100))
	elapsed := time.Since(start)
	fmt.Println("elapsed time: ",elapsed)
}

func findSum(n int) int {
	total := 0;
	i := 3;
	for i < n {
		total = total + i;
		i = i + 3;
	}
	i = 5;
	for i < n {
		if i%3 != 0 {
			total = total + i;
		}

		i = i + 5;
	}
	return total;
}
