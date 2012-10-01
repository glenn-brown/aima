package graph

type FrontierCheapest struct {
	nodes []*Node
	h     func(State) Cost
}

func NewFrontierCheapest(h func(State) Cost) *FrontierCheapest {
	return &FrontierCheapest{[]*Node{}, h}
}

func (q FrontierCheapest) Empty() bool {
	return 0 == len(q.nodes)
}

func (q *FrontierCheapest) Pop() *Node {
	// We will return the first node.
	n := q.nodes[0]
	// Replace the returning node with the last node.
	l := len(q.nodes) - 1
	q.nodes[0] = q.nodes[l]
	q.nodes = q.nodes[:l]
	if l == 0 {
		return n
	}
	// Repeatedly exchange the moved node with any higher priority child.
	c := q.cost(0)
	// fmt.Fprintf(os.Stderr, "Pop() -> %v cost=%v cnt=%v\n", n, c, len(q.nodes));
	for i := 0; ; {
		c1 := 2*i + 1
		c2 := 2*i + 2
		if c1 < l && q.cost(c1).Less(c) && !(c2 < l && q.cost(c2).Less(q.cost(c1))) {
			q.nodes[i], q.nodes[c1] = q.nodes[c1], q.nodes[i]
			i = c1
		} else if c2 < l && q.cost(c2).Less(c) {
			q.nodes[i], q.nodes[c2] = q.nodes[c2], q.nodes[i]
			i = c2
		} else {
			break
		}
	}
	return n
}

func (q *FrontierCheapest) Insert(n *Node) Frontier {
	// Append new entry
	l := len(q.nodes)
	q.nodes = append(q.nodes, n)
	// And exchange it with each lesser parent.
	c := q.cost(l)
	// fmt.Fprintf(os.Stderr, "Insert(%v) cost=%v cnt=%v\n", n, c, l+1)
	for child := l; child > 0; {
		parent := (child - 1) / 2
		if q.cost(parent).Less(c) {
			break
		}
		q.nodes[parent], q.nodes[child] = q.nodes[child], q.nodes[parent]
		child = parent
	}
	return q
}

// Function cost() returns the estimated cost of a solution starting a node I.
func (q *FrontierCheapest) cost(i int) Cost {
	c := q.nodes[i].pathCost
	if q.h != nil {
		c = c.Add(q.h(q.nodes[i].state))
	}
	return c
}
