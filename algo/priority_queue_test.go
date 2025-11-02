package algo

import (
	"container/heap"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	if pq == nil {
		t.Fatal("NewPriorityQueue returned nil")
	}

	if !pq.IsEmpty() {
		t.Error("New priority queue should be empty")
	}

	if pq.Len() != 0 {
		t.Errorf("New priority queue length should be 0, got %d", pq.Len())
	}
}

func TestPriorityQueueBasicOperations(t *testing.T) {
	pq := NewPriorityQueue()

	// Test empty queue operations
	if pq.PopNode() != nil {
		t.Error("PopNode on empty queue should return nil")
	}

	if pq.Peek() != nil {
		t.Error("Peek on empty queue should return nil")
	}

	// Add nodes with different F costs
	node1 := NewNode(1, 1)
	node1.F = 10.0

	node2 := NewNode(2, 2)
	node2.F = 5.0

	node3 := NewNode(3, 3)
	node3.F = 15.0

	pq.PushNode(node1)
	pq.PushNode(node2)
	pq.PushNode(node3)

	if pq.Len() != 3 {
		t.Errorf("Expected length 3, got %d", pq.Len())
	}

	if pq.IsEmpty() {
		t.Error("Queue should not be empty after adding nodes")
	}

	// Peek should return the node with lowest F cost
	peeked := pq.Peek()
	if peeked != node2 {
		t.Errorf("Peek should return node with F=5.0, got F=%.1f", peeked.F)
	}

	// Pop should return nodes in ascending F cost order
	popped1 := pq.PopNode()
	if popped1 != node2 || popped1.F != 5.0 {
		t.Errorf("First pop should return node with F=5.0, got F=%.1f", popped1.F)
	}

	popped2 := pq.PopNode()
	if popped2 != node1 || popped2.F != 10.0 {
		t.Errorf("Second pop should return node with F=10.0, got F=%.1f", popped2.F)
	}

	popped3 := pq.PopNode()
	if popped3 != node3 || popped3.F != 15.0 {
		t.Errorf("Third pop should return node with F=15.0, got F=%.1f", popped3.F)
	}

	if !pq.IsEmpty() {
		t.Error("Queue should be empty after popping all nodes")
	}
}

func TestPriorityQueueContains(t *testing.T) {
	pq := NewPriorityQueue()

	node1 := NewNode(5, 10)
	node1.F = 8.0

	// Check contains before adding
	if pq.Contains(5, 10) {
		t.Error("Queue should not contain node before adding")
	}

	if pq.ContainsNode(node1) {
		t.Error("Queue should not contain node before adding")
	}

	// Add node and check contains
	pq.PushNode(node1)

	if !pq.Contains(5, 10) {
		t.Error("Queue should contain node after adding")
	}

	if !pq.ContainsNode(node1) {
		t.Error("Queue should contain node after adding")
	}

	// Check non-existent coordinates
	if pq.Contains(1, 1) {
		t.Error("Queue should not contain non-existent coordinates")
	}

	// Remove node and check contains
	pq.PopNode()

	if pq.Contains(5, 10) {
		t.Error("Queue should not contain node after removing")
	}
}

func TestPriorityQueueGetNode(t *testing.T) {
	pq := NewPriorityQueue()

	node1 := NewNode(3, 7)
	node1.F = 12.0

	// Get from empty queue
	if pq.GetNode(3, 7) != nil {
		t.Error("GetNode should return nil for empty queue")
	}

	// Add node and get it
	pq.PushNode(node1)

	retrieved := pq.GetNode(3, 7)
	if retrieved != node1 {
		t.Error("GetNode should return the added node")
	}

	// Get non-existent node
	if pq.GetNode(1, 1) != nil {
		t.Error("GetNode should return nil for non-existent coordinates")
	}
}

func TestPriorityQueueDuplicateHandling(t *testing.T) {
	pq := NewPriorityQueue()

	// Add initial node
	node1 := NewNode(2, 3)
	node1.F = 10.0
	node1.G = 5.0
	node1.H = 5.0
	pq.PushNode(node1)

	// Add node at same position with higher F cost (should be ignored)
	node2 := NewNode(2, 3)
	node2.F = 15.0
	node2.G = 8.0
	node2.H = 7.0
	pq.PushNode(node2)

	if pq.Len() != 1 {
		t.Errorf("Queue should still have 1 node after adding duplicate with higher F, got %d", pq.Len())
	}

	retrieved := pq.GetNode(2, 3)
	if retrieved.F != 10.0 {
		t.Errorf("Original node should be preserved, F=%.1f", retrieved.F)
	}

	// Add node at same position with lower F cost (should update)
	node3 := NewNode(2, 3)
	node3.F = 7.0
	node3.G = 3.0
	node3.H = 4.0
	parent := NewNode(1, 2)
	node3.Parent = parent
	pq.PushNode(node3)

	if pq.Len() != 1 {
		t.Errorf("Queue should still have 1 node after updating, got %d", pq.Len())
	}

	retrieved = pq.GetNode(2, 3)
	if retrieved.F != 7.0 || retrieved.G != 3.0 || retrieved.H != 4.0 {
		t.Errorf("Node should be updated: F=%.1f, G=%.1f, H=%.1f", retrieved.F, retrieved.G, retrieved.H)
	}

	if retrieved.Parent != parent {
		t.Error("Parent should be updated")
	}
}

func TestPriorityQueueUpdatePriority(t *testing.T) {
	pq := NewPriorityQueue()

	node1 := NewNode(1, 1)
	node1.F = 10.0

	node2 := NewNode(2, 2)
	node2.F = 5.0

	pq.PushNode(node1)
	pq.PushNode(node2)

	// Update priority of existing node
	updated := pq.UpdatePriority(1, 1, 3.0)
	if !updated {
		t.Error("UpdatePriority should return true for existing node")
	}

	// Check that node1 now comes first (lower F cost)
	peeked := pq.Peek()
	if peeked.X != 1 || peeked.Y != 1 || peeked.F != 3.0 {
		t.Errorf("Updated node should come first: (%d,%d) F=%.1f", peeked.X, peeked.Y, peeked.F)
	}

	// Try to update non-existent node
	updated = pq.UpdatePriority(99, 99, 1.0)
	if updated {
		t.Error("UpdatePriority should return false for non-existent node")
	}
}

func TestPriorityQueueClear(t *testing.T) {
	pq := NewPriorityQueue()

	// Add several nodes
	for i := 0; i < 5; i++ {
		node := NewNode(i, i)
		node.F = float64(i)
		pq.PushNode(node)
	}

	if pq.Len() != 5 {
		t.Errorf("Expected 5 nodes, got %d", pq.Len())
	}

	// Clear the queue
	pq.Clear()

	if !pq.IsEmpty() {
		t.Error("Queue should be empty after clear")
	}

	if pq.Len() != 0 {
		t.Errorf("Queue length should be 0 after clear, got %d", pq.Len())
	}

	// Verify all lookups return nil
	for i := 0; i < 5; i++ {
		if pq.Contains(i, i) {
			t.Errorf("Queue should not contain (%d,%d) after clear", i, i)
		}
	}
}

func TestPriorityQueueHeapProperty(t *testing.T) {
	pq := NewPriorityQueue()

	// Add nodes in random order
	fCosts := []float64{15.0, 3.0, 8.0, 1.0, 12.0, 6.0, 9.0, 2.0}

	for i, f := range fCosts {
		node := NewNode(i, i)
		node.F = f
		pq.PushNode(node)
	}

	// Pop all nodes and verify they come out in ascending F cost order
	var poppedFCosts []float64
	for !pq.IsEmpty() {
		node := pq.PopNode()
		poppedFCosts = append(poppedFCosts, node.F)
	}

	// Verify ascending order
	for i := 1; i < len(poppedFCosts); i++ {
		if poppedFCosts[i] < poppedFCosts[i-1] {
			t.Errorf("F costs not in ascending order: %.1f should not come after %.1f",
				poppedFCosts[i], poppedFCosts[i-1])
		}
	}

	expectedOrder := []float64{1.0, 2.0, 3.0, 6.0, 8.0, 9.0, 12.0, 15.0}
	if len(poppedFCosts) != len(expectedOrder) {
		t.Errorf("Expected %d nodes, got %d", len(expectedOrder), len(poppedFCosts))
	}

	for i, expected := range expectedOrder {
		if i < len(poppedFCosts) && poppedFCosts[i] != expected {
			t.Errorf("At position %d: expected F=%.1f, got F=%.1f", i, expected, poppedFCosts[i])
		}
	}
}

func TestPriorityQueueHeapInterface(t *testing.T) {
	// Test that our PriorityQueue correctly implements heap.Interface
	pq := NewPriorityQueue()

	// Add some nodes
	node1 := NewNode(1, 1)
	node1.F = 10.0

	node2 := NewNode(2, 2)
	node2.F = 5.0

	heap.Push(pq, node1)
	heap.Push(pq, node2)

	// Verify heap property using heap operations
	if pq.Len() != 2 {
		t.Errorf("Expected length 2, got %d", pq.Len())
	}

	// The node with lower F cost should be at the top
	if pq.nodes[0].F != 5.0 {
		t.Errorf("Expected root F cost 5.0, got %.1f", pq.nodes[0].F)
	}

	// Pop using heap interface
	popped := heap.Pop(pq).(*Node)
	if popped.F != 5.0 {
		t.Errorf("Expected to pop node with F=5.0, got F=%.1f", popped.F)
	}
}

// Edge case tests
func TestPriorityQueueEdgeCases(t *testing.T) {
	pq := NewPriorityQueue()

	// Test with nodes having same F cost
	node1 := NewNode(1, 1)
	node1.F = 5.0

	node2 := NewNode(2, 2)
	node2.F = 5.0

	pq.PushNode(node1)
	pq.PushNode(node2)

	// Both should be poppable (order doesn't matter since F costs are equal)
	popped1 := pq.PopNode()
	popped2 := pq.PopNode()

	if popped1 == nil || popped2 == nil {
		t.Error("Should be able to pop both nodes with equal F costs")
	}

	if (popped1 != nil && popped1.F != 5.0) || (popped2 != nil && popped2.F != 5.0) {
		t.Error("Both popped nodes should have F cost 5.0")
	}
}

// Benchmarks
func BenchmarkPriorityQueuePush(b *testing.B) {
	pq := NewPriorityQueue()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node := NewNode(i%1000, i/1000)
		node.F = float64(i % 100)
		pq.PushNode(node)
	}
}

func BenchmarkPriorityQueuePop(b *testing.B) {
	pq := NewPriorityQueue()

	// Pre-fill the queue
	for i := 0; i < b.N; i++ {
		node := NewNode(i%1000, i/1000)
		node.F = float64(i % 100)
		pq.PushNode(node)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.PopNode()
	}
}

func BenchmarkPriorityQueueContains(b *testing.B) {
	pq := NewPriorityQueue()

	// Pre-fill the queue
	for i := 0; i < 1000; i++ {
		node := NewNode(i%100, i/100)
		node.F = float64(i)
		pq.PushNode(node)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Contains(i%100, (i/100)%10)
	}
}
