package pkg

import "strings"

type PrefixMapSum struct {
	SumMap        map[string]int
	SumMapKeyList []string
}

func (p *PrefixMapSum) Insert(key string, value int) {
	_, sumExists := p.SumMap[key]
	if !sumExists {
		p.SumMapKeyList = append(p.SumMapKeyList, key)
	}
	p.SumMap[key] = value
}

func (p *PrefixMapSum) Sum(prefix string) int {
	var sum int
	for _, sumMapKey := range p.SumMapKeyList {
		if strings.HasPrefix(sumMapKey, prefix) {
			sum = sum + p.SumMap[sumMapKey]
		}
	}
	return sum

}
