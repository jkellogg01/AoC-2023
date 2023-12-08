package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type HandBid struct {
	Hand string
	Type int
	Bid  int
}

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	PartOne(lines)
	PartTwo(lines)
}

const cards string = "AKQJT98765432"
const cards2 string = "AKQT98765432J"

func PartOne(lines []string) {
	var hands []HandBid
	for _, line := range lines {
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		hType := handType(split[0])
		if hType == -1 {
			log.Fatalf("invalid hand type: %s", split[0])
		}
		hands = append(hands, HandBid{
			Hand: split[0],
			Type: hType,
			Bid:  bid,
		})
	}
	sortHands(&hands)
	sum := 0
	for i, hand := range hands {
		fmt.Printf("hand %v type %v rank %v bid %v\n", hand.Hand, hand.Type, i+1, hand.Bid)
		sum += hand.Bid * (i + 1)
	}
	log.Println("Part one: ", sum)
}

func PartTwo(lines []string) {
	var hands []HandBid
	for _, line := range lines {
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		hType := handType2(split[0])
		if hType == -1 {
			log.Fatalf("invalid hand type: %s", split[0])
		}
		hands = append(hands, HandBid{
			Hand: split[0],
			Type: hType,
			Bid:  bid,
		})
	}
	sortHands2(&hands)
	sum := 0
	for i, hand := range hands {
		fmt.Printf("hand %v type %v rank %v bid %v\n", hand.Hand, hand.Type, i+1, hand.Bid)
		sum += hand.Bid * (i + 1)
	}
	log.Println("Part Two: ", sum)
}

func handType(s string) int {
	var (
		counts []int
		max    int
	)
	for _, char := range cards {
		count := strings.Count(s, string(char))
		if count > 0 {
			counts = append(counts, count)
		}
		if count > max {
			max = count
		}
	}
	// log.Printf("hand %v has max instances %v\nall instances: %v", s, max, counts)
	switch max {
	case 1:
		return 0
	case 2: // diff between 1 pair and 2 pair
		pairs := 0
		for _, count := range counts {
			if count == 2 {
				pairs++
			}
		}
		return pairs
	case 3: // diff between trips and fh
		for _, count := range counts {
			if count == 2 {
				return 4
			}
		}
		return 3
	case 4:
		return 5
	case 5:
		return 6
	}
	return -1
}

func handType2(s string) int {
	var (
		counts []int
		max    int
	)

	// log.Println(cards2[:len(cards2)-1])
	for _, char := range cards2[:len(cards2)-1] {
		re := regexp.MustCompile(fmt.Sprintf("%c|J", char))
		count := len(re.FindAllString(s, -1))
		if count > 0 {
			counts = append(counts, count)
		}
		if count > max {
			max = count
		}
	}
	// log.Printf("hand %v has max instances %v\nall instances: %v", s, max, counts)
	switch max {
	case 1:
		return 0
	case 2: // diff between 1 pair and 2 pair
		pairs := 0
		for _, count := range counts {
			if count == 2 {
				pairs++
			}
		}
		return pairs
	case 3: // diff between trips and fh
		for _, count := range counts {
			if count == 2 {
				return 4
			}
		}
		return 3
	case 4:
		return 5
	case 5:
		return 6
	}
	return -1
}

func sortHands(h *[]HandBid) {
	// log.Println(*h)
	length := len(*h)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-(i+1); j++ {
			if handsCompare((*h)[j], (*h)[j+1]) {
				// log.Println("swapping")
				// log.Println(*h)
				tmp := (*h)[j]
				(*h)[j] = (*h)[j+1]
				(*h)[j+1] = tmp
				// log.Println(*h)
			}
		}
	}
	// log.Println(*h)
}

func sortHands2(h *[]HandBid) {
	// log.Println(*h)
	length := len(*h)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-(i+1); j++ {
			if handsCompare2((*h)[j], (*h)[j+1]) {
				// log.Println("swapping")
				// log.Println(*h)
				tmp := (*h)[j]
				(*h)[j] = (*h)[j+1]
				(*h)[j+1] = tmp
				// log.Println(*h)
			}
		}
	}
	// log.Println(*h)
}

func handsCompare(a, b HandBid) bool {
	// log.Println(a, b)
	if a.Type != b.Type {
		if a.Type > b.Type {
			// log.Printf("swapped by hand types: %v > %v", a.Type, b.Type)
			return true
		} else {
			return false
		}
	}
	for i := range a.Hand {
		aDex := strings.IndexByte(cards, a.Hand[i])
		bDex := strings.IndexByte(cards, b.Hand[i])
		if aDex == bDex {
			continue
		} else if aDex < bDex {
			return true
		} else if bDex < aDex {
			return false
		}
	}
	// log.Printf("could not compare hands %s and %s", a.Hand, b.Hand)
	return false
}

func handsCompare2(a, b HandBid) bool {
	// log.Println(a, b)
	if a.Type != b.Type {
		if a.Type > b.Type {
			// log.Printf("swapped by hand types: %v > %v", a.Type, b.Type)
			return true
		} else {
			return false
		}
	}
	for i := range a.Hand {
		aDex := strings.IndexByte(cards2, a.Hand[i])
		bDex := strings.IndexByte(cards2, b.Hand[i])
		if aDex == bDex {
			continue
		} else if aDex < bDex {
			return true
		} else if bDex < aDex {
			return false
		}
	}
	// log.Printf("could not compare hands %s and %s", a.Hand, b.Hand)
	return false
}
