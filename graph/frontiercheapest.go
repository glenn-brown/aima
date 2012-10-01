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

func (q *FrontierCheapest) Pop() (n *Node) {
	n = q.nodes[0]
	l := len(q.nodes)
	q.nodes[0] = q.nodes[l-1]
	for i := 1; ; i *= 2 {
		c := q.cost(i - 1)
		if 2*i-1 < l && q.cost(2*i-1).Less(c) {
			q.nodes[i-1], q.nodes[2*i-1] = q.nodes[2*i-1], q.nodes[i-1]
		} else if 2*i < l && q.cost(2*i).Less(c) {
			q.nodes[i-1], q.nodes[2*i] = q.nodes[2*i-1], q.nodes[i-1]
		} else {
			break
		}
	}
	return
}

func (q *FrontierCheapest) Insert(n *Node) Frontier {
	q.nodes = append(q.nodes, n)
	for child := len(q.nodes); child > 1; child >>= 1 {
		parent := child / 2
		if q.cost(parent - 1).Less(q.cost(child - 1)) {
			break
		}
		q.nodes[parent-1], q.nodes[child-1] = q.nodes[child-1], q.nodes[parent-1]
	}
	return q
}

// Function cost() returns the estimated cost of a solution starting a node I.
func (q *FrontierCheapest) cost(i int) Cost {
	c := q.nodes[i].pathCost
	if q.h != nil {
		c.Add(q.h(q.nodes[i].state))
	}
	return c
}
