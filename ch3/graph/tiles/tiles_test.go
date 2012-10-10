package tiles

import (
	"fmt"
	"github.com/glenn-brown/aima/graph"
	"testing"
)

func Tiles(f graph.Frontier, state State) []graph.Action {
	s := Seen{}.Init()
	return graph.Search(&Problem{}, state, f, Cost(0), &s)
}

func ExampleTilesCheapest() {
	fmt.Println(Tiles(graph.NewFrontierCheapest(heuristic),0x51287436))
	// Output: foo
}

func ExampleTilesFifo() {
	fmt.Println(Tiles(&graph.FrontierFifo{},0x51287436))
	// Output: foo
}

func ExampleTilesLifo() {
	fmt.Println(Tiles(&graph.FrontierLifo{},0x51287436))
	// Output: foo
}

func BenchmarkTilesLifo(b *testing.B) {
	for i:=0; i<100; i++{
		Tiles(&graph.FrontierLifo{},Random())
	}
}

func BenchmarkTilesFifo(b *testing.B) {
	for i:=0; i<100; i++{
		Tiles(&graph.FrontierFifo{},Random())
	}
}

func BenchmarkTilesCheapest(b *testing.B) {
	for i:=0; i<100; i++{
		Tiles(&graph.FrontierFifo{},Random())
	}
}

