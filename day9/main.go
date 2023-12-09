package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	log.Print("start")
	log.Print(PartOne(lines))
	log.Print(PartTwo(lines))
}

func PartOne(lines []string) int {
	sum := 0
	for _, line := range lines {
		numsStrs := strings.Split(line, " ")
		values := make([]int, len(numsStrs))
		for i, str := range numsStrs {
			val, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			values[i] = val
		}
		sum += predict(values)
	}
	return sum
}

func predict(nums []int) int {
	allZeroes := true
	for _, num := range nums {
		if num != 0 {
			allZeroes = false
		}
	}
	if allZeroes {
		return 0
	}

	up := make([]int, len(nums)-1)
	for i := range up {
		up[i] = nums[i+1] - nums[i]
	}
	return nums[len(nums)-1] + predict(up)
}

func PartTwo(lines []string) int {
	sum := 0
	for _, line := range lines {
		numsStrs := strings.Split(line, " ")
		values := make([]int, len(numsStrs))
		for i, str := range numsStrs {
			val, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			values[i] = val
		}
		sum += postdict(values)
	}
	return sum
}

func postdict(nums []int) int {
	allZeroes := true
	for _, num := range nums {
		if num != 0 {
			allZeroes = false
		}
	}
	if allZeroes {
		return 0
	}

	up := make([]int, len(nums)-1)
	for i := range up {
		up[i] = nums[i+1] - nums[i]
	}
	return nums[0] - postdict(up)
}
