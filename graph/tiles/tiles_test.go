package tiles

import (
	"github.com/glenn-brown/aima/graph"
	"fmt"
	"os"
	"math/rand"
	"testing"
)

// Actions are named as the cardinal directions, but the values are
// the offset to the tile in the specified direction.
const (
	N, E, W, S = -3, +1, -1, +3 // North, East, West, South
)

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
	for i:=uint(0); i<8; i++ {
		if uint(t) >> (4 * i) & 0xf == val {
			return i
		}
	}
	return 8;
}
func (t State) Get(pos uint) uint {
	if (pos < 8) {
		return (uint(t) >> (4 * pos)) & 0xf
	}
	rv := uint(36)
	for i:=uint(0); i<8; i++ {
		rv -= ((uint(t) >> (4 * i)) & 0xf)
	}
	return rv
}
func (t State) Set(pos, val uint) State {
	if (pos < 8) {
		return (t &^ (0xf << (4 * pos))) | ((State(val) << (4 * pos)))
	}
	return t
}
func Random() State {
	s := State(0x76543210)
	for i:=0; i<100; i++ {
		acts := s.Actions()
		act := acts[rand.Intn(len(acts))]
		zeroPos := s.Pos(0)
		otherPos := uint(int(zeroPos) + act.(int))
		otherVal := s.Get(otherPos)
		nu := s.Set(zeroPos, otherVal).Set(otherPos,0)
		// fmt.Fprintf(os.Stderr, "(%08x,%+d) -%v-> %08x\n", s, act, otherVal, nu)
		s = nu
	}
	fmt.Fprintf(os.Stderr, "Random() -> %08x\n", s)
	return s
}
func (s State) Actions() []graph.Action {
	return act[s.Pos(0)]
}
func (s State) Pack() (packed uint) {
	for i:=uint(0); i<8; i++ {
		packed = 3 * packed + (uint(s) >> (4 * i) & 0xf)
	}
	// fmt.Fprintf(os.Stderr, "Pack(%x) -> %d\n", uint(s), packed)
	return
}

//
// Cost
//

type Cost int

func (a Cost) Less(other graph.Cost) bool {
	return a < other.(Cost)
}
func (a Cost) Add(other graph.Cost) graph.Cost {
	return a + other.(Cost)
}

//
// Problem
//

type Problem struct{}

func (* Problem) IsGoal(s graph.State) bool {
	return s.(State) == 0x76543210
}
func (* Problem) Result(s graph.State, a graph.Action) graph.State {
	ts := s.(State)
	zeroPos := ts.Pos(0)
	otherPos := uint(int(zeroPos) + a.(int))
	otherVal := ts.Get(otherPos)
	rv :=  ts.Set(zeroPos, otherVal).Set(otherPos,0)
	// fmt.Fprintf (os.Stderr, "Result(%x,%d) -> %x\n", s, a, rv)
	return rv
}
func (* Problem) StepCost(s graph.State, a graph.Action) graph.Cost {
	return Cost(1)
}

//
// Seen
//

type Seen graph.Bitmap
func (ts Seen)Init() Seen {
	return Seen(graph.Bitmap(ts).Init(9*9*9*9*9*9*9*9))
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
func h(_s graph.State) graph.Cost {
	manhatten := [9][9]int{
		{0, 1, 2, 1, 2, 3, 2, 3, 4}, {1, 0, 1, 2, 1, 2, 3, 2, 3}, {2, 1, 0, 3, 2, 1, 4, 3, 2},
		{1, 2, 3, 0, 1, 2, 1, 2, 3}, {2, 1, 2, 1, 0, 1, 2, 1, 2}, {3, 2, 1, 2, 1, 0, 3, 2, 1},
		{2, 3, 4, 1, 2, 3, 0, 1, 2}, {3, 2, 3, 2, 1, 2, 1, 0, 1}, {4, 3, 2, 3, 2, 1, 2, 1, 0}}
	ts := _s.(State)
	c := 0
	for i:=uint(0); i<9; i++ {
		c += manhatten[i][(ts>>(4*i))&0xf]
	}
	return Cost(c)
}

//
// Testing
//

func Tiles(f graph.Frontier) []graph.Action {
	s := Seen{}.Init()
	return graph.Search(&Problem{}, Random(), &graph.FrontierFifo{}, Cost(0), &s)
}

func ExampleTilesCheapest() {
	fmt.Println(Tiles(graph.NewFrontierCheapest(h)))
	// Output: foo
}

func ExampleTilesFifo() {
	fmt.Println(Tiles(&graph.FrontierFifo{}))
	// Output: foo
}

func ExampleTilesLifo() {
	fmt.Println(Tiles(&graph.FrontierLifo{}))
	// Output: foo
}

func BenchmarkTilesFifo(b *testing.B) {
	Tiles(&graph.FrontierFifo{})
}

func BenchmarkTilesLifo(b *testing.B) {
	Tiles(&graph.FrontierLifo{})
}
func BenchmarkTilesCheapest(b *testing.B) {
	Tiles(graph.NewFrontierCheapest(h))
}
