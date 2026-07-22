package main

import (
	"math"
)

// MaxFlow calculates the maximum flow from source to sink using Edmonds-Karp
func MaxFlow(g *GraphList, source, sink int) (int, [][]int) {
	// Create residual graph - initially same as capacity
	residual := make([][]int, len(g.stationsNames))
	for i := range residual {
		residual[i] = make([]int, len(g.stationsNames))
		copy(residual[i], g.adjMatrix[i])
	}

	parent := make([]int, len(g.stationsNames))
	maxFlow := 0
	// var paths [][]int

	// While there exists an augmenting path from source to sink
	for g.bfs(source, sink, parent, residual) {

		// Find minimum residual capacity along the path
		pathFlow := math.MaxInt32
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			if residual[u][v] < pathFlow {
				pathFlow = residual[u][v]
			}
		}

		// // path := []int{}
		// for v := sink; v != source; v = parent[v] {
		// 	path = append(path, v)
		// 	// v = parent[v] // move to next node in path
		// }
		// path = append(path, source)
		// reverse(path)
		// paths = append(paths, path)

		// Update residual capacities
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			residual[u][v] -= pathFlow // Decrease forward edge
			residual[v][u] += pathFlow // Increase backward edge
		}
		maxFlow += pathFlow
		// fmt.Println("NEW")
		// fmt.Println(parent)
		// for child, value := range parent {
		// 	if value == 0 {
		// 		continue
		// 	}

		// 	if value == -1 {
		// 		fmt.Println("connection:", g.stationsNames[child+2], "->", g.stationsNames[0])
		// 		continue
		// 	}

		// 	fmt.Println("connection:", g.stationsNames[child], "->", g.stationsNames[value])
		// }
	}
	paths := extractPaths(g, residual, source, sink)
	return maxFlow, paths
}

// helper to reverse path
func reverse(path []int) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
func extractPaths(g *GraphList, residual [][]int, source, sink int) [][]int {
	n := len(g.stationsNames)

	// Build flow-used graph
	flowUsed := make([][]int, n)
	for i := range flowUsed {
		flowUsed[i] = make([]int, n)
	}

	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if g.adjMatrix[u][v] == 1 && residual[u][v] == 0 {
				flowUsed[u][v] = 1
			}
		}
	}

	var paths [][]int

	for {
		visited := make([]bool, n)
		path := []int{}

		if !dfsExtract(source, sink, flowUsed, visited, &path) {
			break
		}

		reverse(path)
		paths = append(paths, path)
	}

	return paths
}

func dfsExtract(u, sink int, flowUsed [][]int, visited []bool, path *[]int) bool {
	visited[u] = true

	if u == sink {
		*path = append(*path, u)
		return true
	}

	for v := 0; v < len(flowUsed); v++ {
		if flowUsed[u][v] == 1 && !visited[v] {
			if dfsExtract(v, sink, flowUsed, visited, path) {
				flowUsed[u][v] = 0 // remove edge so it's not reused
				*path = append(*path, u)
				return true
			}
		}
	}

	return false
}

// bfs performs breadth-first search to find a path from source to sink
// Returns true if a path exists, and fills the parent array with the path
func (g *GraphList) bfs(source, sink int, parent []int, residual [][]int) bool {
	// Initialize visited array
	visited := make([]bool, len(g.stationsNames))

	// Create queue and add source
	queue := []int{source}
	visited[source] = true
	parent[source] = -1

	// Standard BFS
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// Check all adjacent vertices
		for v := 0; v < len(g.stationsNames); v++ {
			// If not visited and has residual capacity
			if !visited[v] && residual[u][v] > 0 {
				queue = append(queue, v)
				parent[v] = u
				visited[v] = true

				// Early exit if we reached the sink
				if v == sink {
					return true
				}
			}
		}
	}

	return false
}
