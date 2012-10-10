// Missionaries and Cannibals.
//
// 3 Missionaries and 3 Cannibals and a boat are on one side of a
// river.  The boat can cross with 1 or 2 people.  The Cannibals will
// eat the Missionaries if they outnumber them on land or on the boat.
// Get everyone across the river.
//
// Throughout, state is represented by 0xMCB where M is the number of
// missionaries on the wrong side, C is the number of cannibals on the
// wrong side, and B is the number of boats on the wrong side.

package main

import "fmt"

// table of legal boat crossings
var δ = [2][5]int{
	{+0x011, +0x101, +0x111, +0x201, +0x021}, // when boat on correct side
	{-0x011, -0x101, -0x111, -0x201, -0x021}, // when boat on wrong side
}

const OVERFLOW_MASK = 0xCCE

func main() {
	seen := map[int]bool{}
	var solve func(int, []int) // Predeclare solve() so it can recur.
	solve = func(state int, steps []int) {
		// If nothing is on the wrong side, report result and stop searching.
		if state == 0 {
			fmt.Printf("%03x\n", steps)
			return
		}
		// Recur if state is valid.
		m, c, b := state>>8&0x3, state>>4&0x3, state&1
		cannibalsEat := (m != 0 && m < c) || (3-m != 0 && 3-m < 3-c)
		if state&OVERFLOW_MASK == 0 && !seen[state] && !cannibalsEat {
			seen[state] = true
			// Blindly try each legal boat configuration, relying on solve() to validate it.
			for _, Δ := range δ[b] {
				solve(state+Δ, append(steps, state))
			}
			delete(seen, state)
		}
	}
	// Initially, 3 missionaries, 3 cannibals, and 1 boat are on the wrong side.
	solve(0x331, []int{})
}
