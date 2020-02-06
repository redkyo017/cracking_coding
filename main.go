package main

import "log"

func main() {
	// node1 := linkedlist.LinkedListNode{Val: 1, Next: nil}
	// node2 := linkedlist.LinkedListNode{Val: 2, Next: &node1}
	// node3 := linkedlist.LinkedListNode{Val: 3, Next: &node2}
	// node4 := linkedlist.LinkedListNode{Val: 4, Next: &node3}
	// node5 := linkedlist.LinkedListNode{Val: 5, Next: &node4}
	// node6 := linkedlist.LinkedListNode{Val: 6, Next: &node5}
	// node7 := linkedlist.LinkedListNode{Val: 7, Next: &node6}
	// node8 := linkedlist.LinkedListNode{Val: 8, Next: &node7}
	// node9 := linkedlist.LinkedListNode{Val: 9, Next: &node8}
	// node10 := linkedlist.LinkedListNode{Val: 10, Next: &node9}
	// node11 := linkedlist.LinkedListNode{Val: 11, Next: &node10}
	// node12 := linkedlist.LinkedListNode{Val: 12, Next: &node11}

	// node1.Next = &node6
	// log.Println(linkedlist.FindBeginLoop(&node12))
	// stack_queue.ImplementAnimalShelter()
	log.Println(repeatedSubstringPattern("abac"))
}

// func strStr(haystack string, needle string) int {
// 	n := len(haystack)
// 	m := len(needle)
// 	if m == 0 {
// 		return 0
// 	}
// 	i := 0
// 	for i < n-m+1 {
// 		match := true
// 		j := 0
// 		for j < m {
// 			if haystack[i+j] != needle[j] {
// 				match = false
// 				break
// 			}
// 			j++
// 		}
// 		if match {
// 			log.Println("matching at index:", i)
// 			return i
// 		}
// 		i++
// 	}
// 	return -1
// }

// return first index of word found in text
// func KMPSearch(haystack string, needle string) int {
// 	n := len(haystack)
// 	m := len(needle)
// 	if m == 0 {
// 		return 0
// 	}

// 	lps := makeLPS(needle)
// 	i, j := 0, 0

// 	for i < n {
// 		if haystack[i] == needle[j] {
// 			i++
// 			j++
// 		}

// 		if j == m {
// 			log.Println("found pattern at ", i-j)
// 			return i - j
// 		} else if i < n && haystack[i] != needle[j] {
// 			if j != 0 {
// 				j = lps[j-1]
// 			} else {
// 				i++
// 			}
// 		}
// 	}
// 	return -1
// }

func makeLPS(s string) []int {
	lps := make([]int, len(s))
	lps[0] = 0
	m := 0 // length of longest previous prefix suffix
	i := 1
	for i < len(s) {
		if s[i] == s[m] {
			m++
			lps[i] = m
			i++
		} else {
			if m != 0 {
				m = lps[m-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	lps := makeLPS(s)
	if lps[len(lps)-1] == 0 {
		return false
	}
	lenSubStr := len(s) - lps[len(lps)-1]
	subString := s[:lenSubStr]
	if len(s)%lenSubStr != 0 {
		return false
	}
	j := 0
	for ; j < len(s); j += lenSubStr {
		if s[j:j+lenSubStr] != subString {
			break
		}
	}
	if j == len(s) {
		return true
	}
	return false
}
