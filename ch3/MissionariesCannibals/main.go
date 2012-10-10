// Missionaries and Cannibals
package main

import "fmt"

// The type state records what is on the wrong side.
type state struct {
	missionaries, cannibals, boats byte
}

const (
	MISSIONARIES = 3
	CANNIBALS    = 3
)

func main() {
	seen := map[int]bool{} // Record states seen in current path.
	steps := []state{}     // Record states seed in current path.
	var solve func(state)  // Declare solve type so it can recur.
	solve = func(s state) {
		// Print solution if all are across, and don't search further
		if s.missionaries == 0 && s.cannibals == 0 {
			fmt.Println(steps)
			return
		}
		// Reject solution if cannibals dominate a side.
		if s.missionaries != 0 && s.cannibals > s.missionaries ||
			s.missionaries != MISSIONARIES && s.cannibals < s.missionaries {
			return
		}
		key := int(s.missionaries)<<3 + int(s.cannibals)<<1 + int(s.boats)
		if !seen[key] {
			steps = append(steps, s)
			seen[key] = true
			// Determine the number of missionaries and cannibals on the same side as the boat.
			maxM, maxC := s.missionaries, s.cannibals
			if s.boats == 0 {
				maxM, maxC = MISSIONARIES-s.missionaries, CANNIBALS-s.cannibals
			}
			// Try sending the boat across with every valid combination of missionaries and cannibals.
			for m := byte(0); m <= maxM; m++ {
				for c := byte(0); c <= maxC; c++ {
					if m+c == 0 || m+c > 2 || m > 0 && c > m {
						continue
					}
					if s.boats != 0 {
						solve(state{s.missionaries - m, s.cannibals - c, 0})
					} else {
						solve(state{s.missionaries + m, s.cannibals + c, 1})
					}
				}
			}
			delete(seen, key)
			steps = steps[:len(steps)-1]
		}
	}
	solve(state{MISSIONARIES, CANNIBALS, 1})
}
