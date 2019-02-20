package main

import "fmt"

//var graph [][]int

func main() {
	graph := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	var path = [5]int{-1, -1, -1, -1, -1}

	fmt.Print(graph)
	fmt.Print(path)
	hamCycle(graph,path)
}

func hamCycle(g [][]int, path [5]int) {
	path[0] = 0
	if hamCycleHelper(g, path, 1) == false {
		fmt.Print("No soultion exist")
	}
	//printHamPath(path)
}

func hamCycleHelper(graph [][]int, path [5]int, pos int) bool {
	// we are on the last vertex
	if pos == len(graph) {
		// check last vertex is adjacent to the first one or not
		if graph[path[pos-1]][path[0]] == 1 {
			return true
		} else {
			return false
		}
	}
	/*
	check if possible to add into hamiltonian cycle
	*/
	for v := 1; v < len(graph); v++ {
		if isSafe(v, graph, path, pos) == true {
			path[pos] = v
			if hamCycleHelper(graph, path, pos+1) == true {
				printHamPath(path)
				return true
			}
			path[pos] = -1
		}

	}
	return false

}

func isSafe(v int, graph [][]int, path [5]int, pos int) bool {
	/* Check if this vertex is an adjacent vertex of
			   the previously added vertex. */
	if graph[path[pos-1]][v] == 0 {
		return false
	}

	/* Check if the vertex has already been included.
	   This step can be optimized by creating an array
	   of size V */
	for i := 0; i < pos; i++ {
		if path[i] == v {
			return false
		}
	}
	return true
}

func printHamPath(path [5]int) {
	fmt.Print("path",path)
}
