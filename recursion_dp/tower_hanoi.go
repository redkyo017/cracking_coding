package recursion_dp

import "log"

func TowerOfHanoi(n int) {
	origin, destination, buffer := []int{}, []int{}, []int{}
	for i := n; i >= 0; i-- {
		origin = append(origin, i)
	}
	log.Printf("origin: %p - des: %p - buff: %p", &origin, &destination, &buffer)
	log.Println("before :", origin, destination, buffer)
	MoveDisks(n, &origin, &destination, &buffer)
	log.Println("after :", origin, destination, buffer)
}

func MoveDisks(n int, origin, destination, buffer *[]int) {
	// log.Printf("con heo origin: %p - des: %p - buff: %p", &origin, &destination, &buffer)
	if n < 0 {
		return
	}
	MoveDisks(n-1, origin, buffer, destination)
	MoveTop(origin, destination)
	MoveDisks(n-1, buffer, destination, origin)
	log.Println("con heo :", origin, destination, buffer)
}

func MoveTop(origin, destination *[]int) {
	if len(*origin) <= 0 {
		return
	}
	top := (*origin)[len(*origin)-1]
	*destination = append(*destination, top)
	*origin = (*origin)[:len(*origin)-1]
}
