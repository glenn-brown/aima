package graph

type Problem struct {
	// Actions returns the actions that may be applied to state.
	Actions func(s State) []Action
	// The estimated (lower bound) cost of getting of n.State to the goal.
	Heuristic func(s State) Cost
	// The never-exceeded cost.
	Infinity Cost
	// InitialState returns the starting State of the search.
	InitialState State
	// IsGoal returns true iff reaching the State ends the search.
	IsGoal func(State) bool
	// Result returns the State resulting from applying Action to State.
	Result func(State, Action) State
	// StepCost returns the Cost of applying Action to State.
	StepCost func(State, Action) Cost
	// Zero returns the 0 Cost for the problem.
	Zero Cost
}

func NewProblem(
	Actions func(s State) []Action,
	Heuristic func(s State) Cost,
	Infinity Cost,
	InitialState State,
	IsGoal func(State) bool,
	Result func(State, Action) State,
	StepCost func(State, Action) Cost,
	Zero Cost)(*Problem) {
	
	return &Problem{Actions, Heuristic, Infinity, InitialState, IsGoal, Result, StepCost, Zero}
}
