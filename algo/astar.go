package algo

import (
	"fmt"
)

// AStar implements the A* pathfinding algorithm.
// A* is a best-first search algorithm that uses a heuristic to guide its search
// toward the goal, guaranteeing the shortest path when using an admissible heuristic.
//
// The algorithm maintains an open set (priority queue) of nodes to explore
// and a closed set of already explored nodes. It selects nodes with the lowest
// f-cost (g + h) for exploration, where g is the actual cost from start and
// h is the heuristic estimate to the goal.
type AStar struct {
	// grid is the search space we're pathfinding within
	grid GridInterface

	// heuristic function estimates cost from current node to goal
	heuristic HeuristicFunc

	// openSet contains nodes to be evaluated, ordered by f-cost
	openSet *PriorityQueue

	// closedSet contains nodes already evaluated (coordinates -> bool)
	closedSet map[coordinate]bool
}

// NewAStar creates a new A* pathfinder instance.
// By default, it uses the Manhattan heuristic which is optimal for 4-way movement.
//
// Returns:
//
//	*AStar: new A* pathfinder instance
func NewAStar() *AStar {
	return &AStar{
		heuristic: Manhattan, // Default to Manhattan heuristic
		openSet:   NewPriorityQueue(),
		closedSet: make(map[coordinate]bool),
	}
}

// SetGrid configures the grid/search space for the pathfinder.
// This must be called before FindPath.
//
// Params:
//
//	grid: grid to search within
func (a *AStar) SetGrid(grid GridInterface) {
	a.grid = grid
}

// SetHeuristic configures the heuristic function for the pathfinder.
// Common heuristics: Manhattan (4-way), Euclidean (any direction), Diagonal (8-way).
//
// Params:
//
//	heuristic: heuristic function to use
func (a *AStar) SetHeuristic(heuristic HeuristicFunc) {
	a.heuristic = heuristic
}

// FindPath finds the optimal path from start to goal using the A* algorithm.
// Returns the complete path including start and goal nodes, or an error if no path exists.
//
// Algorithm:
//  1. Initialize start node and add to open set
//  2. While open set is not empty:
//     - Get node with lowest f-cost
//     - If it's the goal, reconstruct and return path
//     - Move current node to closed set
//     - Examine each neighbor:
//     * Skip obstacles and nodes in closed set
//     * Calculate new g-cost for this path
//     * If new path is better, update neighbor and add to open set
//  3. If open set becomes empty, no path exists
//
// Params:
//
//	start: starting node
//	goal: destination node
//
// Returns:
//
//	[]*Node: path from start to goal (including both endpoints)
//	error: if no path exists or invalid input
func (a *AStar) FindPath(start, goal *Node) ([]*Node, error) {
	// Input validation
	if start == nil || goal == nil {
		return nil, fmt.Errorf("start and goal nodes cannot be nil")
	}

	if a.grid == nil {
		return nil, fmt.Errorf("grid must be set before calling FindPath")
	}

	if a.heuristic == nil {
		return nil, fmt.Errorf("heuristic must be set before calling FindPath")
	}

	// Check if start or goal are obstacles
	if start.IsObstacle {
		return nil, fmt.Errorf("start node at (%d, %d) is an obstacle", start.X, start.Y)
	}

	if goal.IsObstacle {
		return nil, fmt.Errorf("goal node at (%d, %d) is an obstacle", goal.X, goal.Y)
	}

	// If start equals goal, return single-node path
	if start.Equals(goal) {
		return []*Node{start}, nil
	}

	// Initialize algorithm state
	a.reset()

	// Setup start node
	start.G = 0
	start.H = a.heuristic(start, goal)
	start.CalculateF()
	start.Parent = nil

	// Add start node to open set
	a.openSet.PushNode(start)

	// Main A* loop
	for !a.openSet.IsEmpty() {
		// Get node with lowest f-cost
		current := a.openSet.PopNode()

		// Check if we reached the goal
		if current.Equals(goal) {
			return a.reconstructPath(current), nil
		}

		// Move current to closed set
		a.closedSet[coordinate{current.X, current.Y}] = true

		// Examine each neighbor
		neighbors := a.grid.GetNeighbors(current)
		for _, neighbor := range neighbors {
			// Skip obstacles and nodes in closed set
			if neighbor.IsObstacle || a.isInClosedSet(neighbor) {
				continue
			}

			// Calculate new g-cost for this path
			newG := current.G + a.grid.GetCost(current, neighbor)

			// Check if this path to neighbor is better
			inOpenSet := a.openSet.ContainsNode(neighbor)

			if !inOpenSet || newG < neighbor.G {
				// This path is better, record it
				neighbor.G = newG
				neighbor.H = a.heuristic(neighbor, goal)
				neighbor.CalculateF()
				neighbor.Parent = current

				// Add to open set if not already there
				if !inOpenSet {
					a.openSet.PushNode(neighbor)
				}
			}
		}
	}

	// Open set is empty but goal not reached - no path exists
	return nil, fmt.Errorf("no path found from (%d, %d) to (%d, %d)",
		start.X, start.Y, goal.X, goal.Y)
}

// reset clears the algorithm state for a fresh pathfinding operation.
// This clears the open and closed sets and resets the grid nodes.
func (a *AStar) reset() {
	a.openSet.Clear()
	a.closedSet = make(map[coordinate]bool)
	if a.grid != nil {
		a.grid.Reset()
	}
}

// isInClosedSet checks if a node is in the closed set (already fully explored).
//
// Params:
//
//	node: node to check
//
// Returns:
//
//	bool: true if node is in closed set
func (a *AStar) isInClosedSet(node *Node) bool {
	return a.closedSet[coordinate{node.X, node.Y}]
}

// reconstructPath builds the final path by following parent pointers from goal to start.
// Returns the path in forward order (start to goal).
//
// Params:
//
//	goalNode: the goal node with parent chain leading to start
//
// Returns:
//
//	[]*Node: path from start to goal
func (a *AStar) reconstructPath(goalNode *Node) []*Node {
	var path []*Node
	current := goalNode

	// Follow parent chain from goal to start
	for current != nil {
		path = append(path, current)
		current = current.Parent
	}

	// Reverse the path to get start -> goal order
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
