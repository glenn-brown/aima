package graph

type FrontierFifo struct {
	nodes []*Node
}

func (f FrontierFifo) Empty() bool {
	return 0 == len(f.nodes)
}

func (f *FrontierFifo) Pop() (n *Node) {
	n = f.nodes[0]
	f.nodes = f.nodes[1:]
	return
}

func (f *FrontierFifo) Insert(n *Node) *FrontierFifo {
	f.nodes = append(f.nodes, n)
	return f
}
