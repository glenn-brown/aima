package graph

// per AIMA
type Node struct {
	state  State
	parent *Node
	Action
	f Cost			// Cost plus heuristic cost.
	g Cost			// Path cost.
}

func solution(n *Node) []Action {
	r := []Action{}
	for n.Action != nil {
		r = append(r, n.Action)
		n = n.parent
	}
	return r
}
