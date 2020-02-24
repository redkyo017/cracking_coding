package recursion_dp

import "fmt"

type Point struct {
	Row int
	Col int
}

func FindPaths(maze [][]bool) []Point {
	points := []Point{}
	failedPoints := map[string]bool{}
	if maze == nil || len(maze) == 0 {
		return nil
	}
	return points
}

func GetPath(maze [][]bool, row int, col int, failedPoints map[string]bool) bool {
	if row < 0 || col < 0 || !maze[row][col] {
		return false
	}
	key := fmt.Sprintf("%d_%d", row, col)
	if _, ok := failedPoints[key]; ok {
		return false
	}
	return false
}
