package algo

import (
	"testing"
)

// BenchmarkAStar_SmallGrid tests A* performance on small grids (100x100)
// Target: < 1ms pathfinding per docs/PLANNING.md
func BenchmarkAStar_SmallGrid(b *testing.B) {
	grid, err := NewGrid(100, 100, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(99, 99)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

// BenchmarkAStar_MediumGrid tests A* performance on medium grids (500x500)
// Target: < 10ms pathfinding per docs/PLANNING.md
func BenchmarkAStar_MediumGrid(b *testing.B) {
	grid, err := NewGrid(500, 500, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(499, 499)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

// BenchmarkAStar_LargeGrid tests A* performance on large grids (1000x1000)
// Target: < 100ms pathfinding per docs/PLANNING.md
func BenchmarkAStar_LargeGrid(b *testing.B) {
	grid, err := NewGrid(1000, 1000, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(999, 999)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

// BenchmarkAStar_WithObstacles tests performance with obstacles
func BenchmarkAStar_WithObstacles(b *testing.B) {
	grid, err := NewGrid(200, 200, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	// Add random obstacles (about 20% of the grid)
	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y++ {
			if (x*y)%5 == 0 && x != 0 && y != 0 && x != 199 && y != 199 {
				node, _ := grid.GetNode(x, y)
				node.IsObstacle = true
			}
		}
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(199, 199)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

// BenchmarkAStar_DifferentHeuristics compares heuristic performance
func BenchmarkAStar_Manhattan(b *testing.B) {
	grid, err := NewGrid(300, 300, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)
	astar.SetHeuristic(Manhattan)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(299, 299)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

func BenchmarkAStar_Euclidean(b *testing.B) {
	grid, err := NewGrid(300, 300, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)
	astar.SetHeuristic(Euclidean)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(299, 299)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

func BenchmarkAStar_Diagonal(b *testing.B) {
	grid, err := NewGrid(300, 300, EightWay) // Use 8-way for diagonal
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)
	astar.SetHeuristic(Diagonal)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(299, 299)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

// BenchmarkAStar_ShortPath tests performance on short paths
func BenchmarkAStar_ShortPath(b *testing.B) {
	grid, err := NewGrid(100, 100, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(45, 45)
	goal, _ := grid.GetNode(55, 55)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}

// BenchmarkAStar_MemoryAllocation tests memory allocation patterns
func BenchmarkAStar_MemoryAllocation(b *testing.B) {
	grid, err := NewGrid(200, 200, FourWay)
	if err != nil {
		b.Fatalf("Failed to create grid: %v", err)
	}

	astar := NewAStar()
	astar.SetGrid(grid)

	start, _ := grid.GetNode(0, 0)
	goal, _ := grid.GetNode(199, 199)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := astar.FindPath(start, goal)
		if err != nil {
			b.Fatalf("FindPath failed: %v", err)
		}
	}
}
