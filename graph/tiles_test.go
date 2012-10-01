package graph

import (
	"fmt"
	"math/rand"
	"testing"
)

// Actions are named as the cardinal directions, but the values are
// the offset to the tile in the specified direction.
const (
	N, E, W, S = -3, +1, -1, +3 // North, East, West, South
)

// act maps the empty-tile position to a slice of possible moves (actions).
var act = [9][]Action{
	{E, S}, {E, W, S}, {W, S},
	{N, E, S}, {N, E, W, S}, {N, W, S},
	{N, E}, {N, E, W}, {W, S}}

//
// State
//
// We represent board
// 876
// 543
// 21- as 0x8012345678, where 0 means empty and
// the most-significant nibble is the empty position.

type TilesState struct {
	tiles [9]byte
	zeroPos byte
}

func RandomTilesState() TilesState {
	s := TilesState{[9]byte{0,1,2,3,4,5,6,7,8},0}
	for i := range s.tiles {
		r := rand.Intn(9)
		s.tiles[i], s.tiles[r] = s.tiles[r], s.tiles[i]
	}
	for i, v := range s.tiles {
		if v == 0 {
			s.zeroPos = byte(i)
		}
	}
	return s;
}
func (s TilesState) Actions() []Action { return act[s.zeroPos] }
func (s TilesState) ID() (id uint) {
	for i:=0; i<8; i++ {
		id = 9 * id + uint(s.tiles[i])
	}
	return
}

//
// Cost
//

type TilesCost int

func (a TilesCost) Less(other Cost) bool {
	return a < other.(TilesCost)
}
func (a TilesCost) Add(other Cost) Cost {
	return a + other.(TilesCost)
}

//
// Problem
//

type TilesProblem struct{}

func (TilesProblem) IsGoal(s State) bool {
	ts := s.(TilesState)
	for i := range ts.tiles {
		if ts.tiles[i] != byte(i) {
			return false
		}
	}
	return true
}
func (TilesProblem) Result(s State, a Action) State {
	ts := s.(TilesState)
	zeroPos := ts.zeroPos
	other := zeroPos + byte(a.(int))
	ts.tiles[zeroPos], ts.tiles[other] = ts.tiles[other], ts.tiles[zeroPos]
	return ts
}
func (TilesProblem) StepCost(s State, a Action) Cost {
	return TilesCost(1)
}

//
// Seen
//

type TilesSeen struct {
	bitmap *Bitmap
}
func (ts TilesSeen)Init () (*TilesSeen) {
	ts.bitmap = NewBitmap(9*9*9*9*9*9*9*9)
	return &ts;
}
func (ts *TilesSeen) See(s State) {
	ts.bitmap.Set(s.(TilesState).ID(), true)
}
func (ts *TilesSeen) Saw(s State) bool {
	return ts.bitmap.Get(s.(TilesState).ID())
}

//
// Heuristic function
//

// Function h returns a lower bound on cost of solving the puzzle.
// Each piece must move at least the manhatten distance to its correct position.
func h(_s State) Cost {
	manhatten := [9][9]int{
		{0, 1, 2, 1, 2, 3, 2, 3, 4}, {1, 0, 1, 2, 1, 2, 3, 2, 3}, {2, 1, 0, 3, 2, 1, 4, 3, 2},
		{1, 2, 3, 0, 1, 2, 1, 2, 3}, {2, 1, 2, 1, 0, 1, 2, 1, 2}, {3, 2, 1, 2, 1, 0, 3, 2, 1},
		{2, 3, 4, 1, 2, 3, 0, 1, 2}, {3, 2, 3, 2, 1, 2, 1, 0, 1}, {4, 3, 2, 3, 2, 1, 2, 1, 0}}
	s := _s.(TilesState)
	c := 0
	for i, v := range s.tiles {
		c += manhatten[i][v]
	}
	return TilesCost(c)
}

//
// Testing
//

func Tiles(f Frontier) []Action {
	return Search(TilesProblem{}, RandomTilesState(), &FrontierFifo{}, TilesCost(0), TilesSeen{}.Init())
}

func ExampleTilesCheapest() {
	fmt.Println(Tiles(NewFrontierCheapest(h)))
	// Output: foo
}

func BenchmarkTilesFifo(b *testing.B) {
	Tiles(&FrontierFifo{})
}

func BenchmarkTilesLifo(b *testing.B) {
	Tiles(&FrontierLifo{})
}

func BenchmarkTilesCheapest(b *testing.B) {
	Tiles(NewFrontierCheapest(h))
}
