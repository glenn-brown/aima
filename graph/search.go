package graph

//
// Search uses the following interfaces.
//

// Action wraps the client-defined type for specifying actions to
// apply to states.
type Action interface{}

type Problem interface {
	// IsGoal returns true iff reaching the State ends the search.
	IsGoal(State) bool
	// Result returns the State resulting from applying Action to State.
	Result(State, Action) State
	// StepCost returns the Cost of applying Action to State.
	StepCost(State, Action) Cost
}

type State interface {
	// Actions returns the actions that may be applied to state.
	Actions() []Action
}

type Frontier interface {
	// Empty returns true iff the frontier is empty.
	Empty() bool
	// Pop removes the next node to search.
	Pop() *Node
	// Insert adds a Node to the frontier.
	Insert(*Node) Frontier
}

type Cost interface {
	// Less returns true iff the target Cost is less than the parameter Cost.
	Less(Cost) bool
	// Add returns the sum of the target and parameter Costs.
	Add(Cost) Cost
}

type Seen interface {
	See(State)
	Saw(State) bool
}

// Search searches for solutions to Problem p starting in State start.
// Frontier f stores the frontier and selects which node to explore next.
func Search(p Problem, start State, f Frontier, zero Cost, x Seen) []Action {
	// Initialize the frontier using the initial state of the problem.
	n := &Node{start, nil, nil, zero}
	f.Insert(n)
	// Initialize the explored set to be empty.
	// [We keep frontier nodes there, too, for efficiency.]
	x.See(start)
	// "loop do"
	for {
		// "if the frontier is empty, then return failure."
		if f.Empty() {
			return nil
		}
		// "Choose a leaf node and remove it from the frontier"
		n := f.Pop()
		// "If the node contains a goal state then return the corresponding solution."
		if p.IsGoal(n.state) {
			return solution(n)
		}
		// "Add the node to the explored set." [Just leave it in X]

		// "Expand the chosen node, adding the resulting nodes to the frontier
		// only if not in the frontier or explored set."
		for _, a := range n.state.Actions() {
			nu := childNode(p, n, a)
			if !x.Saw(nu.state) {
				f.Insert(nu)
			}
		}
	}
	return nil
}

// per AIMA
type Node struct {
	state  State
	parent *Node
	Action
	pathCost Cost
}

// per AIMA
func childNode(problem Problem, parent *Node, action Action) *Node {
	return &Node{
		problem.Result(parent.state, action),
		parent,
		action,
		parent.pathCost.Add(problem.StepCost(parent.state, action))}
}

func solution(n *Node) (r []Action) {
	for n.Action != nil {
		r = append(r, n.Action)
		n = n.parent
	}
	return
}
