package amount

import (
	"bytes"
	"fmt"
)

func Amount(a string) string {
	if len(a) >= 3 {
		firstpart := a[0 : len(a)-2]
		secondpart := a[len(a)-2:]

		if len(firstpart) > 3 {
			firstpart = insertDot(firstpart, 3)
		}
		a = fmt.Sprintf("%s,%s", firstpart, secondpart)

	} else {
		if len(a) == 1 {
			a = fmt.Sprintf("0%s", a)
		}
		a = fmt.Sprintf("0,%s", a)
	}
	return a
}

func insertDot(s string, n int) string {
	s = revertString(s)
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n1 && i != l1 {
			buffer.WriteRune('.')
		}
	}
	result := buffer.String()
	result = revertString(result)
	return result
}

func revertString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
