package hackerrank

func findSmallestDivisor(s string, t string) int32 {
	sLen := len(s)
	tLen := len(t)
	tLenOrigin := tLen
	if sLen%tLen != 0 {
		return -1
	}
	for sLen > tLen {
		tLen += tLenOrigin
		if sLen == tLen {

		}
	}
	return -1
}

func checkRepeatedString(s string) string {
	return s
}
