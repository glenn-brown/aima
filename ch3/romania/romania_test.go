package romania

import (
	"fmt"
	"github.com/glenn-brown/aima/ch3/graph"
	"testing"
)

func ExampleNew() {
	fmt.Println (graph.RecursiveBestFirstSearch(New(Arad)))
	// Output: [Sibiu RimnicuVilcea Pitesti Bucharest]
}

func BenchmarkAradRBFS(b *testing.B) {
	for i := b.N; i >= 0; i-- {
		graph.RecursiveBestFirstSearch(New(Arad))
	}
}
