package graph

import (
	"fmt"
	"os"
)

// Search searches for solutions to Problem p starting in State start.
// Frontier f stores the frontier and selects which node to explore next.
func Search(p *Problem, f Frontier, x Seen) []Action {
	// Initialize the frontier using the initial state of the problem.
	n := &Node{p.InitialState, nil, nil, p.Zero, p.Zero}
	f.Insert(n)
	// Initialize the explored set to be empty.
	// [We keep frontier nodes there, too, for efficiency.]
	x.See(p.InitialState)
	// "loop do"
	for {
		// "if the frontier is empty, then return failure."
		if f.Empty() {
			fmt.Fprintf(os.Stderr, "No solution\n")
			return nil
		}
		// "Choose a leaf node and remove it from the frontier"
		n := f.Pop()
		// fmt.Fprintf(os.Stderr, "State %x\n", n.state)
		// "If the node contains a goal state then return the corresponding solution."
		if p.IsGoal(n.state) {
			fmt.Fprintf(os.Stderr, "Goal reached: %v\n", n)
			return solution(n)
		}
		// "Add the node to the explored set." [Just leave it in X]

		// "Expand the chosen node, adding the resulting nodes to the frontier
		// only if not in the frontier or explored set."
		for _, a := range p.Actions(n.state) {
			// fmt.Fprintf(os.Stderr, "Action %v\n", a);
			result := p.Result(n.state, a)
			if !x.Saw(result) {
				x.See(result)
				// fmt.Fprintf(os.Stderr, "New State %x\n", n.state)skiplist/
				cost := n.g.Add(p.StepCost(n.state, a)).(Cost)
				nu := &Node{result, n, a, cost, cost}
				nu.f = nu.g.Add(p.Heuristic(nu.state)).(Cost)
				// fmt.Fprintf(os.Stderr, "Inserting state %x\n", nu.state);
				f.Insert(nu)
			} else {
				// FIXME: in general, need to adjust cost.
			}
		}
	}
	return nil
}
