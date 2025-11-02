package algo

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(5, 10)

	if node.X != 5 || node.Y != 10 {
		t.Errorf("Expected position (5, 10), got (%d, %d)", node.X, node.Y)
	}

	if node.G != 0 || node.H != 0 || node.F != 0 {
		t.Errorf("Expected costs to be zero, got G=%.2f, H=%.2f, F=%.2f", node.G, node.H, node.F)
	}

	if node.Parent != nil {
		t.Error("Expected parent to be nil")
	}

	if node.IsObstacle {
		t.Error("Expected IsObstacle to be false")
	}

	if node.Cost != 1.0 {
		t.Errorf("Expected cost to be 1.0, got %.2f", node.Cost)
	}
}

func TestNewObstacle(t *testing.T) {
	obstacle := NewObstacle(3, 7)

	if obstacle.X != 3 || obstacle.Y != 7 {
		t.Errorf("Expected position (3, 7), got (%d, %d)", obstacle.X, obstacle.Y)
	}

	if !obstacle.IsObstacle {
		t.Error("Expected IsObstacle to be true")
	}

	if obstacle.Cost != 1.0 {
		t.Errorf("Expected cost to be 1.0, got %.2f", obstacle.Cost)
	}
}

func TestNodeCalculateF(t *testing.T) {
	node := NewNode(0, 0)
	node.G = 5.5
	node.H = 3.2

	node.CalculateF()

	expected := 8.7
	if node.F != expected {
		t.Errorf("Expected F to be %.1f, got %.1f", expected, node.F)
	}
}

func TestNodeReset(t *testing.T) {
	node := NewNode(2, 4)
	parent := NewNode(1, 3)

	// Set some A* state
	node.G = 10
	node.H = 5
	node.F = 15
	node.Parent = parent

	// Reset should clear A* state but preserve position and grid properties
	node.Reset()

	if node.X != 2 || node.Y != 4 {
		t.Errorf("Reset should preserve position, got (%d, %d)", node.X, node.Y)
	}

	if node.G != 0 || node.H != 0 || node.F != 0 {
		t.Errorf("Reset should clear costs, got G=%.1f, H=%.1f, F=%.1f", node.G, node.H, node.F)
	}

	if node.Parent != nil {
		t.Error("Reset should clear parent")
	}

	// Grid properties should be preserved
	if node.Cost != 1.0 {
		t.Errorf("Reset should preserve cost, got %.1f", node.Cost)
	}
}

func TestNodeEquals(t *testing.T) {
	node1 := NewNode(5, 10)
	node2 := NewNode(5, 10)
	node3 := NewNode(3, 8)

	if !node1.Equals(node2) {
		t.Error("Nodes with same position should be equal")
	}

	if node1.Equals(node3) {
		t.Error("Nodes with different positions should not be equal")
	}

	if node1.Equals(nil) {
		t.Error("Node should not equal nil")
	}
}

func TestNodeString(t *testing.T) {
	node := NewNode(5, 10)
	node.G = 2.5
	node.H = 3.7
	node.F = 6.2

	str := node.String()
	expected := "Node(5,10) [g=2.50, h=3.70, f=6.20] obstacle=false"

	if str != expected {
		t.Errorf("Expected string %q, got %q", expected, str)
	}
}

// Benchmark for node creation
func BenchmarkNewNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewNode(i%100, i%100)
	}
}

// Benchmark for node operations
func BenchmarkNodeCalculateF(b *testing.B) {
	node := NewNode(0, 0)
	for i := 0; i < b.N; i++ {
		node.G = float64(i)
		node.H = float64(i * 2)
		node.CalculateF()
	}
}
