package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func dfs(graph [][]byte, target string, r int, c int) bool {
	// handle base case (when we found the word)
	if len(target) == 0 {
		return true
	}

	// handle invalid cases:
	// - out of bounds
	// - word we are at doesn't match what we want
	if r < 0 || r >= len(graph) || c < 0 || c >= len(graph[0]) || graph[r][c] != target[0] {
		return false
	}

	tempVisited := graph[r][c] // to undo the decision
	graph[r][c] = '.'          // make the decision to visit
	found := dfs(graph, target[1:], r+1, c) || dfs(graph, target[1:], r-1, c) || dfs(graph, target[1:], r, c+1) || dfs(graph, target[1:], r, c-1) ||
		dfs(graph, target[1:], r+1, c+1) || dfs(graph, target[1:], r-1, c-1) || dfs(graph, target[1:], r+1, c-1) || dfs(graph, target[1:], r-1, c+1)
	graph[r][c] = tempVisited // undo the decision
	return found
}

func part1(graph [][]byte) int {
	count := 0
	fmt.Printf("%v\n", graph)
	// go line by line and character by character
	// look for 'X' since that is your starting point
	// visit each neighbor recursively and at each collection, if I am collecting a character I want, continue visiting
	// if I collected a character that isn't what I want, break out of the visit path
	for r, line := range graph {
		for c := range line {
			found := dfs(graph, "XMAS", r, c)
			if found {
				count += 1
				fmt.Printf("DFS found at [%v][%v]\n", r, c)
			}
		}
	}
	return count
}

func main() {
	// convert file to 2d array of strs or chars?
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := [][]byte{}

	for scanner.Scan() {
		line := scanner.Text()
		line_parsed := []byte{}
		for i := range line {
			ch := line[i]
			line_parsed = append(line_parsed, ch)
		}
		graph = append(graph, []byte(line_parsed))
	}
	fmt.Printf("part1: %v\n", part1(graph))
}
