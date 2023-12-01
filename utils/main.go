package main

import (
	"os"
	"strings"
)

func FileLines(filename string) ([]string, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(raw), "\n")
	return lines, nil
}
