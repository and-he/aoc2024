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

func part1_d(graph [][]byte) int {
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

func searchInDirection(graph [][]byte, r int, c int, direction []int, goal []byte) bool {
	// goal is to find 'M', 'A', 'S' when doing: r + direction[0], c + direction[1] three times
	// if, when going in a direction we are out of bounds, return false
	// if, when collecting a new character, if it isn't currently what we want next, return false
	for _, char := range goal {
		r += direction[0]
		c += direction[1]
		if r < 0 || r >= len(graph) || c < 0 || c >= len(graph[0]) || graph[r][c] != char {
			return false
		}
	}
	return true
}

func part1(graph [][]byte) int {
	count := 0
	fmt.Printf("%v\n", graph)
	// go line by line, char by char
	// if current char is an x, look straight in one of 8 directions (within bounds)
	// if a search in one direction 3 letters deep matches "MAS", return true for that direction
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	goal := []byte{'M', 'A', 'S'}
	for r, line := range graph {
		for c := range line {
			if graph[r][c] == 'X' {
				for _, direction := range directions {
					if searchInDirection(graph, r, c, direction, goal) {
						count += 1
					}
				}
			}
		}
	}
	return count
}

func main() {
	// convert file to 2d array of strs or chars?
	file, err := os.Open("./input.txt")
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
