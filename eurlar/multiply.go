package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Print(findSum(100))
}

func findSum(n int) int {
	total := 0;
	i := 3;
	for i < n {
		total = total + i;
		i = i + 3;
	}
	j := 5;
	for j < n {
		if (j%3 != 0) {
			total = total + j;
		}

		j = j + 5;
	}
	return total;
}
