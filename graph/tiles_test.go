package graph

type Action byte

const (
	N, E, W, S = -3, +1, -1, +3 // North, etc.
)

var act = [9][]int{
	{E, S}, {E, W, S}, {W, S},
	{N, E, S}, {N, E, W, S}, {N, W, S},
	{N, E}, {N, E, W}, {W, S}}

//
// State
//
// We represent board
// 876
// 543
// 21- as 0x8876543210, where 0 means empty and
// the first nibble is the empty position.

type TilesState uint64

func NewTilesState(a [3][3]int) (s TileState) {
	var zeroPos int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pos := 3*i + j
			if a[i][j] == 0 {
				zeroPos = pos
			}
			s |= uint64(a[i][j]) << (4 * pos)
		}
	}
	s |= uint64(zeroPos) << 36
}
func (s TilesState) Actions() []Action { return act[s.zeroPos()] }
func (s TilesState) ID() string        { string(s) }
func (s TilesState) zeroPos() (i int) {
	return int(s >> 36)
}

//
// Cost
//

type TilesCost int

func (a TilesCost) Less(other Cost) {
	return a < other.(int)
}
func (a TilesCost) Add(other Cost) {
	return Cost{a + other.(int)}
}

//
// Problem
//

type TilesProblem struct{}

func (TilesProblem) IsGoal(s State) { s.(TilesState) == 0x0876543210 }
func (TilesProblem) Result(s State, a Action) State {
	zp := s.zeroPos()
	otherPos := zp + a
	otherVal := (s >> (4 * otherPos)) & 0xf
	// Swap 0 and the other value in the state.
	s &^= 0xf<<36 | 0xf<<(4*otherPos)
	s |= otherPos<<36 | otherVal<<(4*zp)
	return s
}
func (TilesProblem) StepCost(s State, a Action) Cost {
	return Cost(TilesCost{1})
}

//
// Testing
//

func Tiles(f Frontier) {
	p := TilesProblem{}
	start := NewTilesState([3][3]int{{8, 7, 6}, {5, 4, 3}, {2, 1, 1}})
	fmt.Println(GraphSearch(p, start, FrontierFifo{}, TilesCost{0}))
}

func ExampleTilesCheapest() {
	fmt.Println(Tiles(Cheapest{}))
}

func BenchmarkTilesFifo(b *testing.B) {
	Tiles(FrontierFifo{})
}

func BenchmarkTilesLifo(b *testing.B) {
	Tiles(FrontierLifo{})
}

func BenchmarkTilesCheapest(b *testing.B) {
	Tiles(FrontierCheapest{})
}
