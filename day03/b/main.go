package main

import (
	"fmt"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

type coord struct {
	r int // Row
	c int // Column
}

type wiresteps struct {
	wire1steps int // Steps for wire 1 to get here
	wire2steps int // Steps for wire 2 to get here
}

var grid = make(map[coord]wiresteps)

func main() {
	wires, err := lfs.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	buildgrid(wires)
}

func buildgrid(wires []string) {
	for n, wire := range wires {
		plotWire(wire, n+1)
	}
	// for pos, wiresteps := range grid {
	// 	fmt.Printf("Coord: %v, Wiresteps: %v\n", pos, wiresteps)
	// }
	crossing, dist := findLowestCross()
	fmt.Printf("Lowest distance: %d, found at: %v\n", dist, crossing)
}

func plotWire(wire string, n int) {
	var (
		stepcount int
		pos       coord
	)
	steps := strings.Split(wire, ",") // A range of steps, "," separated
	for _, step := range steps {
		repeats := lfs.SilentAtoi(step[1:])
		switch string(step[0]) {
		case "U": // r increases
			for i := 0; i < repeats; i++ {
				stepcount++
				pos.r++
				if n == 1 {
					if grid[pos].wire1steps == 0 {
						possteps := grid[pos]
						possteps.wire1steps = stepcount
						grid[pos] = possteps
					}
				} else {
					if grid[pos].wire2steps == 0 {
						possteps := grid[pos]
						possteps.wire2steps = stepcount
						grid[pos] = possteps
					}
				}
			}
		case "D": // r decreases
			for i := 0; i < repeats; i++ {
				stepcount++
				pos.r--
				if n == 1 {
					if grid[pos].wire1steps == 0 {
						possteps := grid[pos]
						possteps.wire1steps = stepcount
						grid[pos] = possteps
					}
				} else {
					if grid[pos].wire2steps == 0 {
						possteps := grid[pos]
						possteps.wire2steps = stepcount
						grid[pos] = possteps
					}
				}
			}
		case "R": // c increases
			for i := 0; i < repeats; i++ {
				stepcount++
				pos.c++
				if n == 1 {
					if grid[pos].wire1steps == 0 {
						possteps := grid[pos]
						possteps.wire1steps = stepcount
						grid[pos] = possteps
					}
				} else {
					if grid[pos].wire2steps == 0 {
						possteps := grid[pos]
						possteps.wire2steps = stepcount
						grid[pos] = possteps
					}
				}
			}
		case "L": // c decreases
			for i := 0; i < repeats; i++ {
				stepcount++
				pos.c--
				if n == 1 {
					if grid[pos].wire1steps == 0 {
						possteps := grid[pos]
						possteps.wire1steps = stepcount
						grid[pos] = possteps
					}
				} else {
					if grid[pos].wire2steps == 0 {
						possteps := grid[pos]
						possteps.wire2steps = stepcount
						grid[pos] = possteps
					}
				}
			}
		}

	}
}

func findLowestCross() (coord, int) {
	var lowest int     // Lowest combined steps
	var lowcoord coord // The crossing referenced
	for pos, wiresteps := range grid {
		if wiresteps.wire1steps > 0 && wiresteps.wire2steps > 0 {
			fmt.Println(wiresteps.wire1steps + wiresteps.wire2steps)
			if lowest != 0 && wiresteps.wire1steps+wiresteps.wire2steps < lowest {
				lowest = wiresteps.wire1steps + wiresteps.wire2steps
				lowcoord = pos
			}
			if lowest == 0 {
				lowest = wiresteps.wire1steps + wiresteps.wire2steps
			}

		}

	}
	return lowcoord, lowest
}
