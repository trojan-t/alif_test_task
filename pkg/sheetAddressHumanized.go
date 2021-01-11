package pkg

import (
	"strings"
)

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
