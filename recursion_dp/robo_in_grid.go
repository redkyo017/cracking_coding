package recursion_dp

import (
	"fmt"
)

type Point struct {
	Row int
	Col int
}

func FindPaths(maze [][]bool) []Point {
	if maze == nil || len(maze) == 0 {
		return nil
	}
	points := []Point{}
	failedPoints := map[string]bool{}
	if GetPath(maze, len(maze[0])-1, len(maze)-1, &points, failedPoints) {
		// log.Println("failed", failedPoints)
		return points
	}
	// return points
	return nil
}

func GetPath(maze [][]bool, row int, col int, paths *[]Point, failedPoints map[string]bool) bool {
	if row < 0 || col < 0 || !maze[row][col] {
		return false
	}
	key := fmt.Sprintf("%d_%d", row, col)
	if _, ok := failedPoints[key]; ok {
		return false
	}
	isAtOrigin := row == 0 && col == 0
	if isAtOrigin || GetPath(maze, row, col-1, paths, failedPoints) || GetPath(maze, row-1, col, paths, failedPoints) {
		p := Point{row, col}
		*paths = append(*paths, p)
		return true
	}
	failedPoints[key] = false
	return false
}
