package pkg

import (
	"errors"
)

type Unit string

type UnitConverter struct {
	conversionMap map[Unit]map[Unit]float64
}

var (
	errConversionImpossible = errors.New("Conversion impossible")
)

func (uc *UnitConverter) Convert(value float64, from, to Unit) (float64, error) {
	visited := make(map[Unit]struct{})
	return uc.convert(value, from, to, visited)
}

func (uc *UnitConverter) convert(value float64, from, to Unit, visited map[Unit]struct{}) (float64, error) {
	if factor, found := uc.conversionMap[from][to]; found {
		return value * factor, nil
	}
	for nextTo, factor := range uc.conversionMap[from] {
		if _, seen := visited[nextTo]; seen {
			continue
		}
		visited[nextTo] = struct{}{}
		if result, err := uc.convert(value*factor, nextTo, to, visited); err == nil {
			return result, nil
		}
		delete(visited, nextTo)
	}
	return 0.0, errConversionImpossible
}

func (uc *UnitConverter) AddConversion(factor float64, from, to Unit) {
	if _, exists := uc.conversionMap[from]; !exists {
		uc.conversionMap[from] = make(map[Unit]float64)
	}
	if _, exists := uc.conversionMap[to]; !exists {
		uc.conversionMap[to] = make(map[Unit]float64)
	}
	uc.conversionMap[from][to] = factor
	uc.conversionMap[to][from] = 1.0 / factor
}

func NewUnitConverter() *UnitConverter {
	return &UnitConverter{make(map[Unit]map[Unit]float64)}
}
