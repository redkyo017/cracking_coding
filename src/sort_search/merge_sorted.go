package sort_search

import "log"

func SortMerged(a, b []int, lastA, lastB int) []int {
	indexA := lastA - 1
	indexB := lastB - 1
	indexMerged := lastA + lastB - 1

	for indexB >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexMerged] = a[indexA]
			indexA--
		} else {
			a[indexMerged] = b[indexB]
			indexB--
		}
		log.Println("con heo", indexA, indexB, indexMerged)
		indexMerged--
	}
	return a
}

func ImplementSortMerged() {
	a := []int{1, 3, 5, 7, 9, 11, 0, 0, 0, 0, 0}
	b := []int{2, 6, 10, 14, 18}
	sortMerged := SortMerged(a, b, 6, 5)
	log.Println("con co be be", sortMerged)
}
