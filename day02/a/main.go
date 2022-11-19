// https://adventofcode.com/2019/day/2
// --- Day 2: 1202 Program Alarm ---

package main

import (
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	input, err := lfs.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	_ = InputToInts(input[0])
}

// InputToInts converts the input string to []ints
func InputToInts(input string) []int {
	var intcode []int
	numberstrs := strings.Fields(input)

	for _, n := range numberstrs {
		intcode = append(intcode, lfs.SilentAtoi(n))
	}
	return intcode
}

// IntProcessor processes the intcode program
func IntProcessor(program []int) []int {
	return []int{}
}
