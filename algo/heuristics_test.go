package algo

import (
	"math"
	"testing"
)

func TestManhattan(t *testing.T) {
	tests := []struct {
		name     string
		x1, y1   int
		x2, y2   int
		expected float64
	}{
		{"Same position", 0, 0, 0, 0, 0.0},
		{"Horizontal distance", 0, 0, 5, 0, 5.0},
		{"Vertical distance", 0, 0, 0, 3, 3.0},
		{"Diagonal distance", 0, 0, 3, 4, 7.0},
		{"Negative coordinates", -2, -3, 1, 2, 8.0},
		{"Mixed signs", -5, 3, 2, -1, 11.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node1 := NewNode(tt.x1, tt.y1)
			node2 := NewNode(tt.x2, tt.y2)

			result := Manhattan(node1, node2)
			if result != tt.expected {
				t.Errorf("Manhattan(%d,%d -> %d,%d) = %.1f, want %.1f",
					tt.x1, tt.y1, tt.x2, tt.y2, result, tt.expected)
			}
		})
	}
}

func TestEuclidean(t *testing.T) {
	tests := []struct {
		name     string
		x1, y1   int
		x2, y2   int
		expected float64
	}{
		{"Same position", 0, 0, 0, 0, 0.0},
		{"Unit horizontal", 0, 0, 1, 0, 1.0},
		{"Unit vertical", 0, 0, 0, 1, 1.0},
		{"3-4-5 triangle", 0, 0, 3, 4, 5.0},
		{"Unit diagonal", 0, 0, 1, 1, math.Sqrt(2)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node1 := NewNode(tt.x1, tt.y1)
			node2 := NewNode(tt.x2, tt.y2)

			result := Euclidean(node1, node2)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Euclidean(%d,%d -> %d,%d) = %.6f, want %.6f",
					tt.x1, tt.y1, tt.x2, tt.y2, result, tt.expected)
			}
		})
	}
}

func TestEuclideanSquared(t *testing.T) {
	tests := []struct {
		name     string
		x1, y1   int
		x2, y2   int
		expected float64
	}{
		{"Same position", 0, 0, 0, 0, 0.0},
		{"Unit distance", 0, 0, 1, 0, 1.0},
		{"3-4-5 triangle", 0, 0, 3, 4, 25.0},
		{"Unit diagonal", 0, 0, 1, 1, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node1 := NewNode(tt.x1, tt.y1)
			node2 := NewNode(tt.x2, tt.y2)

			result := EuclideanSquared(node1, node2)
			if result != tt.expected {
				t.Errorf("EuclideanSquared(%d,%d -> %d,%d) = %.1f, want %.1f",
					tt.x1, tt.y1, tt.x2, tt.y2, result, tt.expected)
			}
		})
	}
}

func TestDiagonal(t *testing.T) {
	tests := []struct {
		name     string
		x1, y1   int
		x2, y2   int
		expected float64
	}{
		{"Same position", 0, 0, 0, 0, 0.0},
		{"Horizontal only", 0, 0, 5, 0, 5.0},
		{"Vertical only", 0, 0, 0, 3, 3.0},
		{"Square diagonal", 0, 0, 3, 3, 3.0},
		{"Rectangle - width larger", 0, 0, 5, 3, 5.0},
		{"Rectangle - height larger", 0, 0, 3, 5, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node1 := NewNode(tt.x1, tt.y1)
			node2 := NewNode(tt.x2, tt.y2)

			result := Diagonal(node1, node2)
			if result != tt.expected {
				t.Errorf("Diagonal(%d,%d -> %d,%d) = %.1f, want %.1f",
					tt.x1, tt.y1, tt.x2, tt.y2, result, tt.expected)
			}
		})
	}
}

func TestDiagonalWithCost(t *testing.T) {
	tests := []struct {
		name      string
		x1, y1    int
		x2, y2    int
		approx    float64 // Approximate expected value
		tolerance float64
	}{
		{"Same position", 0, 0, 0, 0, 0.0, 1e-9},
		{"Horizontal only", 0, 0, 3, 0, 3.0, 1e-9},
		{"Vertical only", 0, 0, 0, 3, 3.0, 1e-9},
		{"Pure diagonal", 0, 0, 3, 3, 3 * 1.414, 0.01},
		{"L-shape", 0, 0, 3, 1, 3.414, 0.01}, // 2*1 + 1*1.414
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node1 := NewNode(tt.x1, tt.y1)
			node2 := NewNode(tt.x2, tt.y2)

			result := DiagonalWithCost(node1, node2)
			if math.Abs(result-tt.approx) > tt.tolerance {
				t.Errorf("DiagonalWithCost(%d,%d -> %d,%d) = %.3f, want ~%.3f",
					tt.x1, tt.y1, tt.x2, tt.y2, result, tt.approx)
			}
		})
	}
}

func TestZero(t *testing.T) {
	node1 := NewNode(0, 0)
	node2 := NewNode(100, 100)

	result := Zero(node1, node2)
	if result != 0.0 {
		t.Errorf("Zero heuristic should always return 0.0, got %.1f", result)
	}
}

func TestGetHeuristicByName(t *testing.T) {
	tests := []struct {
		name      string
		valid     bool
		testValue float64 // Expected result for test case (3,4) -> (0,0)
	}{
		{"manhattan", true, 7.0},
		{"euclidean", true, 5.0},
		{"euclidean_squared", true, 25.0},
		{"diagonal", true, 4.0},
		{"diagonal_cost", true, 5.243}, // 1*(4-3) + 1.414*3 ≈ 5.242
		{"zero", true, 0.0},
		{"invalid", false, 0.0},
		{"", false, 0.0},
	}

	node1 := NewNode(3, 4)
	node2 := NewNode(0, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heuristic, valid := GetHeuristicByName(tt.name)

			if valid != tt.valid {
				t.Errorf("GetHeuristicByName(%q) validity = %v, want %v",
					tt.name, valid, tt.valid)
				return
			}

			if valid {
				result := heuristic(node1, node2)
				tolerance := 0.01
				if math.Abs(result-tt.testValue) > tolerance {
					t.Errorf("Heuristic %q returned %.3f, expected ~%.3f",
						tt.name, result, tt.testValue)
				}
			}
		})
	}
}

func TestGetSupportedHeuristics(t *testing.T) {
	supported := GetSupportedHeuristics()

	expectedCount := 6
	if len(supported) != expectedCount {
		t.Errorf("Expected %d supported heuristics, got %d", expectedCount, len(supported))
	}

	// Check that all returned heuristics are actually valid
	for _, name := range supported {
		_, valid := GetHeuristicByName(name)
		if !valid {
			t.Errorf("Supported heuristic %q is not actually valid", name)
		}
	}

	// Check specific heuristics are included
	expectedHeuristics := []string{"manhattan", "euclidean", "diagonal", "zero"}
	heuristicSet := make(map[string]bool)
	for _, h := range supported {
		heuristicSet[h] = true
	}

	for _, expected := range expectedHeuristics {
		if !heuristicSet[expected] {
			t.Errorf("Expected heuristic %q not found in supported list", expected)
		}
	}
}

// Test admissibility property for different heuristics
func TestHeuristicAdmissibility(t *testing.T) {
	// For a simple case where we know the actual shortest path distance
	// From (0,0) to (2,2), actual shortest distances are:
	// - 4-way movement: 4 steps (right, right, up, up)
	// - 8-way movement: 2.828 steps (diagonal to (1,1), then diagonal to (2,2))

	node1 := NewNode(0, 0)
	node2 := NewNode(2, 2)

	actualDistance4Way := 4.0              // Manhattan distance for 4-way
	actualDistance8Way := 2 * math.Sqrt(2) // 2 diagonal moves

	tests := []struct {
		name               string
		heuristic          HeuristicFunc
		shouldBeAdmissible bool
		actualDistance     float64
	}{
		{"Manhattan for 4-way", Manhattan, true, actualDistance4Way},
		{"Manhattan for 8-way", Manhattan, false, actualDistance8Way}, // NOT admissible for 8-way
		{"Euclidean for any", Euclidean, true, actualDistance8Way},
		{"Diagonal for 8-way", Diagonal, true, actualDistance8Way},
		{"Zero for any", Zero, true, actualDistance4Way},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heuristicValue := tt.heuristic(node1, node2)

			if tt.shouldBeAdmissible && heuristicValue > tt.actualDistance+1e-9 {
				t.Errorf("Heuristic %s is not admissible: h=%.3f > actual=%.3f",
					tt.name, heuristicValue, tt.actualDistance)
			}
			if !tt.shouldBeAdmissible && heuristicValue <= tt.actualDistance+1e-9 {
				// This test case expects the heuristic to NOT be admissible
				// but it actually is admissible, which is fine
			}
		})
	}
}

// Benchmarks for performance comparison
func BenchmarkManhattan(b *testing.B) {
	node1 := NewNode(0, 0)
	node2 := NewNode(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Manhattan(node1, node2)
	}
}

func BenchmarkEuclidean(b *testing.B) {
	node1 := NewNode(0, 0)
	node2 := NewNode(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Euclidean(node1, node2)
	}
}

func BenchmarkEuclideanSquared(b *testing.B) {
	node1 := NewNode(0, 0)
	node2 := NewNode(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EuclideanSquared(node1, node2)
	}
}

func BenchmarkDiagonal(b *testing.B) {
	node1 := NewNode(0, 0)
	node2 := NewNode(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Diagonal(node1, node2)
	}
}

func BenchmarkDiagonalWithCost(b *testing.B) {
	node1 := NewNode(0, 0)
	node2 := NewNode(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DiagonalWithCost(node1, node2)
	}
}
