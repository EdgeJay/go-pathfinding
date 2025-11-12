package algo

import (
	"testing"
)

// TestAStar_NewAStar tests the constructor
func TestAStar_NewAStar(t *testing.T) {
	astar := NewAStar()

	if astar == nil {
		t.Fatal("NewAStar() returned nil")
	}

	if astar.openSet == nil {
		t.Error("openSet should be initialized")
	}

	if astar.closedSet == nil {
		t.Error("closedSet should be initialized")
	}

	if astar.heuristic == nil {
		t.Error("heuristic should be set to default (Manhattan)")
	}

	// Test that default heuristic works
	start := NewNode(0, 0)
	goal := NewNode(3, 4)
	distance := astar.heuristic(start, goal)
	expected := float64(7) // Manhattan distance: |3-0| + |4-0| = 7

	if distance != expected {
		t.Errorf("Default heuristic returned %f, expected %f", distance, expected)
	}
}

// TestAStar_SetGrid tests grid assignment
func TestAStar_SetGrid(t *testing.T) {
	astar := NewAStar()
	grid, err := NewGrid(10, 10, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}

	astar.SetGrid(grid)

	if astar.grid != grid {
		t.Error("SetGrid() did not set the grid correctly")
	}
}

// TestAStar_SetHeuristic tests heuristic assignment
func TestAStar_SetHeuristic(t *testing.T) {
	astar := NewAStar()

	// Test setting Euclidean heuristic
	astar.SetHeuristic(Euclidean)

	start := NewNode(0, 0)
	goal := NewNode(3, 4)
	distance := astar.heuristic(start, goal)
	expected := 5.0 // Euclidean distance: sqrt(3^2 + 4^2) = 5

	if distance != expected {
		t.Errorf("Euclidean heuristic returned %f, expected %f", distance, expected)
	}
}

// TestAStar_FindPath_SimpleCase tests basic pathfinding
func TestAStar_FindPath_SimpleCase(t *testing.T) {
	// Create a simple 5x5 grid with no obstacles
	grid, err := NewGrid(5, 5, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}
	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(4, 4)

	path, err := astar.FindPath(start, goal)

	if err != nil {
		t.Fatalf("FindPath() returned error: %v", err)
	}

	if len(path) == 0 {
		t.Fatal("FindPath() returned empty path")
	}

	// Check that path starts with start node and ends with goal node
	if !path[0].Equals(start) {
		t.Errorf("Path should start with start node, got %v", path[0])
	}

	if !path[len(path)-1].Equals(goal) {
		t.Errorf("Path should end with goal node, got %v", path[len(path)-1])
	}

	// For Manhattan heuristic on empty grid, optimal path length should be 9
	// (8 moves: 4 right + 4 down, plus start node = 9 nodes total)
	expectedLength := 9
	if len(path) != expectedLength {
		t.Errorf("Expected path length %d, got %d", expectedLength, len(path))
	}
}

// TestAStar_FindPath_WithObstacles tests pathfinding around obstacles
func TestAStar_FindPath_WithObstacles(t *testing.T) {
	// Create grid with obstacles forming a wall
	grid, err := NewGrid(5, 5, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}

	// Create vertical wall at x=2, blocking direct path
	for y := 1; y < 4; y++ {
		node, _ := grid.GetNode(2, y)
		node.IsObstacle = true
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 2)
	goal, _ := grid.GetNode(4, 2)

	path, err := astar.FindPath(start, goal)

	if err != nil {
		t.Fatalf("FindPath() returned error: %v", err)
	}

	if len(path) == 0 {
		t.Fatal("FindPath() returned empty path")
	}

	// Verify path doesn't go through obstacles
	for _, node := range path {
		if node.IsObstacle {
			t.Errorf("Path goes through obstacle at (%d, %d)", node.X, node.Y)
		}
	}

	// Path should be longer than direct path due to detour
	if len(path) < 6 { // At minimum: start, detour around wall, goal
		t.Errorf("Path seems too short for obstacle avoidance: %d nodes", len(path))
	}
}

// TestAStar_FindPath_NoPath tests when no path exists
func TestAStar_FindPath_NoPath(t *testing.T) {
	grid, err := NewGrid(5, 5, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}

	// Create complete wall separating start and goal
	for y := 0; y < 5; y++ {
		node, _ := grid.GetNode(2, y)
		node.IsObstacle = true
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 2)
	goal, _ := grid.GetNode(4, 2)

	path, err := astar.FindPath(start, goal)

	if err == nil {
		t.Fatal("FindPath() should return error when no path exists")
	}

	if path != nil {
		t.Error("FindPath() should return nil path when no path exists")
	}
}

// TestAStar_FindPath_SameStartGoal tests when start equals goal
func TestAStar_FindPath_SameStartGoal(t *testing.T) {
	grid, err := NewGrid(5, 5, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}
	astar := NewAStar()
	astar.SetGrid(grid)

	node, _ := grid.GetNode(2, 2)

	path, err := astar.FindPath(node, node)

	if err != nil {
		t.Fatalf("FindPath() returned error for same start/goal: %v", err)
	}

	if len(path) != 1 {
		t.Errorf("Expected path length 1 for same start/goal, got %d", len(path))
	}

	if !path[0].Equals(node) {
		t.Error("Single-node path should contain the start/goal node")
	}
}

// TestAStar_FindPath_InputValidation tests error handling for invalid inputs
func TestAStar_FindPath_InputValidation(t *testing.T) {
	astar := NewAStar()
	grid, err := NewGrid(5, 5, FourWay)
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}
	node, _ := grid.GetNode(0, 0)

	// Test nil start node
	_, err2 := astar.FindPath(nil, node)
	if err2 == nil {
		t.Error("FindPath() should return error for nil start node")
	}

	// Test nil goal node
	_, err2 = astar.FindPath(node, nil)
	if err2 == nil {
		t.Error("FindPath() should return error for nil goal node")
	}

	// Set grid for remaining tests
	astar.SetGrid(grid)

	// Test start node as obstacle
	start, _ := grid.GetNode(1, 1)
	goal, _ := grid.GetNode(2, 2)
	start.IsObstacle = true

	_, err2 = astar.FindPath(start, goal)
	if err2 == nil {
		t.Error("FindPath() should return error when start is obstacle")
	}

	// Test goal node as obstacle
	start.IsObstacle = false
	goal.IsObstacle = true

	_, err2 = astar.FindPath(start, goal)
	if err2 == nil {
		t.Error("FindPath() should return error when goal is obstacle")
	}
}

// TestAStar_FindPath_NoGrid tests error when grid is not set
func TestAStar_FindPath_NoGrid(t *testing.T) {
	astar := NewAStar()
	start := NewNode(0, 0)
	goal := NewNode(1, 1)

	_, err := astar.FindPath(start, goal)
	if err == nil {
		t.Error("FindPath() should return error when grid is not set")
	}
}
