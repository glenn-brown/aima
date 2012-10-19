package tiles

import (
	"fmt"
	"github.com/glenn-brown/aima/ch3/graph"
	"testing"
)

// 0x51287436
const X State = 0x12034567

func Tiles(f graph.Frontier, state State) []graph.Action {
	s := Seen{}.Init()
	p := New(state, 0x01234567)
	return graph.Search(p, f, &s)
}

func ExampleTilesPriorityQ() {
	fmt.Println(Tiles(graph.NewPriorityQ(heuristic),X))
	// Output: [E E]
}

func ExampleTilesFifo() {
	fmt.Println(Tiles(&graph.Fifo{},X))
	// Output: [E E]
}

func ExampleTilesLifo() {
	fmt.Println(Tiles(&graph.Lifo{},X))
	// Output: [E E]
}

func ExampleTilesRbfs() {
	fmt.Println(graph.RecursiveBestFirstSearch(New(X,0x01234567)))
}

func BenchmarkTilesLifo(b *testing.B) {
	for i:=0; i<100; i++{
		Tiles(&graph.Lifo{},Random())
	}
}

func BenchmarkTilesFifo(b *testing.B) {
	for i:=0; i<100; i++{
		Tiles(&graph.Fifo{},Random())
	}
}

func BenchmarkTilesPriorityQ(b *testing.B) {
	for i:=0; i<100; i++{
		Tiles(&graph.Fifo{},Random())
	}
}

