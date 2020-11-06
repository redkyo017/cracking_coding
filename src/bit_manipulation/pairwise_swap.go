package bit_manipulation

import "log"

func SwapOddEventBits(x int) int {
	log.Printf("x: %b", x)
	oddShift := x & 0xaaaaaaaa
	log.Printf("odd mask: %b - %d", 0xaaaa, 0xaaaa)
	log.Printf("odd: %d - %08b", oddShift, oddShift)
	evenShift := x & 0x55555555
	log.Printf("even mask: %b - %d", 0x5555, 0x5555)
	log.Printf("even: %d - %08b", evenShift, evenShift)
	right := oddShift >> 1
	left := evenShift << 1
	log.Printf("left: %08b - right: %08b", left, right)
	result := right | left
	log.Printf("result: %08b", result)
	return result
}
