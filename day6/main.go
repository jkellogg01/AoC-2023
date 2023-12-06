package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	PartOne(lines)
	PartTwo(lines)
}

func PartOne(lines []string) {
	num := regexp.MustCompile(`\d+`)
	timeStrs := num.FindAllString(lines[0], -1)
	var times []int
	for _, str := range timeStrs {
		val, _ := strconv.Atoi(str)
		times = append(times, val)
	}

	distStrs := num.FindAllString(lines[1], -1)
	var dists []int
	for _, str := range distStrs {
		val, _ := strconv.Atoi(str)
		dists = append(dists, val)
	}

	// if total race time = t, speed = t - runtime and vice versa
	// so speed + runtime = t && speed * runtime > distance for any winning scenario
	// answer is # of winning scenarios for each race multiplied all together
	result := 1
	for i := range times {
		time := times[i]
		dist := dists[i]
		result *= numWins(time, dist)
	}
	log.Printf("Part One: %v", result)
}

func PartTwo(lines []string) {
	num := regexp.MustCompile(`\d+`)
	timeStr := strings.Join(num.FindAllString(lines[0], -1), "")
	time, _ := strconv.Atoi(timeStr)
	distStr := strings.Join(num.FindAllString(lines[1], -1), "")
	dist, _ := strconv.Atoi(distStr)

	result := numWins(time, dist)
	log.Printf("Part Two: %v", result)
}

func numWins(t, d int) int {
	var result int
	for i := 1; i < t; i++ {
		dist := i * (t - i)
		if dist > d {
			result++
		}
	}
	return result
}
