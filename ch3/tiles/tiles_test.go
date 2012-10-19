package tiles

import (
	"fmt"
	"github.com/glenn-brown/aima/ch3/graph"
	"testing"
)

func Tiles(f graph.Frontier, state State) []graph.Action {
	s := Seen{}.Init()
	p := New(0x12345678, 0x01234567)
	return graph.Search(p, state, f, Cost(0), &s)
}

func ExampleTilesPriorityQ() {
	fmt.Println(Tiles(graph.NewPriorityQ(heuristic),0x51287436))
	// Output: foo
}

func ExampleTilesFifo() {
	fmt.Println(Tiles(&graph.Fifo{},0x51287436))
	// Output: foo
}

func ExampleTilesLifo() {
	fmt.Println(Tiles(&graph.Lifo{},0x51287436))
	// Output: foo
}

func ExampleTilesRbfs() {
	fmt.Println(graph.RecursiveBestFirstSearch(New(0x51287436,0x01234567)))
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

