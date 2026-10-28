// Package coffeemachine places an office coffee machine, ported (and
// finished) from python/misc/coffee_machine.py.
//
// The office is a grid; office[r][c] is true where a developer sits. Pick
// the cell that minimizes the total Manhattan distance to every developer.
package coffeemachine

import "sort"

// Locate uses the fact that Manhattan distance splits into a row part and
// a column part, and each part is minimized at the median. O(R*C) to
// collect positions plus sorting. ok is false when there are no
// developers.
func Locate(office [][]bool) (row, col int, ok bool) {
	var rows, cols []int
	for r := range office {
		for c, dev := range office[r] {
			if dev {
				rows = append(rows, r)
				cols = append(cols, c)
			}
		}
	}
	if len(rows) == 0 {
		return 0, 0, false
	}
	// rows are collected in order already; cols are not.
	sort.Ints(cols)
	return rows[(len(rows)-1)/2], cols[(len(cols)-1)/2], true
}

// Cost is the total distance from (row, col) to every developer.
func Cost(office [][]bool, row, col int) int {
	total := 0
	for r := range office {
		for c, dev := range office[r] {
			if dev {
				total += abs(r-row) + abs(c-col)
			}
		}
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
