# Go Pathfinding Algorithms

A learning project for implementing pathfinding algorithms in Go, starting with A* (A-star).

> Scope & milestones per [docs/PLANNING.md](docs/PLANNING.md).

## Pre-requisites
- Go >= 1.21
- OS: macOS/Linux/Windows (amd64/arm64)

## Quick Start
```bash
git clone https://github.com/edgejay/go-pathfinding.git
cd go-pathfinding
go mod download
go test ./... -race
```

## Basic Usage

```go
package main

import (
    "fmt"
    "github.com/edgejay/go-pathfinding/algo"
)

func main() {
    // Create a 10x10 grid with 4-way movement
    grid, err := algo.NewGrid(10, 10, algo.FourWay)
    if err != nil {
        panic(err)
    }
    
    // Add some obstacles
    grid.SetObstacle(5, 5)
    grid.SetObstacle(5, 6)
    grid.SetObstacle(5, 7)
    
    // Create A* pathfinder
    astar := algo.NewAStar()
    astar.SetGrid(grid)
    astar.SetHeuristic(algo.Manhattan) // Best for 4-way movement
    
    // Find path from top-left to bottom-right
    start, _ := grid.GetNode(0, 0)
    goal, _ := grid.GetNode(9, 9)
    
    path, err := astar.FindPath(start, goal)
    if err != nil {
        fmt.Printf("No path found: %v\n", err)
        return
    }
    
    fmt.Printf("Path found with %d steps:\n", len(path))
    for i, node := range path {
        fmt.Printf("%d: (%d, %d)\n", i, node.X, node.Y)
    }
}
```

## Available Heuristics

- **Manhattan**: `algo.Manhattan` - Optimal for 4-way movement
- **Euclidean**: `algo.Euclidean` - Best for unrestricted movement
- **Diagonal**: `algo.Diagonal` - Optimal for 8-way movement
- **Zero**: `algo.Zero` - Converts A* to Dijkstra's algorithm

## Development

### Run Tests
```bash
go test ./... -race -coverprofile=coverage.out
go tool cover -func=coverage.out
```

### Run Benchmarks
```bash
go test ./algo -bench=BenchmarkAStar -benchmem
```

### Performance Results
Current benchmark results on our test hardware:
- **Small grids (100x100)**: ~0.21ms ✅ (Target: < 1ms)
- **Medium grids (500x500)**: ~2.2ms ✅ (Target: < 10ms)  
- **Large grids (1000x1000)**: ~7.5ms ✅ (Target: < 100ms)
- **Test coverage**: 98.2% ✅ (Target: ≥90%)

## Architecture

```
├── algo/                      # Core pathfinding algorithms
│   ├── astar.go              # A* implementation
│   ├── grid.go               # Grid representation and utilities
│   ├── heuristics.go         # Heuristic function implementations
│   ├── interfaces.go         # Core interfaces
│   ├── node.go               # Node structure for pathfinding
│   └── priority_queue.go     # Priority queue for A* open set
├── cmd/                      # CLI applications (future)
├── docs/                     # Documentation and analysis
│   ├── PLANNING.md           # Project planning and milestones
│   ├── TASKS.md              # Task tracking
│   └── benchmarks/           # Performance benchmark results
└── README.md                 # This file
```

## Learning Goals

This project serves as a hands-on learning platform for:
1. **Mastering A* Algorithm**: Understanding heuristics, optimality, and performance
2. **Go Programming**: Idiomatic Go, data structures, interfaces, and testing
3. **Algorithm Analysis**: Benchmarking, complexity analysis, and optimization

## Current Status: Milestone M1 Complete ✅

- [x] Core A* algorithm implementation
- [x] Comprehensive test suite (98.2% coverage)
- [x] Performance benchmarking
- [x] Grid-based pathfinding with obstacles
- [x] Multiple heuristic functions
- [x] Full documentation

## Next Steps: M2 - Enhanced Features

- [ ] Support for weighted terrain
- [ ] 8-way movement patterns
- [ ] Path smoothing and optimization
- [ ] Memory allocation improvements

## License

MIT
