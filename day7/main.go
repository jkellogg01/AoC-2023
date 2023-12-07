package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	PartOne(lines)
	PartTwo(lines)
}

func PartOne(lines []string) {}

func PartTwo(lines []string) {}
