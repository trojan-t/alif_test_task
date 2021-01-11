package pkg

import (
	"math"
	"strconv"
	"strings"
)

func EgyptianFraction(top, bottom float64) string {
	bottomList := []float64{}
	for top != 0 {
		x := math.Ceil(bottom / top)
		bottomList = append(bottomList, x)
		top = x*top - bottom
		bottom *= x
	}
	var sb strings.Builder
	for _, currentBottom := range bottomList {
		sb.WriteString("1 / ")
		sb.WriteString(strconv.Itoa(int(currentBottom)))
		sb.WriteString(", ")
	}
	return sb.String()
}
