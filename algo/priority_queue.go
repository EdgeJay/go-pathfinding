package algo

import (
	"container/heap"
)

// PriorityQueue implements a priority queue for A* pathfinding using Go's heap interface.
// It maintains nodes ordered by their F cost (total estimated cost) in ascending order.
// This ensures that A* always explores the most promising node next.
type PriorityQueue struct {
	nodes []*Node
	// nodeIndex maps node coordinates to their index in the heap for fast lookups
	nodeIndex map[coordinate]int
}

// coordinate represents a position in the grid for efficient map lookups
type coordinate struct {
	x, y int
}

// NewPriorityQueue creates a new empty priority queue for A* pathfinding.
//
// Returns:
//
//	*PriorityQueue: new priority queue instance
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		nodes:     make([]*Node, 0),
		nodeIndex: make(map[coordinate]int),
	}
	heap.Init(pq)
	return pq
}

// Len returns the number of nodes in the priority queue.
// This is required by the heap.Interface.
func (pq *PriorityQueue) Len() int {
	return len(pq.nodes)
}

// Less compares two nodes by their F cost (total estimated cost).
// Returns true if node i has lower F cost than node j.
// This is required by the heap.Interface and ensures min-heap ordering.
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.nodes[i].F < pq.nodes[j].F
}

// Swap exchanges two nodes in the priority queue and updates their indices.
// This is required by the heap.Interface.
func (pq *PriorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]

	// Update indices in the lookup map
	pq.nodeIndex[coordinate{pq.nodes[i].X, pq.nodes[i].Y}] = i
	pq.nodeIndex[coordinate{pq.nodes[j].X, pq.nodes[j].Y}] = j
}

// Push adds a new node to the priority queue.
// This is required by the heap.Interface. Use PushNode instead for type safety.
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	pq.nodeIndex[coordinate{node.X, node.Y}] = len(pq.nodes)
	pq.nodes = append(pq.nodes, node)
}

// Pop removes and returns the node with the lowest F cost.
// This is required by the heap.Interface. Use PopNode instead for type safety.
func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.nodes)
	if n == 0 {
		return nil
	}

	node := pq.nodes[n-1]
	pq.nodes = pq.nodes[:n-1]
	delete(pq.nodeIndex, coordinate{node.X, node.Y})
	return node
}

// PushNode adds a node to the priority queue.
// If a node at the same position already exists with higher F cost,
// it updates the existing node instead of adding a duplicate.
//
// Params:
//
//	node: node to add to the queue
func (pq *PriorityQueue) PushNode(node *Node) {
	coord := coordinate{node.X, node.Y}

	// Check if node already exists
	if index, exists := pq.nodeIndex[coord]; exists {
		// Update existing node if new F cost is lower
		existingNode := pq.nodes[index]
		if node.F < existingNode.F {
			existingNode.G = node.G
			existingNode.H = node.H
			existingNode.F = node.F
			existingNode.Parent = node.Parent
			heap.Fix(pq, index)
		}
	} else {
		// Add new node
		heap.Push(pq, node)
	}
}

// PopNode removes and returns the node with the lowest F cost.
// Returns nil if the queue is empty.
//
// Returns:
//
//	*Node: node with lowest F cost, or nil if queue is empty
func (pq *PriorityQueue) PopNode() *Node {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(pq).(*Node)
}

// Contains checks if a node at the given coordinates exists in the queue.
//
// Params:
//
//	x, y: coordinates to check
//
// Returns:
//
//	bool: true if a node at these coordinates is in the queue
func (pq *PriorityQueue) Contains(x, y int) bool {
	_, exists := pq.nodeIndex[coordinate{x, y}]
	return exists
}

// ContainsNode checks if the specified node exists in the queue.
//
// Params:
//
//	node: node to check for
//
// Returns:
//
//	bool: true if the node is in the queue
func (pq *PriorityQueue) ContainsNode(node *Node) bool {
	return pq.Contains(node.X, node.Y)
}

// GetNode retrieves a node at the given coordinates from the queue.
// Returns nil if no node exists at those coordinates.
//
// Params:
//
//	x, y: coordinates to look up
//
// Returns:
//
//	*Node: node at the coordinates, or nil if not found
func (pq *PriorityQueue) GetNode(x, y int) *Node {
	if index, exists := pq.nodeIndex[coordinate{x, y}]; exists {
		return pq.nodes[index]
	}
	return nil
}

// Peek returns the node with the lowest F cost without removing it.
// Returns nil if the queue is empty.
//
// Returns:
//
//	*Node: node with lowest F cost, or nil if queue is empty
func (pq *PriorityQueue) Peek() *Node {
	if pq.Len() == 0 {
		return nil
	}
	return pq.nodes[0]
}

// Clear removes all nodes from the priority queue.
func (pq *PriorityQueue) Clear() {
	pq.nodes = pq.nodes[:0]
	pq.nodeIndex = make(map[coordinate]int)
}

// IsEmpty returns true if the priority queue contains no nodes.
//
// Returns:
//
//	bool: true if queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// UpdatePriority updates the F cost of a node already in the queue
// and re-heapifies to maintain ordering.
//
// Params:
//
//	x, y: coordinates of the node to update
//	newF: new F cost for the node
//
// Returns:
//
//	bool: true if node was found and updated
func (pq *PriorityQueue) UpdatePriority(x, y int, newF float64) bool {
	if index, exists := pq.nodeIndex[coordinate{x, y}]; exists {
		pq.nodes[index].F = newF
		heap.Fix(pq, index)
		return true
	}
	return false
}
