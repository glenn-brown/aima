package graph

// Action wraps the client-defined type for specifying actions to
// apply to states.  It is opaque to the generic graph algorithms in this package.
type Action interface{}

type Frontier interface {
	// Empty returns true iff the frontier is empty.
	Empty() bool
	// Pop removes the next node to search.
	Pop() *Node
	// Insert adds a Node to the frontier.
	Insert(*Node) Frontier
}

type Seen interface {
	See(State)
	Saw(State) bool
}

type Solution []Action

type State interface {}

