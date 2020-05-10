package graph

import "sync"

// Thread safe graph to work with goroutines
type ThreadSafeGraph struct {
	Graph map[string][]string
	lock  sync.RWMutex
}

// Create a new thread safe graph
func NewThreadSafeGraph() *ThreadSafeGraph {
	return &ThreadSafeGraph{Graph: make(map[string][]string)}
}

// Add a new edge to this graph
func (graph *ThreadSafeGraph) AddEdge(from string, to string) {
	graph.lock.Lock()
	defer graph.lock.Unlock()
	graph.Graph[from] = append(graph.Graph[from], to)
}

// Check if a node exists
func (graph *ThreadSafeGraph) NodeExists(node string) bool {
	graph.lock.RLock()
	defer graph.lock.RUnlock()
	_, ok := graph.Graph[node]
	return ok
}

// get size of graph
func (graph *ThreadSafeGraph) Size() int {
	graph.lock.RLock()
	defer graph.lock.RUnlock()
	return len(graph.Graph)
}
