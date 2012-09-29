package graph

type FrontierCheapest struct {
	nodes []*Node
}

func (q FrontierCheapest) Empty() bool {
	return 0 == len(q.nodes)
}

func (q *FrontierCheapest) Pop() (n *Node) {
	n = q.nodes[0]
	l := len(q.nodes)
	q.nodes[0] = q.nodes[l-1]
	for i := 1; ; i *= 2 {
		c := q.nodes[i-1].pathCost
		if 2*i-1 < l && q.nodes[2*i-1].pathCost.Less(c) {
			q.nodes[i-1], q.nodes[2*i-1] = q.nodes[2*i-1], q.nodes[i-1]
		} else if 2*i < l && q.nodes[2*i].pathCost.Less(c) {
			q.nodes[i-1], q.nodes[2*i] = q.nodes[2*i-1], q.nodes[i-1]
		} else {
			break
		}
	}
	return
}

func (q *FrontierCheapest) Insert(n *Node) *FrontierCheapest {
	q.nodes = append(q.nodes, n)
	for child := len(q.nodes); child > 1; child >>= 1 {
		parent := child / 2
		if q.nodes[parent-1].pathCost.Less(q.nodes[child-1].pathCost) {
			break
		}
		q.nodes[parent-1], q.nodes[child-1] = q.nodes[child-1], q.nodes[parent-1]
	}
	return q
}
