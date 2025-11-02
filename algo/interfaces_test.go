package algo

import (
	"testing"
)

// Test that Grid implements GridInterface
func TestGridInterfaceCompliance(t *testing.T) {
	grid, err := NewGrid(5, 5, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}

	// This line will cause a compile error if Grid doesn't implement GridInterface
	var _ GridInterface = grid

	// Test that we can use Grid through the interface
	var gridInterface GridInterface = grid

	// Test interface methods
	node, err := gridInterface.GetNode(2, 2)
	if err != nil {
		t.Errorf("Interface GetNode failed: %v", err)
	}

	neighbors := gridInterface.GetNeighbors(node)
	if len(neighbors) == 0 {
		t.Error("Interface GetNeighbors returned no neighbors")
	}

	isObstacle := gridInterface.IsObstacle(2, 2)
	if isObstacle {
		t.Error("New grid position should not be obstacle")
	}

	neighbor := neighbors[0]
	cost := gridInterface.GetCost(node, neighbor)
	if cost <= 0 {
		t.Error("Interface GetCost should return positive cost")
	}

	// Test reset through interface
	gridInterface.Reset()
}

// Test heuristic function type
func TestHeuristicFunc(t *testing.T) {
	// Create a simple Manhattan distance heuristic for testing
	manhattan := func(current, goal *Node) float64 {
		dx := float64(current.X - goal.X)
		dy := float64(current.Y - goal.Y)
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		return dx + dy
	}

	// Test the heuristic function
	node1 := NewNode(0, 0)
	node2 := NewNode(3, 4)

	distance := manhattan(node1, node2)
	expected := 7.0 // |3-0| + |4-0| = 7

	if distance != expected {
		t.Errorf("Expected Manhattan distance %.1f, got %.1f", expected, distance)
	}
}

// Example of how interfaces will be used together
func TestInterfaceUsage(t *testing.T) {
	// Create a grid
	grid, _ := NewGrid(5, 5, FourWay)

	// Add some obstacles
	grid.SetObstacle(2, 1)
	grid.SetObstacle(2, 2)
	grid.SetObstacle(2, 3)

	// Get nodes
	start, _ := grid.GetNode(1, 2)
	goal, _ := grid.GetNode(3, 2)

	// Define a heuristic
	euclidean := func(current, goal *Node) float64 {
		dx := float64(current.X - goal.X)
		dy := float64(current.Y - goal.Y)
		return dx*dx + dy*dy // Squared Euclidean for efficiency
	}

	// Test heuristic calculation
	h := euclidean(start, goal)
	if h != 4.0 { // (3-1)^2 + (2-2)^2 = 4
		t.Errorf("Expected heuristic 4.0, got %.1f", h)
	}

	// This demonstrates how the interfaces will work together
	// when we implement the actual A* algorithm
	var gridInterface GridInterface = grid
	neighbors := gridInterface.GetNeighbors(start)

	// Should have neighbors despite obstacles
	if len(neighbors) == 0 {
		t.Error("Start node should have some accessible neighbors")
	}
}

// Test interface with different movement types
func TestInterfaceMovementTypes(t *testing.T) {
	tests := []struct {
		name         string
		movementType MovementType
		expectedMin  int // Minimum neighbors for center node
		expectedMax  int // Maximum neighbors for center node
	}{
		{"4-way movement", FourWay, 4, 4},
		{"8-way movement", EightWay, 8, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, _ := NewGrid(5, 5, tt.movementType)
			var gridInterface GridInterface = grid

			centerNode, _ := gridInterface.GetNode(2, 2)
			neighbors := gridInterface.GetNeighbors(centerNode)

			if len(neighbors) < tt.expectedMin || len(neighbors) > tt.expectedMax {
				t.Errorf("Expected %d-%d neighbors, got %d", tt.expectedMin, tt.expectedMax, len(neighbors))
			}
		})
	}
}
