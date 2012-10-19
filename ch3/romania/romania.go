package romania

import "github.com/glenn-brown/aima/ch3/graph"

type City int

const (
	Arad City = iota
	Bucharest
	Craiova
	Drobeta
	Eforie
	Fagaras
	Giurgiu
	Hirsova
	Iasi
	Lugoj
	Mehadia
	Neamt
	Oradea
	Pitesti
	RimnicuVilcea
	Sibiu
	Timisoara
	Urziceni
	Vaslui
	Zerind
	Cities
)

func (c City) String() string {
	switch c {
	case Arad:
		return "Arad"
	case Bucharest:
		return "Bucharest"
	case Craiova:
		return "Craiova"
	case Drobeta:
		return "Drobeta"
	case Eforie:
		return "Eforie"
	case Fagaras:
		return "Fagaras"
	case Giurgiu:
		return "Giurgiu"
	case Hirsova:
		return "Hirsova"
	case Iasi:
		return "Iasi"
	case Lugoj:
		return "Lugoj"
	case Mehadia:
		return "Mehadia"
	case Neamt:
		return "Neamt"
	case Oradea:
		return "Oradea"
	case Pitesti:
		return "Pitesti"
	case RimnicuVilcea:
		return "RimnicuVilcea"
	case Sibiu:
		return "Sibiu"
	case Timisoara:
		return "Timisoara"
	case Urziceni:
		return "Urziceni"
	case Vaslui:
		return "Vaslui"
	case Zerind:
		return "Zerind"
	}
	return "?City?"
}

var roads = []struct {
	a City
	d int
	b City
}{
	{Oradea, 71, Zerind},
	{Oradea, 151, Sibiu},
	{Zerind, 75, Arad},
	{Arad, 140, Sibiu},
	{Sibiu, 99, Fagaras},
	{Fagaras, 211, Bucharest},
	{Bucharest, 85, Urziceni},
	{Urziceni, 142, Vaslui},
	{Vaslui, 92, Iasi},
	{Iasi, 87, Neamt},
	{Urziceni, 98, Hirsova},
	{Hirsova, 86, Eforie},
	{Bucharest, 90, Giurgiu},
	{Bucharest, 101, Pitesti},
	{Pitesti, 97, RimnicuVilcea},
	{RimnicuVilcea, 80, Sibiu},
	{RimnicuVilcea, 146, Craiova},
	{Craiova, 120, Drobeta},
	{Drobeta, 75, Mehadia},
	{Mehadia, 70, Lugoj},
	{Lugoj, 111, Timisoara},
	{Timisoara, 118, Arad},
}

// Straight line distances to Bucharest:
var toBucharest = []int{
	Arad: 366, Bucharest: 0, Craiova: 160, Drobeta: 242, Eforie: 161, Fagaras: 176, Giurgiu: 77,
	Hirsova: 151, Iasi: 226, Lugoj: 244, Mehadia: 241, Neamt: 234, Oradea: 380, Pitesti: 100,
	RimnicuVilcea: 193, Sibiu: 253, Timisoara: 329, Urziceni: 80, Vaslui: 199, Zerind: 374}

// Straight-line distance between cities.
var edge = [Cities]map[City]int{}

func init() {
	for i := City(Cities - 1); i >= 0; i-- {
		edge[i] = map[City]int{}
	}
	for _, r := range roads {
		edge[r.a][r.b] = r.d
		edge[r.b][r.a] = r.d
	}
}

func New(from City) *graph.Problem {
	return &graph.Problem{
		Actions: func(s graph.State) []graph.Action {
			rv := []graph.Action{}
			for k := range edge[s.(City)] {
				rv = append(rv, k)
			}
			return rv
		},
		Heuristic:    func(s graph.State) graph.Cost { return graph.NewCost(toBucharest[s.(City)]) },
		Infinity:     graph.NewCost(1<<31 - 1),
		InitialState: from,
		IsGoal:       func(s graph.State) bool { return s.(City) == Bucharest },
		Result:       func(s graph.State, a graph.Action) graph.State { return a.(City) },
		StepCost: func(s graph.State, a graph.Action) graph.Cost {
			return graph.NewCost(edge[s.(City)][a.(City)])
		},
		Zero: graph.NewCost(0),
	}
}
