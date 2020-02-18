package bit_manipulation

// Write a function to determine the number of bits you would need to flip to convert integer A to integer B.
// SOLUTION
// EXAMPLE
// Input: 29 (or: 111(1), 15 (or: (1111) Output: 2
// This seemingly complex problem is actually rather straightforward.To approach this, ask yourself how you would figure out which bits in two numbers are different. Simple: with an XOR

func BitSwapRequired(a int, b int) int {
	count := 0
	for c := a ^ b; c != 0; c = c & (c - 1) {
		count++
	}
	return count
}
