package recursion_dp

import (
	"log"
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
	return b[i].Height < b[j].Height
}
func (b Boxes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func CreateStack(boxes Boxes) int {
	sort.Sort(boxes)
	log.Println(boxes)
	return 0
}

func ImplementStackOfBox() {
	boxes := []Box{
		Box{Height: 3},
		Box{Height: 1},
		Box{Height: 7},
		Box{Height: 9},
		Box{Height: 5},
	}
	CreateStack(Boxes(boxes))
}
