package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parseIds(line string) (int, int, error) {
	split_line_contents := strings.Split(line, "   ")
	id1, err := strconv.Atoi(split_line_contents[0])
	if err != nil {
		return -1, -1, err
	}
	id2, err := strconv.Atoi(split_line_contents[1])
	if err != nil {
		return -1, -1, err
	}
	return id1, id2, nil
}

// part two:
// create a frequency mapping of the right list
// then for each num in left, multiple left_current by the mapping, if found, from right
// if mapping isn't found, multiple it by 0

func main() {
	file, err := os.Open("input-pc.txt")
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left_ids := &IntHeap{}
	right_ids := &IntHeap{}
	heap.Init(left_ids)
	heap.Init(right_ids)

	// part 2
	right_freq := make(map[int]int)
	left_list := []int{}
	for scanner.Scan() {
		id_left, id_right, err := parseIds(scanner.Text())
		// for part 2, as I scan:
		// - create that right frequency mapping
		// - append left to a list
		if err != nil {
			log.Fatalf("error parsing line [%s]: %v", scanner.Text(), err)
		}
		heap.Push(left_ids, id_left)
		heap.Push(right_ids, id_right)

		// compute right freq
		_, exists := right_freq[id_right]
		if !exists {
			right_freq[id_right] = 1
		} else {
			right_freq[id_right] += 1
		}
		// create left list
		left_list = append(left_list, id_left)
	}

	// part 1
	fmt.Printf("Computing part 1\n")
	sum := 0
	for (left_ids.Len() > 0) && (right_ids.Len() > 0) {
		current_left := heap.Pop(left_ids).(int)
		current_right := heap.Pop(right_ids).(int)
		diff := current_left - current_right
		if diff < 0 {
			diff = -diff
		}
		sum += diff
		// fmt.Printf("%d <-> %d --> %d\n", current_left, current_right, diff)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in scanner: %v", err)
	}
	fmt.Printf("sum = %d\n", sum)
	fmt.Printf("Computing part 1 done\n")

	// part 2
	fmt.Printf("Computing part 2\n")
	// fmt.Printf("left_list: %v\n", left_list)
	// fmt.Printf("right_freq: %#v\n", right_freq)
	sum_2 := 0
	for _, val := range left_list {
		current_freq_in_right, exists := right_freq[val]
		if exists {
			sum_2 += val * current_freq_in_right
		} else {
			continue
		}
	}
	fmt.Printf("sum_2: %d\n", sum_2)
	fmt.Printf("Computing part 2 done\n")
}
