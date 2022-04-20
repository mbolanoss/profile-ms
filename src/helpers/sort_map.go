package helpers

import (
	"sort"
)

type PairInt_Int struct {
	Key int
	Value int
}

type PairString_Int struct {
	Key string
	Value int
}

type PairListInt_Int []PairInt_Int
type PairListString_Int []PairString_Int

func (p PairListInt_Int) Len() int           { return len(p) }
func (p PairListInt_Int) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairListInt_Int) Less(i, j int) bool { return p[i].Value < p[j].Value }

func (p PairListString_Int) Len() int           { return len(p) }
func (p PairListString_Int) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairListString_Int) Less(i, j int) bool { return p[i].Value < p[j].Value }

func SortMapInt_Int(m map[int]int) PairListInt_Int {
	p := make(PairListInt_Int, len(m))

	i := 0
	for k, v := range m {
		p[i] = PairInt_Int{k, v}
		i++
	}

	sort.Sort(p)
	
	return p
}

func SortMapString_Int(m map[string]int) PairListString_Int {
	p := make(PairListString_Int, len(m))

	i := 0
	for k, v := range m {
		p[i] = PairString_Int{k, v}
		i++
	}

	sort.Sort(p)
	
	return p
}