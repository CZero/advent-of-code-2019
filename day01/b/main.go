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

// CalcSumFuel calculates the sum of fuel needed for all modules
func CalcSumFuel(input []string) int {
	var sum int // Total fuel needed for the modules and their fuel
	for _, line := range input {
		sum += CalcFuelNeed(lfs.SilentAtoi(line))
	}
	return sum
}

// CalcFuelNeed calculates the fuel needed for a module, recursively calculating the fuel needed for the fuel too.
func CalcFuelNeed(mass int) int {
	need := mass/3 - 2
	if need > 0 {
		return need + CalcFuelNeed(need)
	} else {
		return 0
	}
}
