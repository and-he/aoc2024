package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func determineSafeReport(report []int) bool {
	// first check if mono decreasing
	// 		if abs difference > 3 or < 1, unsafe
	// 		if at end + mono dec, return safe
	// if not mono dec, check mono inc
	// 		if abs difference > 3 or < 1, unsafe
	// return true

	a, b := 0, 1
	for b < len(report) {
		// if adjacent levels differ by more than 3 or less than 1, unsafe
		diff := abs(report[a] - report[b])
		if (diff > 3) || (diff < 1) {
			// fmt.Printf("diff invalid on indices %d and %d on report %v", a, b, report)
			return false
		}
		// if not mono dec, break
		if report[a] <= report[b] {
			// fmt.Printf("[breaking...]")
			break
		}
		a += 1
		b += 1
	}
	// fmt.Printf("[breaking successful..., a=%d b=%d]", a, b)
	if (b == len(report)) && (report[a-1] > report[b-1]) {
		return true
	}

	// check mono inc
	a, b = 0, 1
	for b < len(report) {
		// if adjacent levels differ by more than 3 or less than 1, unsafe
		diff := abs(report[a] - report[b])
		if (diff > 3) || (diff < 1) {
			// fmt.Printf("diff invalid on indices %d and %d on report %v", a, b, report)
			return false
		}
		// if not mono inc, break
		if report[a] >= report[b] {
			// fmt.Printf("[breaking2...]")
			break
		}
		a += 1
		b += 1
	}
	// fmt.Printf("[breaking2 successful..., a=%d b=%d]", a, b)
	if (b == len(report)) && (report[a-1] < report[b-1]) {
		return true
	}
	return false
}

func removeAtIndex(report []int, index int) []int {
	// not sure why the below commented doesn't work, maybe to do with pointers?
	// return append(report[:index], report[index+1:]...)
	newReport := []int{}
	newReport = append(newReport, report[:index]...)
	newReport = append(newReport, report[index+1:]...)

	return newReport
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	num_reports_safe := 0

	for scanner.Scan() {
		split_line_contents := strings.Split(scanner.Text(), " ")

		report := []int{}
		for _, level_str := range split_line_contents {
			level, err := strconv.Atoi(level_str)
			if err != nil {
				log.Fatalf("Error converting level to int: [%s]\n", level_str)
			}
			report = append(report, level)

		}
		if determineSafeReport(report) {
			num_reports_safe += 1
			continue
		}
		// fmt.Printf("level: %v --> safe?=%v\n", report, determineSafeReport(report))
		// now operate on report
		for i := range len(report) {
			reportNew := removeAtIndex(report, i)
			// fmt.Printf("reportNew: %v\n", reportNew)
			if determineSafeReport(reportNew) {
				num_reports_safe += 1
				break
			}
		}
	}
	fmt.Printf("num_reports_safe: %d\n", num_reports_safe)
}
