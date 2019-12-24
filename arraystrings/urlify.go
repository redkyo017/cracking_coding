package arraystrings

const space uint8 = 32
const mod uint8 = 37
const two uint8 = 50
const zero uint8 = 48

func URLify(url string) (result string) {
	var trueLength, spaceCount, index = 0, 0, 0
	byte_array := []byte(url)

	for i := len(byte_array) - 1; i >= 0; i-- {
		v := byte_array[i]
		if v != space {
			trueLength = i
			break
		}
	}
	if trueLength != 0 {
		trueLength += 1
	}

	for i := 0; i < trueLength; i++ {
		if byte_array[i] == space {
			spaceCount++
		}
	}
	index = trueLength + spaceCount*2

	newStr := make([]byte, index)
	for i := trueLength - 1; i >= 0; i-- {
		if byte_array[i] == space {
			newStr[index-1] = zero
			newStr[index-2] = two
			newStr[index-3] = mod
			index -= 3
		} else {
			newStr[index-1] = byte_array[i]
			index--
		}
	}
	result = string(newStr)
	return result
}
