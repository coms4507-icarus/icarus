package graph

import "sync"

// Thread safe graph to work with goroutines
type ThreadSafeGraph struct {
	graph map[string][]string
	lock  sync.RWMutex
}

// Create a new thread safe graph
func NewThreadSafeGraph() *ThreadSafeGraph {
	return &ThreadSafeGraph{graph: make(map[string][]string)}
}

// Add a new edge to this graph
func (graph *ThreadSafeGraph) AddEdge(from string, to string) {
	graph.lock.Lock()
	graph.graph[from] = append(graph.graph[from], to)
	graph.lock.Unlock()
}

// Check if a node exists
func (graph *ThreadSafeGraph) NodeExists(node string) bool {
	graph.lock.RLock()
	_, ok := graph.graph[node]
	graph.lock.RUnlock()
	return ok
}

// Get a json-serializable copy of the graph
func (graph *ThreadSafeGraph) Graph() map[string][]string {
	var mapCopy map[string][]string
	graph.lock.RLock()
	for key, value := range graph.graph {
		copy(mapCopy[key], value)
	}
	graph.lock.RUnlock()
	return mapCopy
}

// get size of graph
func (graph *ThreadSafeGraph) Size() int {
	graph.lock.RLock()
	size := len(graph.graph)
	graph.lock.RUnlock()
	return size
}
