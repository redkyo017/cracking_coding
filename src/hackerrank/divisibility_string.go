package hackerrank

import (
	"log"
	"strings"
)

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
			return int32(len(checkRepeatedString(t)))
		}
	}
	return -1
}

func checkRepeatedString(s string) string {
	doubleString := s + s
	trim := doubleString[1 : len(doubleString)-1]
	existIndex := strings.Index(trim, s)
	if existIndex != -1 {
		return string(s[:existIndex+1])
	}
	return s
}

func FindSmallestDivisorImplement() {
	s := "bcdbcdbcdbcd"
	t := "bcdbcd"
	log.Println(findSmallestDivisor(s, t))
}
