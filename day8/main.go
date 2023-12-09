package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("start")
	log.Print(PartOne(raw))
	log.Print(PartTwo(raw))
}

func PartOne(data []byte) int {
	dirs, nodeData, _ := strings.Cut(string(data), "\n\n")
	nodeRow, err := regexp.Compile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	if err != nil {
		log.Fatal(err)
	}

	nodes := make(map[string][2]string)
	nodeRows := nodeRow.FindAllStringSubmatch(nodeData, -1)
	for _, row := range nodeRows {
		nodes[row[1]] = [2]string{row[2], row[3]}
	}

	return walkMap("AAA", "ZZZ", dirs, &nodes)
}

func PartTwo(data []byte) int {
	dirs, nodeData, _ := strings.Cut(string(data), "\n\n")
	nodeRow, err := regexp.Compile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	if err != nil {
		log.Fatal(err)
	}

	nodes := make(map[string][2]string)
	nodeRows := nodeRow.FindAllStringSubmatch(nodeData, -1)
	for _, row := range nodeRows {
		nodes[row[1]] = [2]string{row[2], row[3]}
	}

	// find all of the numbers of steps to get from each node
	// that ends with A to any node that ends in Z,
	// then find the least common multiple of those numbers
	stepCounts := make([]int, 0)
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			stepCounts = append(stepCounts, walkMap(k, `\w\wZ`, dirs, &nodes))
		}
	}

	return leastCommonMultiple(stepCounts)
}

func walkMap(start, end, dirs string, nodes *map[string][2]string) int {
	steps := 0
	curr := start
	finish := regexp.MustCompile(end)
	for {
		if finish.MatchString(curr) {
			return steps
		}
		idx := steps % len(dirs)
		switch dirs[idx] {
		case 'L':
			curr = (*nodes)[curr][0]
		case 'R':
			curr = (*nodes)[curr][1]
		}
		steps++
	}
}

func leastCommonMultiple(nums []int) int {
	fmt.Printf("finding least common multiple of:\n%v\n", nums)
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Printf("max: %v\n", max)
	for i := max; true; i += max {
		valid := true
		for _, num := range nums {
			if i%num != 0 {
				valid = false
			}
		}
		if valid {
			return i
		}
		// fmt.Printf("%v is not a common multiple\n", i)
	}
	return -1
}

// func lcmForTheHaters(nums []int) int {

// }
