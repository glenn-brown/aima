package romania

import (
	"github.com/glenn-brown/aima/ch3/graph"
	"testing"
)

func BenchmarkAradRBFS(b *testing.B) {
	for i := b.N; i >= 0; i-- {
		graph.RecursiveBestFirstSearch(New(Arad))
	}
}
