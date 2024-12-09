package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(graph [][]string) int {
	count := 0
	fmt.Printf("%v\n", graph)
	// go line by line and character by character
	// look for 'X' since that is your starting point
	// visit each neighbor recursively and at each collection, if I am collecting a character I want, continue visiting
	// if I collected a character that isn't what I want, break out of the visit path
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

	graph := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		line_parsed := []string{}
		for i := range line {
			ch := line[i]
			line_parsed = append(line_parsed, string(ch))
		}
		graph = append(graph, line_parsed)
	}
	part1(graph)
}
