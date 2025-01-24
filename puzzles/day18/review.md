# Day 18

## Analysis

The Historians and I are on a 2D grid that is NxN size, and bytes are falling down.
The puzzle input denotes the X,Y coordinate of where these bytes will land. I.E `5,4` will land on the `grid[4][5]` where grid is a 2D array.
The puzzle prompts to start at a specific number of bytes that dropped and find the shortest path from [0, 0] to [N, N].
As an example, the sample starts with 12 bytes. Therefore, the grid is initialized to be:

```
...#...
..#..#.
....#..
...#..#
..#..#.
.#..#..
#.#....
```

## Approach

Use BFS algorithm to find the shortest path from start to end point on a 2D grid.
Since the grid is unweighted, this explores all possible paths layer by layer. Ensuring
it reaches the target in the minimum number of steps.
Use multiple 2D arrays to track visited, distance, and retracing helps rebuild the path.
Initially, without any bytes falling, the quickest way to traverse to the end would be going
in this order:

```go
directions := []Coordinate{
		{X: 1, Y: 0},  // Right
		{X: 0, Y: 1},  // Down
		{X: -1, Y: 0}, // Left
		{X: 0, Y: -1}, // Up
	}
```

This is crucial to rebuilding the correct path because if not, the path would not match
the sample given.

## Solution Debrief

For part 1, I initially didn't use a 2D array to track parenting neighbors since it wasn't
part of the problem. However, that shortly was my downfall in part 2. Luckily, the BFS
setup was there and all it took was to add the current node to the adjacent neighbor.

```go
trace[adj.Y][adj.X] = &curr
```

Then using the function to rebuild the path once algorithm reached the very end.

```go
buildPath := func(parent [][]*Coordinate, endPoint Coordinate) []Coordinate {
		var path []Coordinate
		for current := &endPoint; current != nil; current = parent[current.Y][current.X] {
			path = append([]Coordinate{*current}, path...)
		}
		return path
	}
```

With the path, it was trivial to just check if the dropped byte will block the path if it lands anywhere in the path, causing the program to recalculate the path.
We repeat this until no path can be found.

Stats

| **Part** | **Average Time** | **Average Bytes Allocated** | **Average Memory Allocations** |
| -------- | ---------------- | --------------------------- | ------------------------------ |
| 1        | 563,715 ns/op    | 1,167,701 B/op              | 5,339 allocs/op                |
| 2        | 20,848,934 ns/op | 46,479,618 B/op             | 158,542 allocs/op              |

## Self-Reflect

Due to previous bad experiences with creating pathfinding algorithms, this one I tried
to keep it as simple as possible by utilizing multiple 2D graphs. Although it may not be
memory efficient, I was able to breeze through this problem without creating a graph.
Understanding how to create unit tests in Go have been immensely helpful in ensuring that
each partial issue was addressed.
