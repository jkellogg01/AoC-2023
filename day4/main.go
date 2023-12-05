package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Num  int
	Win  []int
	Pool []int
}

type ManyCard struct {
	Card   *Card
	Amount int
}

func main() {
	raw, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")

	PartOne(lines)
	PartTwo(lines)
}

func PartOne(lines []string) {
	sum := 0
	for i, line := range lines {
		card := newCard(line, i+1)
		wins := card.Wins()
		// log.Println(card, wins)
		if wins > 0 {
			points := math.Pow(2, float64(wins-1))
			sum += int(points)
		}
	}
	log.Println(sum)
}

func PartTwo(lines []string) {
	cards := make([]ManyCard, 0)
	for i, line := range lines {
		card := newCard(line, i+1)
		many := ManyCard{
			Card:   card,
			Amount: 1,
		}
		cards = append(cards, many)
	}
	var sum uint64 = 0
	for i, card := range cards {
		wins := card.Card.Wins()
		// log.Printf("card %v wins %v times with %v copies\n", card.Card.Num, wins, card.Amount)
		if wins > 10 {
			log.Println("something bad has happened.", card)
		}
		for j := range cards[i:min(i+wins, len(cards))] {
			cards[i+j+1].Amount += card.Amount
			// // log.Println(i + j + 2)
			// // log.Println(inCard)
		}
		sum += 1 + uint64(wins*card.Amount)

		// log.Printf("added %v, new sum %v\n", wins*card.Amount, sum)
	}
	log.Println(sum)
}

func (c *Card) Wins() int {
	wins := 0
	for _, win := range c.Win {
		for _, num := range c.Pool {
			if num == win {
				wins++
				break
			}
		}
	}
	return wins
}

func newCard(line string, num int) *Card {
	// // log.Println(line)
	_, nums, _ := strings.Cut(line, ": ")
	// // log.Println(nums)
	win, pool, _ := strings.Cut(nums, " | ")
	// // log.Println(win, pool)
	result := new(Card)
	result.Num = num
	for _, num := range strings.Split(win, " ") {
		// // log.Println(num)
		if num == "" {
			continue
		}
		val, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		result.Win = append(result.Win, val)
	}
	for _, num := range strings.Split(pool, " ") {
		// // log.Println(num)
		if num == "" {
			continue
		}
		val, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		result.Pool = append(result.Pool, val)
	}
	return result
}
