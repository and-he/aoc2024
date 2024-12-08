package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./test2.txt")
	if err != nil {
		log.Fatalf("Error loading input file: %v\n", err)
	}
	fileContent := string(data)
	// fmt.Println(fileContent)

	fmt.Printf("part1: %v\n", part1(fileContent))
	fmt.Printf("part2: %v\n", part2(fileContent))
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
	// create mapping of
	//  indicies[0] -> matches
	//  so if I have a match of mul(x,x) at indices [1 9], I'd get the mapping 1:mul(x,x)
	//  no, it should match index to the token where token is one of the following:
	// 		- mul(x, y)
	// 		- do()
	// 		- don't()
	// get the keys of this mapping as an int array (indices)
	// indices.sort()
	// multiplier = 1
	// as you go through indices, if it is a do, make multiplier = 1
	// if it is a don't, make multiplier = 0
	// sum += multiplier * x * y

	mapping := make(map[int][]int)

	fmt.Printf("fileContent: [%v]\n", fileContent)
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	mul_matches := re.FindAllStringSubmatch(fileContent, -1)
	fmt.Printf("mul_matches = %v\n", mul_matches)
	mul_indices := re.FindAllStringSubmatchIndex(fileContent, -1)
	fmt.Printf("mul_indices: %v\n", mul_indices)
	for i, mul_index := range mul_indices {
		// mul_index_2 := []int{mul_index[0], mul_index[1]}
		mul_match_indicies := []int{}
		mul_match_start, err := strconv.Atoi(mul_matches[i][1])
		if err != nil {
			log.Fatalf("Error converting mul_matches[i][1]: %v\n", err)
		}
		mul_match_indicies = append(mul_match_indicies, mul_match_start)

		mul_match_end, err := strconv.Atoi(mul_matches[i][2])
		if err != nil {
			log.Fatalf("Error converting mul_matches[i][2]: %v\n", err)
		}
		mul_match_indicies = append(mul_match_indicies, mul_match_end)
		mapping[mul_index[0]] = mul_match_indicies
	}

	fmt.Printf("mapping: %#v\n", mapping)

	do_re := regexp.MustCompile("do\\(\\)")
	do_matches := do_re.FindAllStringSubmatch(fileContent, -1)
	fmt.Printf("do_matches: %v\n", do_matches)
	do_indices := do_re.FindAllStringSubmatchIndex(fileContent, -1)
	fmt.Printf("do_indices: %v\n", do_indices)

	dont_re := regexp.MustCompile("don't\\(\\)")
	dont_matches := dont_re.FindAllStringSubmatch(fileContent, -1)
	fmt.Printf("dont_matches: %v\n", dont_matches)
	dont_indices := dont_re.FindAllStringSubmatchIndex(fileContent, -1)
	fmt.Printf("dont_indices: %v\n", dont_indices)
	return 0
}
