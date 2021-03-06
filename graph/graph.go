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
	mapCopy := make(map[string][]string)
	graph.lock.RLock()
	for key, value := range graph.graph {
		mapCopy[key] = value // arrays are value type for some reason
	}
	graph.lock.RUnlock()
	return mapCopy
}

// Get neighbours of a node
func (graph *ThreadSafeGraph) GetNeighbours(node string) []string {
	graph.lock.RLock()
	neighbours := graph.graph[node]
	graph.lock.RUnlock()
	return neighbours
}

// get size of graph
func (graph *ThreadSafeGraph) Size() int {
	graph.lock.RLock()
	size := len(graph.graph)
	graph.lock.RUnlock()
	return size
}
