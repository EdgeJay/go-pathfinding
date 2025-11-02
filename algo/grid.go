package algo

import (
	"errors"
	"fmt"
)

// MovementType defines the type of movement allowed in the grid
type MovementType int

const (
	// FourWay allows movement in 4 directions: up, down, left, right
	FourWay MovementType = iota
	// EightWay allows movement in 8 directions: 4-way + diagonals
	EightWay
)

// Grid represents a 2D pathfinding grid containing nodes.
// It provides methods for accessing nodes, checking boundaries,
// and getting neighbors for pathfinding algorithms.
type Grid struct {
	// Grid dimensions
	Width, Height int

	// 2D array of nodes stored as grid[y][x] (row-major order)
	nodes [][]*Node

	// Movement configuration
	MovementType MovementType
}

// NewGrid creates a new grid with the specified dimensions.
// All nodes are initialized as empty (non-obstacle) nodes.
//
// Params:
//
//	width, height: grid dimensions (must be > 0)
//	movementType: type of movement allowed (FourWay or EightWay)
//
// Returns:
//
//	*Grid: new grid instance
//	error: if dimensions are invalid
func NewGrid(width, height int, movementType MovementType) (*Grid, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("grid dimensions must be positive")
	}

	grid := &Grid{
		Width:        width,
		Height:       height,
		MovementType: movementType,
		nodes:        make([][]*Node, height),
	}

	// Initialize all nodes
	for y := 0; y < height; y++ {
		grid.nodes[y] = make([]*Node, width)
		for x := 0; x < width; x++ {
			grid.nodes[y][x] = NewNode(x, y)
		}
	}

	return grid, nil
}

// GetNode returns the node at the specified coordinates.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	*Node: node at the specified position
//	error: if coordinates are out of bounds
func (g *Grid) GetNode(x, y int) (*Node, error) {
	if !g.IsValidPosition(x, y) {
		return nil, fmt.Errorf("position (%d, %d) is out of bounds for grid %dx%d", x, y, g.Width, g.Height)
	}
	return g.nodes[y][x], nil
}

// SetObstacle marks a node as an obstacle.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	error: if coordinates are out of bounds
func (g *Grid) SetObstacle(x, y int) error {
	if !g.IsValidPosition(x, y) {
		return fmt.Errorf("position (%d, %d) is out of bounds for grid %dx%d", x, y, g.Width, g.Height)
	}
	g.nodes[y][x].IsObstacle = true
	return nil
}

// ClearObstacle removes obstacle status from a node.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	error: if coordinates are out of bounds
func (g *Grid) ClearObstacle(x, y int) error {
	if !g.IsValidPosition(x, y) {
		return fmt.Errorf("position (%d, %d) is out of bounds for grid %dx%d", x, y, g.Width, g.Height)
	}
	g.nodes[y][x].IsObstacle = false
	return nil
}

// IsObstacle checks if a position contains an obstacle.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	bool: true if position is an obstacle or out of bounds
func (g *Grid) IsObstacle(x, y int) bool {
	if !g.IsValidPosition(x, y) {
		return true // Out of bounds considered as obstacle
	}
	return g.nodes[y][x].IsObstacle
}

// IsValidPosition checks if coordinates are within grid bounds.
//
// Params:
//
//	x, y: grid coordinates
//
// Returns:
//
//	bool: true if position is within bounds
func (g *Grid) IsValidPosition(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

// GetNeighbors returns all valid neighbors of a node based on movement type.
// Neighbors are filtered to exclude obstacles and out-of-bounds positions.
//
// Params:
//
//	node: the node to get neighbors for
//
// Returns:
//
//	[]*Node: slice of neighboring nodes
func (g *Grid) GetNeighbors(node *Node) []*Node {
	neighbors := make([]*Node, 0, 8) // Pre-allocate for max 8 neighbors

	directions := g.getDirections()

	for _, dir := range directions {
		newX := node.X + dir[0]
		newY := node.Y + dir[1]

		if g.IsValidPosition(newX, newY) && !g.nodes[newY][newX].IsObstacle {
			neighbors = append(neighbors, g.nodes[newY][newX])
		}
	}

	return neighbors
}

// GetCost returns the movement cost from one node to another.
// For now, this returns the base movement cost, but can be extended
// for weighted grids or terrain-based costs.
//
// Params:
//
//	from, to: source and destination nodes
//
// Returns:
//
//	float64: movement cost
func (g *Grid) GetCost(from, to *Node) float64 {
	// Calculate Euclidean distance for cost
	dx := float64(to.X - from.X)
	dy := float64(to.Y - from.Y)

	// Apply terrain cost multiplier
	baseCost := 1.0
	if dx != 0 && dy != 0 {
		// Diagonal movement costs sqrt(2) ≈ 1.414
		baseCost = 1.414
	}

	return baseCost * to.Cost
}

// Reset clears all A* algorithm state from all nodes in the grid.
// This allows reusing the grid for multiple pathfinding operations.
func (g *Grid) Reset() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.nodes[y][x].Reset()
		}
	}
}

// getDirections returns the movement directions based on movement type.
// Returns slice of [dx, dy] offsets.
func (g *Grid) getDirections() [][2]int {
	// 4-way movement: up, right, down, left
	fourWay := [][2]int{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}

	if g.MovementType == FourWay {
		return fourWay
	}

	// 8-way movement: 4-way + diagonals
	return [][2]int{
		{0, -1},  // Up
		{1, -1},  // Up-Right
		{1, 0},   // Right
		{1, 1},   // Down-Right
		{0, 1},   // Down
		{-1, 1},  // Down-Left
		{-1, 0},  // Left
		{-1, -1}, // Up-Left
	}
}

// String returns a string representation of the grid for debugging.
// Obstacles are shown as '#', empty cells as '.'.
func (g *Grid) String() string {
	result := fmt.Sprintf("Grid %dx%d (%s movement):\n", g.Width, g.Height, g.movementTypeString())

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.nodes[y][x].IsObstacle {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}

// movementTypeString returns a string representation of the movement type
func (g *Grid) movementTypeString() string {
	switch g.MovementType {
	case FourWay:
		return "4-way"
	case EightWay:
		return "8-way"
	default:
		return "unknown"
	}
}
