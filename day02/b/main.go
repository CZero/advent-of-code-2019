// Day 2: A
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

func main() {
	code, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	for a := 0; a <= 99; a++ {
		for b := 0; b <= 99; b++ {
			answer := runIntcode(code[0], a, b)
			if answer == 19690720 {
				fmt.Printf("The answer: %d\nAchieved with %d\n", answer, 100*a+b)
				break
			}
		}
	}
}

// runIntcode is the coderunner
func runIntcode(code string, a, b int) int {
	rcode := resolveIntcode(code)
	rcode = preAdjust(rcode, a, b)
	executed := runSteps(rcode)
	// fmt.Println(code)
	// fmt.Println(executed)
	return executed[0]

}

// resolveIntcode turns the string into []int (1,2,3,4 -> []int{1,2,3,4})
func resolveIntcode(code string) (rescode []int) {
	exploded := strings.Split(code, ",")
	for _, num := range exploded {
		rescode = append(rescode, libaoc.SilentAtoi(num))
	}
	return rescode
}

// runSteps follows through the steps in the code, 1 is add, 2 is multi, 99 is end
func runSteps(input []int) []int {
	var pos int // Tracks where we are
	for {
		switch input[pos] {
		case 1: // Addition
			num1, num2 := input[input[pos+1]], input[input[pos+2]]
			input[input[pos+3]] = num1 + num2
			pos += 4
		case 2: // Multiplication
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
func preAdjust(input []int, a, b int) []int {
	input[1] = a
	input[2] = b
	return input
}
