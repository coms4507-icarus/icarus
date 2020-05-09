package icarus

import "sync"

type ThreadSafeGraph struct {
	graph map[string][]string
	*sync.RWMutex
}

func newThreadSafeGraph() *ThreadSafeGraph {
	return &ThreadSafeGraph{graph: make(map[string][]string)}
}

func (graph *ThreadSafeGraph) addEdge(from string, to string) {
	graph.graph[from] = append(graph.graph[from], to)
}
