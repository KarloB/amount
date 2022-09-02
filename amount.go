package amount

import (
	"bytes"
	"fmt"
	"strings"
)

func Amount(in float64) string {
	res := strings.Replace(fmt.Sprintf("%.2f", in), ".", ",", 1)
	if ix := strings.Index(res, ","); ix > 0 {
		parts := strings.Split(res, ",")
		parts[0] = insertDots(parts[0], 3) // dot every 3 digits
		res = strings.Join(parts, ",")
	}
	return res
}

func insertDots(s string, n int) string {
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
