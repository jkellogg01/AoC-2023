package main

import (
	"fmt"
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

var errFoundNoDigit error = fmt.Errorf("no digit at position")

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	calVals := make([]int, 0)
	for _, line := range strings.Split(string(data), "\n") {
		// log.Println(line)
		first := -1
		for i := range line {
			found, err := digitAtPosition(line, i)
			if err == errFoundNoDigit {
				continue
			} else if err != nil {
				log.Fatal(err)
			}
			log.Println("first digit:", found)
			first = found
			break
		}
		if first == -1 {
			log.Fatal("found no digits on this line")
		}
		last := -1
		for i := len(line) - 1; i >= 0; i-- {
			found, err := digitAtPosition(line, i)
			if err == errFoundNoDigit {
				continue
			} else if err != nil {
				log.Fatal(err)
			}
			log.Println("last digit:", found)
			last = found
			break
		}
		if last == -1 {
			// This shouldn't ever happen because the backwards loop will just capture the same digit that the forward loop does.
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

/*
2023/12/01 13:50:09 fivefour7nineseven1qtcdqbp1four
2023/12/01 13:50:09 first digit: 5
2023/12/01 13:50:09 last digit: 1
2023/12/01 13:50:09 value:  51
*/

func digitAtPosition(s string, idx int) (int, error) {
	if strings.Contains(digits, string(s[idx])) {
		return strconv.Atoi(string(s[idx]))
	}
	for k, v := range digitsSpelled {
		// log.Println(k, v)
		digIdx := strings.Index(s[idx:], k)
		if digIdx == 0 {
			return v, nil
		}
	}
	// log.Println(s[idx:])
	return -1, errFoundNoDigit
}
