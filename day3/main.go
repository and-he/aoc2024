package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error loading input file: %v\n", err)
	}
	fileContent := string(data)
	// fmt.Println(fileContent)

	fmt.Printf("part1: %v", part1(fileContent))
}

func part1(fileContent string) int {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(fileContent, -1)

	sum := 0
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("%v", err)
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("%v", err)
		}

		sum += num1 * num2
	}
	return sum
}

func part2(fileContent string) int {
	// use new regex in the form: [do]0+/[don't]0+(mul([0-9]+,[0-9]+))
	return 0
}
