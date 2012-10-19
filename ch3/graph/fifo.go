package graph

type Fifo struct {
	nodes []*Node
}

func (f Fifo) Empty() bool {
	return 0 == len(f.nodes)
}

func (f *Fifo) Pop() (n *Node) {
	n = f.nodes[0]
	f.nodes = f.nodes[1:]
	return
}

func (f *Fifo) Insert(n *Node) Frontier {
	f.nodes = append(f.nodes, n)
	return f
}
