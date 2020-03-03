package recursion_dp

func MakeChange(n int) int {
	unitChanges := []int{25, 10, 5, 1}
	storeDup := make([][]int, n+1)
	for k, _ := range storeDup {
		storeDup[k] = make([]int, len(unitChanges))
	}
	return MakeChangeFast(n, unitChanges, 0, &storeDup)
	// return MakeChangeSlow(n, unitChanges, 0)
}

func MakeChangeSlow(amount int, changeUnits []int, index int) int {
	if index >= len(changeUnits)-1 {
		return 1
	}
	unitAmount := changeUnits[index]
	ways := 0
	for i := 0; i*unitAmount <= amount; i++ {
		amountRemaining := amount - i*unitAmount
		ways += MakeChangeSlow(amountRemaining, changeUnits, index+1)
	}
	return ways
}

func MakeChangeFast(amount int, changeUnits []int, index int, storeDuplicate *[][]int) int {
	if (*storeDuplicate)[amount][index] > 0 {
		return (*storeDuplicate)[amount][index]
	}
	if index >= len(changeUnits)-1 {
		return 1
	}
	unitAmount := changeUnits[index]
	ways := 0
	for i := 0; i*unitAmount <= amount; i++ {
		amountRemaining := amount - i*unitAmount
		ways += MakeChangeFast(amountRemaining, changeUnits, index+1, storeDuplicate)
	}
	(*storeDuplicate)[amount][index] = ways
	return ways
}
