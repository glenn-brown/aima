
package tiles

import (
	"fmt"
	"github.com/glenn-brown/aima/ch3/graph"
	"math/rand"
	"os"
)

type Action int

// Actions are named as the cardinal directions, but the values are
// the offset to the tile in the specified direction.
const (
	N, E, W, S = Action(-3), Action(+1), Action(-1), Action(+3) // North, East, West, South
)
func (a Action) String() string {
	switch a {
	case N: return "N"
	case E: return "E"
	case W: return "W"
	case S: return "S"
	}
	panic("bad action")
}

// act maps the empty-tile position to a slice of possible moves (actions).
var act = [9][]graph.Action{
	{E, S}, {E, W, S}, {W, S},
	{N, E, S}, {N, E, W, S}, {N, W, S},
	{N, E}, {N, E, W}, {N, W}}

//
// State
//
// We represent board
// 876
// 543
// 21- as 0x12345678, where 0 means empty and
// the most-significant nibble is the empty position.

type State uint32

func (t State) Pos(val uint) (pos uint) {
	for i := uint(0); i < 8; i++ {
		if uint(t)>>(4*i)&0xf == val {
			return i
		}
	}
	return 8
}
func (t State) Get(pos uint) uint {
	if pos < 8 {
		return (uint(t) >> (4 * pos)) & 0xf
	}
	rv := uint(36)
	for i := uint(0); i < 8; i++ {
		rv -= ((uint(t) >> (4 * i)) & 0xf)
	}
	return rv
}
func (t State) Set(pos, val uint) State {
	if pos < 8 {
		return (t &^ (0xf << (4 * pos))) | (State(val) << (4 * pos))
	}
	return t
}
func Random() State {
	s := State(0x76543210)
	for i := 0; i < 100; i++ {
		acts := Actions(s)
		act := acts[rand.Intn(len(acts))]
		zeroPos := s.Pos(0)
		otherPos := uint(int(zeroPos) + act.(int))
		otherVal := s.Get(otherPos)
		nu := s.Set(zeroPos, otherVal).Set(otherPos, 0)
		fmt.Fprintf(os.Stderr, "(%08x,%+d) -%v-> %08x\n", s, act, otherVal, nu)
		s = nu
	}
	fmt.Fprintf(os.Stderr, "Random() -> %08x\n", s)
	return s
}
func Actions(s graph.State) []graph.Action {
	return act[s.(State).Pos(0)]
}
func (s State) Pack() (packed uint) {
	for i := uint(0); i < 8; i++ {
		packed = 3*packed + (uint(s) >> (4 * i) & 0xf)
	}
	// fmt.Fprintf(os.Stderr, "Pack(%x) -> %d\n", uint(s), packed)
	return
}

//
// Problem
//

func New(initialState, finalState State) *graph.Problem {
	return &graph.Problem{
		Actions,
		heuristic,
		graph.NewCost(1<<31-1),
		initialState,
		func(s graph.State) bool { return s.(State) == finalState },
		Result,
		StepCost,
		graph.NewCost(0)}
}
func Result(s graph.State, a graph.Action) graph.State {
	ts := s.(State)
	zeroPos := ts.Pos(0)
	otherPos := uint(int(zeroPos) + int(a.(Action)))
	otherVal := ts.Get(otherPos)
	rv := ts.Set(zeroPos, otherVal).Set(otherPos, 0)
	// fmt.Fprintf (os.Stderr, "Result(%x,%d) -> %x\n", s, a, rv)
	return rv
}
func StepCost(s graph.State, a graph.Action) graph.Cost {
	return graph.NewCost(1)
}

//
// Seen
//

type Seen graph.Bitmap

func (ts Seen) Init() Seen {
	return Seen(graph.Bitmap(ts).Init(9 * 9 * 9 * 9 * 9 * 9 * 9 * 9))
}
func (ts *Seen) See(s graph.State) {
	(*graph.Bitmap)(ts).Set(s.(State).Pack(), true)
}
func (ts *Seen) Saw(s graph.State) bool {
	return (*graph.Bitmap)(ts).Get(s.(State).Pack())
}

//
// Heuristic function
//

// Function h returns a lower bound on cost of solving the puzzle.
// Each piece must move at least the manhatten distance to its correct position.
func heuristic(_s graph.State) graph.Cost {
	manhatten := [9][9]int{
		{0, 1, 2, 1, 2, 3, 2, 3, 4}, {1, 0, 1, 2, 1, 2, 3, 2, 3}, {2, 1, 0, 3, 2, 1, 4, 3, 2},
		{1, 2, 3, 0, 1, 2, 1, 2, 3}, {2, 1, 2, 1, 0, 1, 2, 1, 2}, {3, 2, 1, 2, 1, 0, 3, 2, 1},
		{2, 3, 4, 1, 2, 3, 0, 1, 2}, {3, 2, 3, 2, 1, 2, 1, 0, 1}, {4, 3, 2, 3, 2, 1, 2, 1, 0}}
	ts := _s.(State)
	c := 0
	for i := uint(0); i < 9; i++ {
		tile := ts.Get(i)
		if tile != 0 {
			c += manhatten[i][tile]
		}
	}
	return graph.NewCost(c)
}
