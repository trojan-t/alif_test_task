package pkg

import (
	"strings"
)

//def encoding(n):
//code = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//base = len(code)
//letter1 = []
//
//while n > 0:
//letter = code[(n-1) % base]
//letter1.append(letter)
//n = (n-1) // base # result is different because of (n-1) only for cases of 'Z'
//
//return ''.join(reversed(letter1))
//
//# letter_ord * base^(pos from right)
//
//col = encoding(27)
//
//print(col)

func SheetAddressHumanized(index int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLen := len(alphabet)
	var reversedAddress []string

	for index > 0 {
		letter := alphabet[(index-1)%alphabetLen]
		reversedAddress = append(reversedAddress, string(letter))
		index = (index - 1) / alphabetLen
	}
	var sb strings.Builder
	for i := len(reversedAddress) - 1; i >= 0; i-- {
		sb.WriteString(reversedAddress[i])
	}
	return sb.String()

}
