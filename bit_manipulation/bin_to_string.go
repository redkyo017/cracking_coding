package bit_manipulation

import "log"

func PrintBinary(num float32) string {
	if num > 1 || num < 0 {
		return "ERROR"
	}
	binString := "."
	for num > 0 {
		if len(binString) >= 32 {
			return "ERROR"
		}
		n := num * 2
		log.Println("con co", num, n)
		if n >= 1 {
			binString += "1"
			num = n - 1
		} else {
			binString += "0"
			num = n
		}
	}
	return binString
}
