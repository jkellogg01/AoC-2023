package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const digits string = "0123456789"

type Number struct {
	Near   [3][]rune
	Val    int
	Len    int
	Center [2]int
}

type Gear struct {
	Parts    []int
	Location [2]int
}

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	log.Println(PartOne(lines))
	log.Println(PartTwo(lines))
}

func PartOne(lines []string) int {
	sum := 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if strings.Contains(digits, string(line[j])) {
				num := newNumber(i, j, lines)
				j += num.Len
				if num.isPartNumber() {
					// log.Printf("%v is a part number", num.Val)
					sum += num.Val
				}
			}
		}
	}
	return sum
}

func PartTwo(lines []string) int {
	partNumbers := make([]*Number, 0)
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if strings.Contains(digits, string(line[j])) {
				num := newNumber(i, j, lines)
				j += num.Len
				if num.isPartNumber() {
					partNumbers = append(partNumbers, num)
				}
			}
		}
	}
	gears := make([]*Gear, 0)
	for _, num := range partNumbers {
		gear := num.hasGear()
		if gear != [2]int{-1, -1} {
			linkGear(gear[0], gear[1], &gears, num.Val)
			// log.Println(newGear)
		}
	}
	sum := 0
	for _, gear := range gears {
		// log.Println(gear.Parts)
		if len(gear.Parts) == 2 {
			// log.Println("valid gear")
			sum += gear.Ratio()
		}
	}
	return sum
}

func newNumber(cRow int, sCol int, lines []string) *Number {
	result := new(Number)
	result.Center = [2]int{
		max(sCol-1, 0),
		max(cRow-1, 0),
	}
	numLength := 0
	for i := sCol; i < len(lines[cRow]); i++ {
		if !strings.Contains(digits, string(lines[cRow][i])) {
			break
		}
		numLength++
	}
	result.Len = numLength
	value, err := strconv.Atoi(lines[cRow][sCol : sCol+numLength])
	if err != nil {
		log.Fatal(err)
	}
	result.Val = value
	near := new([3][]rune)
	for i, row := range lines[max(cRow-1, 0):min(cRow+2, len(lines))] {
		for _, char := range row[max(sCol-1, 0):min(sCol+numLength+1, len(row))] {
			near[i] = append(near[i], char)
		}
	}
	result.Near = *near
	// log.Printf( /*"\nNear: %v\n%v\n%v*/ "\nVal: %v\nLen: %v\nCenter: %v\n" /*result.Near[0], result.Near[1], result.Near[2],*/, result.Val, result.Len, result.Center)
	return result
}

func (num *Number) isPartNumber() bool {
	for _, row := range num.Near {
		for _, char := range row {
			if char != '.' && !strings.ContainsRune(digits, char) {
				return true
			}
		}
	}
	return false
}

func (num *Number) hasGear() [2]int {
	for i, row := range num.Near {
		for j, char := range row {
			if char == '*' {
				return [2]int{num.Center[0] + j, num.Center[1] + i}
			}
		}
	}
	return [2]int{-1, -1}
}

func linkGear(x, y int, all *[]*Gear, part int) *Gear {
	for _, gear := range *all {

		if gear.Location == [2]int{x, y} {
			gear.Parts = append(gear.Parts, part)
			return gear
		}
	}
	return newGear(x, y, all, part)
}

func newGear(x, y int, all *[]*Gear, part int) *Gear {
	result := &Gear{
		Location: [2]int{
			x,
			y,
		},
		Parts: make([]int, 0),
	}
	result.Parts = append(result.Parts, part)
	*all = append(*all, result)
	return result
}

func (g *Gear) Ratio() int {
	ratio := 1
	for _, part := range g.Parts {
		ratio *= part
	}
	return ratio
}
