package algo

// HeuristicFunc represents a heuristic function for pathfinding algorithms.
// It estimates the cost from the current node to the goal node.
// For A* to be optimal, the heuristic must be admissible (never overestimate).
//
// Params:
//
//	current: the current node
//	goal: the goal node
//
// Returns:
//
//	float64: estimated cost from current to goal
type HeuristicFunc func(current, goal *Node) float64

// Pathfinder represents a pathfinding algorithm that can find paths between nodes.
// This interface allows different pathfinding algorithms (A*, Dijkstra, etc.)
// to be used interchangeably.
type Pathfinder interface {
	// FindPath finds the optimal path from start to goal node.
	// Returns the path as a slice of nodes from start to goal,
	// or an error if no path exists.
	//
	// Params:
	//   start: starting node
	//   goal: destination node
	// Returns:
	//   []Node: path from start to goal (including both endpoints)
	//   error: if no path exists or invalid input
	FindPath(start, goal *Node) ([]*Node, error)

	// SetHeuristic configures the heuristic function for the pathfinder.
	// This allows switching between different heuristics (Manhattan, Euclidean, etc.)
	//
	// Params:
	//   heuristic: heuristic function to use
	SetHeuristic(heuristic HeuristicFunc)

	// SetGrid configures the grid/search space for the pathfinder.
	//
	// Params:
	//   grid: grid to search within
	SetGrid(grid GridInterface)
}

// GridInterface represents a searchable space for pathfinding algorithms.
// This interface abstracts the underlying representation (2D grid, graph, etc.)
// and provides the necessary operations for pathfinding.
type GridInterface interface {
	// GetNode returns the node at the specified coordinates.
	//
	// Params:
	//   x, y: coordinates
	// Returns:
	//   *Node: node at the position
	//   error: if position is invalid
	GetNode(x, y int) (*Node, error)

	// GetNeighbors returns all valid neighbors of a node.
	// Neighbors should exclude obstacles and out-of-bounds positions.
	//
	// Params:
	//   node: node to get neighbors for
	// Returns:
	//   []*Node: slice of neighboring nodes
	GetNeighbors(node *Node) []*Node

	// IsObstacle checks if a position is an obstacle or out of bounds.
	//
	// Params:
	//   x, y: coordinates to check
	// Returns:
	//   bool: true if position is blocked
	IsObstacle(x, y int) bool

	// GetCost returns the movement cost from one node to another.
	// This supports weighted grids and different terrain types.
	//
	// Params:
	//   from, to: source and destination nodes
	// Returns:
	//   float64: movement cost
	GetCost(from, to *Node) float64

	// Reset clears any algorithm-specific state from all nodes.
	// This allows reusing the grid for multiple pathfinding operations.
	Reset()
}
