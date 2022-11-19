package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

type coord struct {
	c int // column
	r int // row
}

var grid = make(map[coord]int) // A global grid, containing the wires

func main() {
	lines, err := lfs.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	dist, crosses := buildGrid(lines)
	fmt.Printf("Dist: %d and a list of the crossings:\n%v", dist, crosses)

}

func buildGrid(wires []string) (distance int, crosses []coord) {
	for num, wire := range wires {
		plotWire(wire, num+1)
	}
	// fmt.Println(grid)
	crosses = getCrosses()
	for _, cross := range crosses {
		if distance == 0 { // It's the first, so it's going to be the shortest
			distance = manhattanDistance(cross)
		} else if manhattanDistance(cross) < distance { // This one is shorter than a previous shortest
			distance = manhattanDistance(cross)
		}

	}
	return distance, crosses
}

func plotWire(wire string, num int) {
	var pos = coord{0, 0}
	wiresteps := strings.Split(wire, ",")
	for _, wirestep := range wiresteps {
		steps := lfs.SilentAtoi(wirestep[1:]) // get the number from the step
		switch string(wirestep[0]) {          // Evaluate the instructin (U,D,L,R)
		case "U": // We will increase r this amount. Each step we will check for a crossing.
			for i := 0; i < steps; i++ {
				pos.r++
				if grid[pos] != num {
					grid[pos] += num
				}
			}
		case "D": // We will decrease r this amount. Each step we will check for a crossing.
			for i := 0; i < steps; i++ {
				pos.r--
				if grid[pos] != num {
					grid[pos] += num
				}
			}
		case "R": // We will increase c this amount. Each step we will check for a crossing.
			for i := 0; i < steps; i++ {
				pos.c++
				if grid[pos] != num {
					grid[pos] += num
				}
			}
		case "L": // We will decrease c this amount. Each step we will check for a crossing.
			for i := 0; i < steps; i++ {
				pos.c--
				if grid[pos] != num {
					grid[pos] += num
				}
			}
		default:
			panic(fmt.Sprintf("Unknown instuction: %s", wirestep[0]))
		}
	}
	return
}

func getCrosses() (crosses []coord) {
	for coord, n := range grid {
		if n > 2 {
			crosses = append(crosses, coord)
		}
	}
	return crosses
}

// manhattanDistance calculates the manhattan distance of a coordinate distance to 0,0
func manhattanDistance(input coord) int {
	return int(math.Abs(float64(input.c))) + int(math.Abs(float64(input.r))) // We use absolute values to ignore +/-
}
