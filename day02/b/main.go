// Day 2: A
package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	code, err := lfs.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	runIntcode(code[0])
}

// runIntcode is the coderunner
func runIntcode(code string) {
	rcode := resolveIntcode(code)
	rcode = preAdjust(rcode)
	executed, err := runSteps(rcode)
	// fmt.Println(code)
	// fmt.Println(executed)
	fmt.Printf("The answer: %d\n", executed[0])
}

// resolveIntcode turns the string into []int (1,2,3,4 -> []int{1,2,3,4})
func resolveIntcode(code string) (rescode []int) {
	exploded := strings.Split(code, ",")
	for _, num := range exploded {
		rescode = append(rescode, lfs.SilentAtoi(num))
	}
	return rescode
}

// runSteps follows through the steps in the code, 1 is add, 2 is multi, 99 is end
func runSteps(input []int) ([]int, error) {
	var pos int // Tracks where we are
	for {
		switch input[pos] {
		case 1: // Addition
			if input[pos+1] > len(input-1) || input[pos+2] > len(input-1) {
				return []int{}, errors.New("Out of bounds")
			}
			num1, num2 := input[input[pos+1]], input[input[pos+2]]
			input[input[pos+3]] = num1 + num2
			pos += 4
		case 2: // Multiplication
			if input[pos+1] > len(input-1) || input[pos+2] > len(input-1) {
				return []int{}, errors.New("Out of bounds")
			}
			num1, num2 := input[input[pos+1]], input[input[pos+2]]
			input[input[pos+3]] = num1 * num2
			pos += 4
		case 99: // End
			return input
		default: // We really shouldn't get here
			panic(fmt.Sprintf("Invalid code %d", input[pos]))
		}
	}
}

// preAdjust is an addition to restore the previous state of the real input! (it crashes the other examples!)
func preAdjust(input []int) []int {
	input[1] = 64
	input[2] = 72
	return input
}
