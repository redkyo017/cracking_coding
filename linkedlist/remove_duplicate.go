package linkedlist

import "log"

func RemoveDup(node LinkedListNode) LinkedListNode {
	log.Println("node ", node)
	// cach 1: dung hash table (buffer allowed)
	hashmap := make(map[int]bool)
	run := &node
	var prev *LinkedListNode

	for run != nil {
		if _, ok := hashmap[run.Val]; ok {
			prev.Next = prev.Next.Next
		} else {
			hashmap[run.Val] = true
			prev = run
		}
		run = run.Next
	}
	log.Println("hash table", hashmap)

	// // cach 2: dung 2 con tro (no buffer allowed)
	// current := &node
	// for current != nil {
	// 	runner := current
	// 	for runner.Next != nil {
	// 		if runner.Next.Val == current.Val {
	// 			runner.Next = runner.Next.Next
	// 		} else {
	// 			runner = runner.Next
	// 		}
	// 	}
	// 	current = current.Next
	// }

	test := &node
	for test != nil {
		log.Println("con meo", test)
		test = test.Next
	}

	return node
}
