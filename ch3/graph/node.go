package graph

// per AIMA
type Node struct {
	state  State
	parent *Node
	Action
	f Cost // Cost plus heuristic cost.
	g Cost // Path cost.
}

func solution(n *Node) []Action {
	// Build actions from finish to start.
	r := []Action{}
	for n.Action != nil {
		r = append(r, n.Action)
		n = n.parent
	}
	// Reverse the order.
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
