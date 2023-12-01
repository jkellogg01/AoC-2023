package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const digits string = "0123456789"

var digitsSpelled map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	calVals := make([]int, 0)
	for _, line := range strings.Split(string(data), "\n") {
		first := -1
		for i, char := range line {
			found := false
			if strings.Contains(digits, string(char)) {
				log.Println("first digit is numeric: ", string(char))
				first, err = strconv.Atoi(string(char))
				if err != nil {
					log.Fatal(err)
				}
				found = true
			}
			for k, v := range digitsSpelled {
				dex := strings.Index(line[i:], k)
				if dex == 0 {
					log.Println("last digit is spelled: ", k)
					first = v
					found = true
				}
			}
			if found {
				break
			}
		}
		if first == -1 {
			log.Fatal("found no digits on this line")
		}
		last := -1
		for i := len(line) - 1; i >= 0; i-- {
			found := false
			if strings.Contains(digits, string(line[i])) {
				log.Println("last digit is numeric: ", string(line[i]))
				last, err = strconv.Atoi(string(line[i]))
				if err != nil {
					log.Fatal(err)
				}
				found = true
			}
			for k, v := range digitsSpelled {
				dex := strings.Index(line[i:], k)
				if dex == 0 {
					log.Println("last digit is spelled: ", k)
					last = v
					found = true
				}
			}
			if found {
				break
			}
		}
		if last == -1 {
			log.Println("only one digit on this line: ", first)
			last = first
		}
		value := first*10 + last
		log.Println("value: ", value)
		calVals = append(calVals, value)
	}

	sum := 0
	for _, value := range calVals {
		sum += value
	}
	log.Println(sum)
}
