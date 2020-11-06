package arraystrings

import (
	"log"
	"strconv"
)

func StringCompression(input string) (output string) {
	log.Println("con co", input)
	compressedLength, countRepeat := 0, 0
	for i := 0; i < len(input); i++ {
		countRepeat++
		if i+1 >= len(input) || input[i] != input[i+1] {
			compressedLength += countRepeat
			countRepeat = 0
		}
	}

	if compressedLength > len(input) {
		return input
	}
	output = compressBad(input)
	return output
}

func compressBad(in string) (out string) {
	countRepeat := 0
	for i := 0; i < len(in); i++ {
		countRepeat++
		if i+1 >= len(in) || in[i] != in[i+1] {
			out += string(in[i]) + strconv.Itoa(countRepeat)
			countRepeat = 0
		}
	}
	if len(out) >= len(in) {
		return in
	}
	return out
}
