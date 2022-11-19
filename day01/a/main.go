// https://adventofcode.com/2019/day/1
// --- Day 1: The Tyranny of the Rocket Equation ---

package main

import (
	"fmt"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	input, err := lfs.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fuel needed: %d\n", CalcSumFuel(input))
}

func CalcSumFuel(input []string) int {
	sum := 0
	for _, line := range input {
		sum += CalcModuleFuel(line)
	}
	return sum
}

func CalcModuleFuel(mass string) int {
	m := lfs.SilentAtoi(mass)
	return m/3 - 2
}
