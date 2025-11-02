# Project Tasks (auto-curated by Copilot; source of truth)

> All tasks must trace back to docs/PLANNING.md. Use checkboxes. Keep sections small and current.

## Milestone: M1 – Core A* Algorithm (Foundation)
- [ ] Design data model for grid-based pathfinding (ref: docs/PLANNING.md §M1)
- [ ] Implement Node and Grid structures in package `algo`
- [ ] Create priority queue implementation for A* open set
- [ ] Implement basic A* algorithm v1 in `algo/astar.go`
- [ ] Add Manhattan, Euclidean, and Diagonal heuristic functions
- [ ] Create comprehensive unit tests for edge cases (empty grid, no path, obstacles)
- [ ] Add benchmarks for small/medium/large grids
- [ ] Document public API functions in `algo` package (GoDoc style)
- [ ] Achieve ≥90% test coverage for core algorithm
- [ ] Update README with basic usage example

## Milestone: M2 – Enhanced Features & Optimization
- [ ] Support 4-way and 8-way movement patterns
- [ ] Implement weighted terrain/movement costs
- [ ] Add configurable tie-breaking strategies for A*
- [ ] Implement path smoothing/post-processing
- [ ] Optimize memory allocation patterns
- [ ] Add extended benchmark suite for different scenarios
- [ ] Profile memory usage and document patterns
- [ ] Performance comparison against different grid complexities

## Milestone: M3 – CLI Tool & Visualization
- [ ] Create CLI application in `cmd/pathfinder/`
- [ ] Implement map loading from JSON/text formats
- [ ] Add ASCII visualization of paths and exploration
- [ ] Support configurable algorithm parameters via CLI flags
- [ ] Add export functionality for results and statistics
- [ ] Create cross-platform build scripts
- [ ] Write usage documentation and examples
- [ ] Test CLI on different platforms

## Milestone: M4 – Analysis & Documentation
- [ ] Write comprehensive algorithm analysis in `docs/analysis/`
- [ ] Create performance comparison reports
- [ ] Document learning notes and insights
- [ ] Write step-by-step tutorial in `docs/tutorial/`
- [ ] Add code examples and API documentation
- [ ] Create contribution guidelines
- [ ] Update README with complete examples and setup

## Housekeeping
- [ ] CI: run `go test ./... -race -coverprofile=coverage.out`
- [ ] Lint: `golangci-lint run`
- [ ] Benchmarks recorded in `docs/benchmarks/`
- [ ] Ensure Go mod tidy and dependency management
- [ ] Cross-platform compatibility testing