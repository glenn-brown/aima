// A trivial greedy agent with no memory.
package agent

import "github.com/glenn-brown/vu"

type Agent struct{}

func New() *Agent {
	return &Agent{}
}

func (a *Agent) Move(goal vu.Point, percept vu.Points) vu.Point {
	if goal.Equals(vu.Point{0, 0}) {
		return goal
	}
	if len(percept) == 0 {
		return vu.Point{0, 0}
	}
	best := percept[0]
	for i := 1; i < len(percept); i++ {
		if percept[i].Sub(goal).Len() < best.Sub(goal).Len() {
			best = percept[i]
		}
	}
	return best
}

func (a *Agent) Render(w, h, d float64) {
}
