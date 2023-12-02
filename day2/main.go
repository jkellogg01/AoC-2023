package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := fileLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	idList := partOne(lines)
	powerList := partTwo(lines)

	idSum, powerSum := 0, 0
	for _, id := range idList {
		idSum += id
	}
	for _, power := range powerList {
		powerSum += power
	}
	log.Println(idSum, powerSum)
}

func partOne(lines []string) []int {
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var idList []int
	for i, line := range lines {
		gameID := i + 1
		_, gameData, _ := strings.Cut(line, ": ")
		outcomes := strings.Split(gameData, "; ")
		possible := true
		for _, outcome := range outcomes {
			if !validOutcome(rules, outcome) {
				log.Printf("game %d is impossible", gameID)
				possible = false
				break
			}
		}
		if possible {
			idList = append(idList, gameID)
		}
	}

	return idList
}

func partTwo(lines []string) []int {
	var powerList []int

	for _, line := range lines {
		_, game, _ := strings.Cut(line, ": ")
		mins := minCubes(game)

		power := 1
		for _, v := range mins {
			power *= v
		}
		powerList = append(powerList, power)
	}

	return powerList
}

func fileLines(filename string) ([]string, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(raw), "\n")
	return lines, nil
}

func validOutcome(rules map[string]int, outcome string) bool {
	for _, group := range strings.Split(outcome, ", ") {
		parts := strings.Split(group, " ")
		// log.Println(parts)
		amt, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range rules {
			if strings.Contains(parts[1], k) {
				if v < amt {
					return false
				}
				break
			}
		}
	}
	return true
}

func minCubes(game string) map[string]int {
	mins := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, outcome := range strings.Split(game, "; ") {
		for _, group := range strings.Split(outcome, ", ") {
			parts := strings.Split(group, " ")
			amt, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			for k, v := range mins {
				if strings.Contains(parts[1], k) {
					mins[k] = max(amt, v)
				}
			}
		}
	}
	return mins
}
