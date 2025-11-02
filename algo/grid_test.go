package algo

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	tests := []struct {
		name         string
		width        int
		height       int
		movementType MovementType
		expectError  bool
	}{
		{"Valid 5x5 grid", 5, 5, FourWay, false},
		{"Valid 10x20 grid", 10, 20, EightWay, false},
		{"Invalid zero width", 0, 5, FourWay, true},
		{"Invalid zero height", 5, 0, FourWay, true},
		{"Invalid negative width", -1, 5, FourWay, true},
		{"Invalid negative height", 5, -1, FourWay, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, err := NewGrid(tt.width, tt.height, tt.movementType)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if grid.Width != tt.width || grid.Height != tt.height {
				t.Errorf("Expected dimensions %dx%d, got %dx%d", tt.width, tt.height, grid.Width, grid.Height)
			}

			if grid.MovementType != tt.movementType {
				t.Errorf("Expected movement type %v, got %v", tt.movementType, grid.MovementType)
			}

			// Check that all nodes are initialized
			for y := 0; y < tt.height; y++ {
				for x := 0; x < tt.width; x++ {
					node, err := grid.GetNode(x, y)
					if err != nil {
						t.Errorf("Error getting node at (%d, %d): %v", x, y, err)
					}
					if node.X != x || node.Y != y {
						t.Errorf("Node at (%d, %d) has wrong coordinates (%d, %d)", x, y, node.X, node.Y)
					}
					if node.IsObstacle {
						t.Errorf("New node at (%d, %d) should not be an obstacle", x, y)
					}
				}
			}
		})
	}
}

func TestGridGetNode(t *testing.T) {
	grid, _ := NewGrid(5, 5, FourWay)

	// Valid positions
	node, err := grid.GetNode(2, 3)
	if err != nil {
		t.Errorf("Unexpected error getting valid node: %v", err)
	}
	if node.X != 2 || node.Y != 3 {
		t.Errorf("Expected node at (2, 3), got (%d, %d)", node.X, node.Y)
	}

	// Invalid positions
	invalidPositions := [][2]int{
		{-1, 0}, {0, -1}, {5, 0}, {0, 5}, {-1, -1}, {5, 5},
	}

	for _, pos := range invalidPositions {
		_, err := grid.GetNode(pos[0], pos[1])
		if err == nil {
			t.Errorf("Expected error for position (%d, %d), but got none", pos[0], pos[1])
		}
	}
}

func TestGridObstacles(t *testing.T) {
	grid, _ := NewGrid(5, 5, FourWay)

	// Initially no obstacles
	if grid.IsObstacle(2, 2) {
		t.Error("New grid should have no obstacles")
	}

	// Set obstacle
	err := grid.SetObstacle(2, 2)
	if err != nil {
		t.Errorf("Unexpected error setting obstacle: %v", err)
	}

	if !grid.IsObstacle(2, 2) {
		t.Error("Position should be an obstacle after setting")
	}

	// Clear obstacle
	err = grid.ClearObstacle(2, 2)
	if err != nil {
		t.Errorf("Unexpected error clearing obstacle: %v", err)
	}

	if grid.IsObstacle(2, 2) {
		t.Error("Position should not be an obstacle after clearing")
	}

	// Test out of bounds
	if !grid.IsObstacle(-1, 0) {
		t.Error("Out of bounds positions should be considered obstacles")
	}

	// Test error cases
	err = grid.SetObstacle(-1, 0)
	if err == nil {
		t.Error("Expected error when setting obstacle out of bounds")
	}

	err = grid.ClearObstacle(10, 10)
	if err == nil {
		t.Error("Expected error when clearing obstacle out of bounds")
	}
}

func TestGridNeighbors(t *testing.T) {
	// Test 4-way movement
	grid4, _ := NewGrid(5, 5, FourWay)
	centerNode, _ := grid4.GetNode(2, 2)

	neighbors4 := grid4.GetNeighbors(centerNode)
	if len(neighbors4) != 4 {
		t.Errorf("Expected 4 neighbors for 4-way movement, got %d", len(neighbors4))
	}

	// Expected neighbor positions for (2,2) in 4-way
	expected4 := [][2]int{{2, 1}, {3, 2}, {2, 3}, {1, 2}}
	found := make(map[[2]int]bool)
	for _, neighbor := range neighbors4 {
		found[[2]int{neighbor.X, neighbor.Y}] = true
	}

	for _, exp := range expected4 {
		if !found[exp] {
			t.Errorf("Expected neighbor at (%d, %d) not found", exp[0], exp[1])
		}
	}

	// Test 8-way movement
	grid8, _ := NewGrid(5, 5, EightWay)
	centerNode8, _ := grid8.GetNode(2, 2)

	neighbors8 := grid8.GetNeighbors(centerNode8)
	if len(neighbors8) != 8 {
		t.Errorf("Expected 8 neighbors for 8-way movement, got %d", len(neighbors8))
	}

	// Test corner node (should have fewer neighbors)
	cornerNode, _ := grid4.GetNode(0, 0)
	cornerNeighbors := grid4.GetNeighbors(cornerNode)
	if len(cornerNeighbors) != 2 {
		t.Errorf("Expected 2 neighbors for corner node in 4-way, got %d", len(cornerNeighbors))
	}

	// Test with obstacles
	grid4.SetObstacle(2, 1) // Block one neighbor
	blockedNeighbors := grid4.GetNeighbors(centerNode)
	if len(blockedNeighbors) != 3 {
		t.Errorf("Expected 3 neighbors after blocking one, got %d", len(blockedNeighbors))
	}
}

func TestGridCost(t *testing.T) {
	grid, _ := NewGrid(5, 5, EightWay)

	node1, _ := grid.GetNode(1, 1)
	node2, _ := grid.GetNode(2, 1) // Horizontal neighbor
	node3, _ := grid.GetNode(2, 2) // Diagonal neighbor

	// Horizontal/vertical movement should cost 1.0
	cost1 := grid.GetCost(node1, node2)
	if cost1 != 1.0 {
		t.Errorf("Expected horizontal cost 1.0, got %.3f", cost1)
	}

	// Diagonal movement should cost ~1.414
	cost2 := grid.GetCost(node1, node3)
	if cost2 < 1.4 || cost2 > 1.42 {
		t.Errorf("Expected diagonal cost ~1.414, got %.3f", cost2)
	}

	// Test with custom node cost
	node3.Cost = 2.0
	cost3 := grid.GetCost(node1, node3)
	if cost3 < 2.8 || cost3 > 2.85 {
		t.Errorf("Expected weighted diagonal cost ~2.828, got %.3f", cost3)
	}
}

func TestGridReset(t *testing.T) {
	grid, _ := NewGrid(3, 3, FourWay)

	// Set some A* state
	node, _ := grid.GetNode(1, 1)
	parent, _ := grid.GetNode(0, 0)
	node.G = 5.0
	node.H = 3.0
	node.F = 8.0
	node.Parent = parent

	// Reset grid
	grid.Reset()

	// Check that A* state is cleared
	if node.G != 0 || node.H != 0 || node.F != 0 || node.Parent != nil {
		t.Error("Grid reset should clear A* state from all nodes")
	}

	// But position should be preserved
	if node.X != 1 || node.Y != 1 {
		t.Error("Grid reset should preserve node positions")
	}
}

func TestGridString(t *testing.T) {
	grid, _ := NewGrid(3, 2, FourWay)
	grid.SetObstacle(1, 0)
	grid.SetObstacle(2, 1)

	str := grid.String()

	// Should contain grid info and representation
	if len(str) == 0 {
		t.Error("Grid string should not be empty")
	}

	// Should show obstacles as # and empty as .
	// Expected pattern: ".#." on first row, "..#" on second row
	expectedPattern := ".#.\n..#\n"
	if !containsPattern(str, expectedPattern) {
		t.Errorf("Grid string should contain pattern %q, got %q", expectedPattern, str)
	}
}

// Helper function to check if string contains pattern
func containsPattern(s, pattern string) bool {
	// Simple substring check for this test
	for i := 0; i <= len(s)-len(pattern); i++ {
		if s[i:i+len(pattern)] == pattern {
			return true
		}
	}
	return false
}

// Benchmarks
func BenchmarkNewGrid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewGrid(50, 50, FourWay)
	}
}

func BenchmarkGetNeighbors4Way(b *testing.B) {
	grid, _ := NewGrid(100, 100, FourWay)
	node, _ := grid.GetNode(50, 50)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = grid.GetNeighbors(node)
	}
}

func BenchmarkGetNeighbors8Way(b *testing.B) {
	grid, _ := NewGrid(100, 100, EightWay)
	node, _ := grid.GetNode(50, 50)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = grid.GetNeighbors(node)
	}
}
