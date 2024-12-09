package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./input.txt")
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

func extractProductFromMulExpr(expr string, multiplier int) int {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(expr, -1)
	fmt.Printf("extractProductFromMulExpr extracted matches = %v\n", matches)
	firstDigitStr := matches[0][1]
	firstDigit, err := strconv.Atoi(firstDigitStr)
	if err != nil {
		log.Fatalf("extractProductFromMulExpr, failed to convert firstDigitStr: %v, err: %v\n", firstDigitStr, err)
	}
	secondDigitStr := matches[0][2]
	secondDigit, err := strconv.Atoi(secondDigitStr)
	if err != nil {
		log.Fatalf("extractProductFromMulExpr, failed to convert secondDigitStr: %v, err: %v\n", secondDigitStr, err)
	}
	return firstDigit * secondDigit * multiplier
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

	mapping := make(map[int]string)

	fmt.Printf("fileContent: [%v]\n", fileContent)
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	mul_matches := re.FindAllStringSubmatch(fileContent, -1)
	fmt.Printf("mul_matches = %v\n", mul_matches)
	mul_indices := re.FindAllStringSubmatchIndex(fileContent, -1)
	fmt.Printf("mul_indices: %v\n", mul_indices)
	for i, mul_index := range mul_indices {
		mul_match := mul_matches[i][0]
		mapping[mul_index[0]] = mul_match
	}

	fmt.Printf("mapping after getting mul_match: %#v\n", mapping)

	do_re := regexp.MustCompile("do\\(\\)")
	do_matches := do_re.FindAllStringSubmatch(fileContent, -1)
	fmt.Printf("do_matches: %v\n", do_matches)
	do_indices := do_re.FindAllStringSubmatchIndex(fileContent, -1)
	fmt.Printf("do_indices: %v\n", do_indices)
	for i, do_index := range do_indices {
		do_match := do_matches[i][0]
		mapping[do_index[0]] = do_match
	}
	fmt.Printf("mapping after getting do_match: %#v\n", mapping)

	dont_re := regexp.MustCompile("don't\\(\\)")
	dont_matches := dont_re.FindAllStringSubmatch(fileContent, -1)
	fmt.Printf("dont_matches: %v\n", dont_matches)
	dont_indices := dont_re.FindAllStringSubmatchIndex(fileContent, -1)
	fmt.Printf("dont_indices: %v\n", dont_indices)
	for i, dont_index := range dont_indices {
		dont_match := dont_matches[i][0]
		mapping[dont_index[0]] = dont_match
	}

	fmt.Printf("mapping after getting dont_match: %#v\n", mapping)

	keys := []int{}
	for key := range mapping {
		keys = append(keys, key)
	}

	fmt.Printf("Keys are: %#v\n", keys)

	sort.Ints(keys)
	fmt.Printf("Sorted indices keys: %#v\n", keys)

	multiplier := 1
	sum := 0
	for _, key := range keys {
		token := mapping[key]
		if token == "do()" {
			fmt.Printf("enabled at key = [%v]\n", key)
			multiplier = 1
		} else if token == "don't()" {
			fmt.Printf("disabled at key = [%v]\n", key)
			multiplier = 0
		} else { // token is a mul
			sum += extractProductFromMulExpr(token, multiplier)
		}
		fmt.Printf("mapping[key] = %v\n", mapping[key])
	}

	return sum
}
