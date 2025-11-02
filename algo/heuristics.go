package algo

import (
	"math"
)

// Manhattan calculates the Manhattan distance heuristic between two nodes.
// This heuristic is optimal for 4-way movement (up, down, left, right only).
// It's admissible and consistent, guaranteeing optimal A* paths.
//
// Formula: |x1 - x2| + |y1 - y2|
//
// Params:
//
//	current: current node
//	goal: goal node
//
// Returns:
//
//	float64: Manhattan distance
func Manhattan(current, goal *Node) float64 {
	dx := float64(current.X - goal.X)
	dy := float64(current.Y - goal.Y)

	// Calculate absolute values
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

// Euclidean calculates the Euclidean distance heuristic between two nodes.
// This is the straight-line distance and is admissible for any movement.
// More accurate than Manhattan for 8-way movement but computationally expensive.
//
// Formula: sqrt((x1 - x2)² + (y1 - y2)²)
//
// Params:
//
//	current: current node
//	goal: goal node
//
// Returns:
//
//	float64: Euclidean distance
func Euclidean(current, goal *Node) float64 {
	dx := float64(current.X - goal.X)
	dy := float64(current.Y - goal.Y)

	return math.Sqrt(dx*dx + dy*dy)
}

// EuclideanSquared calculates the squared Euclidean distance heuristic.
// This avoids the expensive square root operation while maintaining
// the same relative ordering of nodes. Use when you only need to compare
// distances, not their actual values.
//
// Formula: (x1 - x2)² + (y1 - y2)²
//
// Params:
//
//	current: current node
//	goal: goal node
//
// Returns:
//
//	float64: squared Euclidean distance
func EuclideanSquared(current, goal *Node) float64 {
	dx := float64(current.X - goal.X)
	dy := float64(current.Y - goal.Y)

	return dx*dx + dy*dy
}

// Diagonal calculates the diagonal distance heuristic (also known as Chebyshev distance).
// This is optimal for 8-way movement where diagonal moves cost the same as
// straight moves. It's the maximum of the horizontal and vertical distances.
//
// Formula: max(|x1 - x2|, |y1 - y2|)
//
// Params:
//
//	current: current node
//	goal: goal node
//
// Returns:
//
//	float64: diagonal distance
func Diagonal(current, goal *Node) float64 {
	dx := float64(current.X - goal.X)
	dy := float64(current.Y - goal.Y)

	// Calculate absolute values
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	// Return the maximum
	if dx > dy {
		return dx
	}
	return dy
}

// DiagonalWithCost calculates a more accurate diagonal distance heuristic
// for 8-way movement where diagonal moves cost √2 ≈ 1.414 times straight moves.
// This provides better guidance for A* when diagonal movement is more expensive.
//
// Formula: D * max(dx, dy) + (D√2 - D) * min(dx, dy)
// where D = 1.0 (cost of straight movement)
//
// Params:
//
//	current: current node
//	goal: goal node
//
// Returns:
//
//	float64: diagonal distance with movement cost
func DiagonalWithCost(current, goal *Node) float64 {
	dx := float64(current.X - goal.X)
	dy := float64(current.Y - goal.Y)

	// Calculate absolute values
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	// Cost constants
	const straightCost = 1.0
	const diagonalCost = 1.414213562373095 // √2

	// Calculate min and max distances
	var minDist, maxDist float64
	if dx < dy {
		minDist = dx
		maxDist = dy
	} else {
		minDist = dy
		maxDist = dx
	}

	return straightCost*(maxDist-minDist) + diagonalCost*minDist
}

// Zero returns zero heuristic (equivalent to Dijkstra's algorithm).
// This turns A* into Dijkstra's algorithm, which explores uniformly
// in all directions. Useful for comparison or when no good heuristic exists.
//
// Params:
//
//	current: current node (unused)
//	goal: goal node (unused)
//
// Returns:
//
//	float64: always 0.0
func Zero(current, goal *Node) float64 {
	return 0.0
}

// GetHeuristicByName returns a heuristic function by name.
// This is useful for configuration and testing different heuristics.
//
// Supported names:
//   - "manhattan": Manhattan distance
//   - "euclidean": Euclidean distance
//   - "euclidean_squared": Squared Euclidean distance
//   - "diagonal": Diagonal/Chebyshev distance
//   - "diagonal_cost": Diagonal with movement costs
//   - "zero": Zero heuristic (Dijkstra)
//
// Params:
//
//	name: name of the heuristic
//
// Returns:
//
//	HeuristicFunc: heuristic function
//	bool: true if name is valid
func GetHeuristicByName(name string) (HeuristicFunc, bool) {
	switch name {
	case "manhattan":
		return Manhattan, true
	case "euclidean":
		return Euclidean, true
	case "euclidean_squared":
		return EuclideanSquared, true
	case "diagonal":
		return Diagonal, true
	case "diagonal_cost":
		return DiagonalWithCost, true
	case "zero":
		return Zero, true
	default:
		return nil, false
	}
}

// GetSupportedHeuristics returns a list of all supported heuristic names.
//
// Returns:
//
//	[]string: list of heuristic names
func GetSupportedHeuristics() []string {
	return []string{
		"manhattan",
		"euclidean",
		"euclidean_squared",
		"diagonal",
		"diagonal_cost",
		"zero",
	}
}
