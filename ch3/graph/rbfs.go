package graph

func RecursiveBestFirstSearch(p *Problem) (solution []Action) {
	rv, _ := RBFS(p, &Node{p.InitialState, nil, nil, p.Zero, p.Zero}, p.Infinity)
	return rv
}

func RBFS(p *Problem, n *Node, f_limit Cost) (result []Action, c Cost) {
	if p.IsGoal(n.state) {
		return solution(n), n.g
	}
	successors := []*Node{}
	for _, a := range p.Actions(n.state) {
		successors = append(successors, childNode(p,n,a))
	}
	if len(successors) == 0 {
		return nil, p.Infinity
	}
	for _, s := range(successors) {
		s.f = max (s.g.Add(p.Heuristic(s.state)), n.f)
	}
	for {
		best, alternative := lowestFValueNodeIn(p, successors)
		if f_limit.Less(best.f) {
			return nil, best.f
		}
		result, best.f = RBFS(p, best, min(f_limit,alternative))
		if result != nil {
			return result, p.Infinity
		}
	}
	panic ("Never get here.")
}

func childNode(p *Problem, n *Node, a Action) *Node {
	g := n.g.Add(p.StepCost(n.state, a))
	nu := &Node{p.Result(n.state, a), n, a, g, g}
	nu.f = nu.g.Add(p.Heuristic(nu.state))
	return nu
}

func lowestFValueNodeIn(p *Problem, nn []*Node) (best *Node, alternative Cost) {
	if 0 == len(nn) {
		return nil, p.Infinity
	}
	best = nn[0]
	alternative = p.Infinity
	for i:=1; i<len(nn); i++ {
		n := nn[i]
		if n.f.Less(best.f) {
			alternative = best.f
			best = n
		}
	}
	return
}

func min(a, b Cost) Cost {
	if a.Less(b) {
		return a
	}
	return b
}
func max(a, b Cost) Cost {
	if a.Less(b) {
		return b
	}
	return a
}
