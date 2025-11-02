package algo

import "fmt"

// Node represents a single cell in the pathfinding grid.
// It contains position coordinates, A* algorithm costs (g, h, f),
// and parent tracking for path reconstruction.
type Node struct {
	// Position coordinates in the grid
	X, Y int

	// A* algorithm costs
	G float64 // Cost from start node to this node
	H float64 // Heuristic cost from this node to goal
	F float64 // Total cost (G + H)

	// Parent node for path reconstruction
	Parent *Node

	// Grid properties
	IsObstacle bool    // Whether this node is an obstacle
	Cost       float64 // Movement cost multiplier for weighted grids (default: 1.0)
}

// NewNode creates a new node at the specified coordinates.
// By default, the node is not an obstacle and has a movement cost of 1.0.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	*Node: new node instance
func NewNode(x, y int) *Node {
	return &Node{
		X:          x,
		Y:          y,
		G:          0,
		H:          0,
		F:          0,
		Parent:     nil,
		IsObstacle: false,
		Cost:       1.0,
	}
}

// NewObstacle creates a new obstacle node at the specified coordinates.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	*Node: new obstacle node instance
func NewObstacle(x, y int) *Node {
	node := NewNode(x, y)
	node.IsObstacle = true
	return node
}

// CalculateF updates the F cost based on current G and H values.
// F = G + H represents the total estimated cost of the path through this node.
func (n *Node) CalculateF() {
	n.F = n.G + n.H
}

// Reset clears the A* algorithm state while preserving position and grid properties.
// This allows reusing the node for multiple pathfinding operations.
func (n *Node) Reset() {
	n.G = 0
	n.H = 0
	n.F = 0
	n.Parent = nil
}

// Equals checks if two nodes have the same position coordinates.
// Used for comparing nodes in pathfinding algorithms.
//
// Params:
//
//	other: node to compare with
//
// Returns:
//
//	bool: true if positions match
func (n *Node) Equals(other *Node) bool {
	if other == nil {
		return false
	}
	return n.X == other.X && n.Y == other.Y
}

// String returns a string representation of the node for debugging.
// Format: "Node(x,y) [g=G, h=H, f=F] obstacle=IsObstacle"
func (n *Node) String() string {
	return fmt.Sprintf("Node(%d,%d) [g=%.2f, h=%.2f, f=%.2f] obstacle=%t",
		n.X, n.Y, n.G, n.H, n.F, n.IsObstacle)
}
