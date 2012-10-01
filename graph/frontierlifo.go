package graph

// FrontierLifo is a simple Last In First Out (LIFO) queue.
type FrontierLifo struct {
	nodes []*Node
}

// Empty returns true iff the LIFO is empty.
func (l FrontierLifo) Empty() bool {
	return 0 == len(l.nodes)
}

// Pop returns the most recently added entry in the LIFO.
func (l *FrontierLifo) Pop() (n *Node) {
	n = l.nodes[len(l.nodes)-1]
	l.nodes = l.nodes[:len(l.nodes)-1]
	return
}

// Insert adds an entry into the LIFO.
func (l *FrontierLifo) Insert(n *Node) Frontier {
	l.nodes = append(l.nodes, n)
	return l
}
