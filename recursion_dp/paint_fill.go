package recursion_dp

import "log"

const (
	Black = iota
	White
	Red
	Yellow
	Green
)

func PaintFill() {
	screen := [][]int{
		[]int{Black, White, Red, Yellow, Green},
		[]int{Black, White, Red, Yellow, Green},
		[]int{Black, White, White, White, Green},
		[]int{Black, White, White, White, Green},
		[]int{Black, White, Red, Yellow, Green},
	}
	for _, line := range screen {
		log.Println("before:", line)
	}
	newColor := Red
	log.Println("old color:", screen[2][1])
	log.Println("new color:", newColor)
	ClickPixel(&screen, 2, 1, newColor)
	for _, line := range screen {
		log.Println("after:", line)
	}
}

func ClickPixel(screen *[][]int, row, col int, newColor int) {
	if (*screen)[row][col] == newColor {
		return
	}
	Paint(screen, row, col, (*screen)[row][col], newColor)
}

func Paint(screen *[][]int, row, col int, oldColor int, newColor int) {
	if row < 0 || row >= len(*screen) || col < 0 || col >= len((*screen)[0]) {
		return
	}
	if (*screen)[row][col] == oldColor {
		(*screen)[row][col] = newColor
		Paint(screen, row-1, col, oldColor, newColor) // up
		Paint(screen, row+1, col, oldColor, newColor) // down
		Paint(screen, row, col-1, oldColor, newColor) // left
		Paint(screen, row, col+1, oldColor, newColor) // right
	}
}
