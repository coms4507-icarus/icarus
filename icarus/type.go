package icarus

import "sync"

type ThreadSafeGraph struct {
	Graph map[string][]string
	*sync.RWMutex
}

func NewThreadSafeGraph() *ThreadSafeGraph {
	return &ThreadSafeGraph{Graph: make(map[string][]string)}
}

func (graph *ThreadSafeGraph) addEdge(from string, to string) {
	graph.Graph[from] = append(graph.Graph[from], to)
}
