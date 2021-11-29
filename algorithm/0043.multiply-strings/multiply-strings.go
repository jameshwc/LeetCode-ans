package main

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	if strings.EqualFold(num1, "0") || strings.EqualFold(num2, "0") {
		return "0"
	}
	i := len(num2) - 1
	cnt := 0
	var singleLayerMultiplyResult []string
	for i > -1 {
		var sb strings.Builder
		for j := 0; j < cnt; j++ {
			sb.WriteByte('0')
		}
		carry := 0
		for j := len(num1) - 1; j > -1; j-- {
			n := int(num2[i]-'0')*int(num1[j]-'0') + carry
			carry = n / 10
			n = n % 10
			sb.WriteByte(byte(n + '0'))
		}
		if carry > 0 {
			sb.WriteByte(byte(carry + '0'))
		}
		singleLayerMultiplyResult = append(singleLayerMultiplyResult, reverse(sb.String()))
		cnt++
		i--
	}
	s := "0"
	for n := range singleLayerMultiplyResult {
		s = add(s, singleLayerMultiplyResult[n])
	}
	return s
}

func add(num1, num2 string) string {
	var sb strings.Builder
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0
	for i > -1 || j > -1 {
		n := carry
		if i > -1 {
			n += int(num1[i] - '0')
		}
		if j > -1 {
			n += int(num2[j] - '0')
		}
		carry = n / 10
		n = n % 10
		sb.WriteByte(byte(n + '0'))
		i--
		j--
	}
	if carry == 1 {
		sb.WriteByte('1')
	}
	return reverse(sb.String())
}

func reverse(a string) string {
	chars := []rune(a)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
func main() {
	fmt.Println(multiply("99999", "0"))
}
