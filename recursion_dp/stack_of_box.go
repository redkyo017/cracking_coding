package recursion_dp

import (
	"log"
	"math"
	"sort"
)

type Box struct {
	Height int
}

type Boxes []Box

func (b Boxes) Len() int {
	return len(b)
}
func (b Boxes) Less(i, j int) bool {
	return b[i].Height > b[j].Height
}
func (b Boxes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func CreateStack(boxes Boxes) int {
	sort.Sort(boxes)
	log.Println("con meo", boxes)
	maxHeight := 0
	stackMap := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		// slow solution
		// height := CreateStackSlow(boxes, i)
		// fast solution
		height := CreateStackFaster(boxes, i, stackMap)
		maxHeight = int(math.Max(float64(height), float64(maxHeight)))
	}
	return maxHeight
}

func ImplementStackOfBox() {
	boxes := []Box{
		Box{Height: 3},
		Box{Height: 1},
		Box{Height: 7},
		Box{Height: 9},
		Box{Height: 3},
		Box{Height: 4},
		Box{Height: 5},
	}
	log.Println("con co be be", CreateStack(Boxes(boxes)))
}

func CreateStackSlow(boxes []Box, bottomIndex int) int {
	bottom := boxes[bottomIndex]
	maxHeight := 0
	for i := bottomIndex + 1; i < len(boxes); i++ {
		if boxes[i].Height < bottom.Height {
			height := CreateStackSlow(boxes, i)
			maxHeight = int(math.Max(float64(height), float64(maxHeight)))
		}
	}
	maxHeight += bottom.Height
	return maxHeight
}

func CreateStackFaster(boxes []Box, bottomIndex int, stackMap []int) int {
	if bottomIndex < len(boxes) && stackMap[bottomIndex] > 0 {
		return stackMap[bottomIndex]
	}
	bottom := boxes[bottomIndex]
	maxHeight := 0
	for i := bottomIndex + 1; i < len(boxes); i++ {
		if boxes[i].Height < bottom.Height {
			height := CreateStackFaster(boxes, i, stackMap)
			maxHeight = int(math.Max(float64(height), float64(maxHeight)))
		}
	}
	maxHeight += bottom.Height
	stackMap[bottomIndex] = maxHeight
	return maxHeight
}
