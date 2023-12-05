package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
			maps[len(maps)-1].Append(line)
		}
	}

	wg := new(sync.WaitGroup)
	var seedLocations []int
	for i, seed := range seeds {
		wg.Add(1)
		log.Printf("displacing seed #%v (seed location %v)", i, seed)
		go func(seed int) {
			curr := seed
			for _, m := range maps {
				curr = m.Displace(curr)
			}
			seedLocations = append(seedLocations, curr)
			wg.Done()
		}(seed)
	}
	wg.Wait()

	min := -1
	for _, loc := range seedLocations {
		if loc < min || min < 0 {
			min = loc
		}
	}
	log.Printf("part one: minimum seed location %v\n", min)
}

func PartTwo(lines []string) {
	seedGlob, _ := strings.CutPrefix(lines[0], "seeds: ")
	var seeds []int
	log.Println(seedGlob)
	seedPairs := regexp.MustCompile(`(\d+) (\d+)`)
	for i, pair := range seedPairs.FindAllStringSubmatch(seedGlob, -1) {
		start, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err)
		}
		dist, err := strconv.Atoi(pair[2])
		if err != nil {
			log.Fatal(err)
		}
		end := start + dist
		log.Println(i, start, dist)
		for i := start; i < end; i++ {
			seeds = append(seeds, i)
		}
		// log.Println("finished appending seeds")
	}
	log.Println(len(seeds))

	var maps []Map
	numRow := regexp.MustCompile(`\d+ \d+ \d+`)
	for _, line := range lines[1:] {
		if line == "" {
			maps = append(maps, make(Map, 0))
			continue
		}
		if numRow.MatchString(line) {
			maps[len(maps)-1].Append(line)
		}
	}

	wg := new(sync.WaitGroup)
	var seedLocations []int
	for _, seed := range seeds {
		wg.Add(1)
		log.Printf("displacing seed at %v", seed)
		go func(seed int) {
			curr := seed
			for _, m := range maps {
				curr = m.Displace(curr)
			}
			seedLocations = append(seedLocations, curr)
			log.Printf("finished displacing seed %v", seed)
			wg.Done()
		}(seed)
	}
	wg.Wait()

	min := -1
	for _, loc := range seedLocations {
		if loc < min || min < 0 {
			min = loc
		}
	}
	log.Printf("part two: minimum seed location %v\n", min)
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

func (m *Map) Append(line string) {
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
	*m = append(*m, next)
}
