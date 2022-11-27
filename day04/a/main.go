// Secure container https://adventofcode.com/2019/day/4
package main

import (
	"a/libaoc"
	"fmt"
	"strconv"
)

func main() {
	var (
		lower int = 156218 // Start range
		upper int = 652527 // End range
	)
	_, _ = lower, upper
	codes := findCodes(lower, upper)
	fmt.Printf("Possible codes:\n%v\n\nNumber: %d\n", codes, len(codes))
}

func findCodes(lower, upper int) []int {
	var codes []int
	for n := lower; n <= upper; {
		if validate(n) {
			codes = append(codes, n)
		}
		n = nextIncrement(n)
	}
	return codes
}

// nextIncrement gets the next number. It's just an incrementer, but it cheats up on the decrease rule
func nextIncrement(input int) (result int) {
	for {
		input++
		numstring := strconv.Itoa(input)
		if string(numstring[len(numstring)-1]) == "0" {
			return findFirst(input)
		}
		return input
	}
}

// findFirst finds the first possible combination after the number given
func findFirst(start int) int {
	number := numtostringslice(start)
	rolling := 0 // Indicates if we had a decrease, and contains the higher number
	for i := 1; i < len(number); i++ {
		if rolling != 0 { // We haven't run into a decrease yet
			number[i] = strconv.Itoa(rolling)
		} else {
			if libaoc.SilentAtoi(number[i-1]) > libaoc.SilentAtoi(number[i]) {
				number[i] = number[i-1]
				rolling = libaoc.SilentAtoi(number[i])
			}
		}

	}
	return stringslicetonum(number)
}

// validate checks if num is a possible code: Never decreases and has at least 1 pair of the same adjacent digits
func validate(input int) bool {
	number := numtostringslice(input)
	pairs := 0
	for i := 1; i < len(number); i++ {
		if libaoc.SilentAtoi(number[i-1]) > libaoc.SilentAtoi(number[i]) {
			return false
		}
		if libaoc.SilentAtoi(number[i-1]) == libaoc.SilentAtoi(number[i]) {
			pairs++
		}
	}
	if pairs == 0 {
		return false
	}
	return true
}

// numtostringslice turns a number into []string
func numtostringslice(num int) (result []string) {
	numstring := strconv.Itoa(num)
	for i := 0; i < len(numstring); i++ {
		result = append(result, string(numstring[i]))
	}
	return result
}

// stringslicetonum turns []string into int
func stringslicetonum(input []string) int {
	var built string
	for _, char := range input {
		built += char
	}
	return libaoc.SilentAtoi(built)
}
