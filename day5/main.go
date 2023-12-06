package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Offset struct {
	Dest   int
	Source int
	Length int
}

type Map []*Offset

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
	seedGlob, _ := strings.CutPrefix(lines[0], "seeds: ")
	var seeds []int
	for _, str := range strings.Split(seedGlob, " ") {
		seed, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seed)
	}

	var maps []Map
	numRow, err := regexp.Compile(`\d+ \d+ \d+`)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines[1:] {
		if line == "" {
			maps = append(maps, make(Map, 0))
			continue
		}
		if numRow.MatchString(line) {
			maps[len(maps)-1] = append(maps[len(maps)-1], newOffset(line))
		}
	}
	log.Printf("found %v maps", len(maps))

	min := -1
	for _, seed := range seeds {
		curr := seed
		for _, m := range maps {
			curr = m.Displace(curr)
		}
		// log.Printf("seed %v -> %v (offset %v)", seed, curr, curr-seed)
		if curr < min || min < 0 {
			min = curr
			log.Printf("new minimum: %v\n", min)
		}
	}

	log.Printf("part one: minimum seed location %v\n", min)
}

func PartTwo(lines []string) {
	seedGlob, _ := strings.CutPrefix(lines[0], "seeds: ")
	var seeds []struct {
		Start int
		End   int
	}
	seedPairs := regexp.MustCompile(`(\d+) (\d+)`)
	for _, pair := range seedPairs.FindAllStringSubmatch(seedGlob, -1) {
		start, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err)
		}
		dist, err := strconv.Atoi(pair[2])
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, struct {
			Start int
			End   int
		}{Start: start, End: start + dist})
	}

	var maps []Map
	numRow := regexp.MustCompile(`\d+ \d+ \d+`)
	for _, line := range lines[1:] {
		if line == "" {
			maps = append(maps, make(Map, 0))
			continue
		}
		if numRow.MatchString(line) {
			maps[len(maps)-1] = append(maps[len(maps)-1], newOffset(line))
		}
	}
	log.Printf("found %v maps", len(maps))

	min := -1
	// lastOffset := 0
	for _, sr := range seeds {
		log.Printf("parsing seed range: %v - %v", sr.Start, sr.End)
		for i := sr.Start; i <= sr.End; i++ {
			curr := i
			for _, m := range maps {
				curr = m.Displace(curr)
			}
			// if i == sr.Start || i == sr.End {
			// 	log.Printf("seed %v -> %v (offset %v)", i, curr, curr-i)
			// }
			// if lastOffset != curr-i {
			// 	log.Printf("OFFSET CHANGE: %v", curr-i)
			// 	log.Printf("seed %v -> %v (offset %v)", i, curr, curr-i)
			// 	lastOffset = curr - i
			// }
			if curr < min || min < 0 {
				min = curr
				log.Printf("new minimum: %v", min)
			}
		}
	}

	log.Printf("part two: minimum seed location %v", min)
}

func (m *Map) Displace(n int) int {
	for _, offset := range *m {
		if n < offset.Source || n > offset.Source+offset.Length {
			continue
		}
		distance := offset.Dest - offset.Source
		return n + distance
	}
	return n
}

func newOffset(line string) *Offset {
	// dest start, source start, range length
	next := new(Offset)
	var row []int
	for _, str := range strings.Split(line, " ") {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		row = append(row, num)
	}
	next.Dest = row[0]
	next.Source = row[1]
	next.Length = row[2]
	return next
}
