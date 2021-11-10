package iterator

func getInitialized(length int, isOrdered bool) []int {
	initialized := make([]int, length)
	if isOrdered {
		for i := 0; i < length; i++ {
			initialized[i] = 0
		}
	} else {
		for i := 0; i < length; i++ {
			initialized[i] = i
		}
	}
	initialized[length-1]--
	return initialized
}

func isDuplicated(arr []int) bool {
	m := map[int]bool{}
	for _, v := range arr {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

func isOrdered(arr []int, ascending bool, allowSame bool) bool {
	for i := 0; i < len(arr)-1; i++ {
		if ascending {
			if allowSame {
				if arr[i] > arr[i+1] {
					return false
				}
			} else {
				if arr[i] >= arr[i+1] {
					return false
				}
			}
		} else {
			if allowSame {
				if arr[i] < arr[i+1] {
					return false
				}
			} else {
				if arr[i] <= arr[i+1] {
					return false
				}
			}
		}
	}
	return true
}

type combinationGenerator struct {
	source         []interface{}
	nPick          int
	current        []int
	allowDuplicate bool
}

/* This is not work now...
func NewCombinations[T any] (source []T){
    ...
}
*/

func NewIntCombinations(source []int, nPick int) *combinationGenerator {
	interfaced := make([]interface{}, len(source))
	for i, v := range source {
		interfaced[i] = v
	}
	return &combinationGenerator{
		source:         interfaced,
		nPick:          nPick,
		current:        getInitialized(nPick, true),
		allowDuplicate: false,
	}
}

func NewIntCombinationsWithReplacement(source []int, nPick int) *combinationGenerator {
	interfaced := make([]interface{}, len(source))
	for i, v := range source {
		interfaced[i] = v
	}
	return &combinationGenerator{
		source:         interfaced,
		nPick:          nPick,
		current:        getInitialized(nPick, true),
		allowDuplicate: true,
	}
}
func (g *combinationGenerator) GetCurrent() []interface{} {
	rst := make([]interface{}, len(g.current))
	for i, v := range g.current {
		rst[i] = g.source[v]
	}
	return rst
}
func (g *combinationGenerator) Next() bool {
	carrying := 1
	for i := len(g.current) - 1; i >= 0; i-- {
		next := g.current[i] + carrying
		g.current[i] = next % len(g.source)
		carrying = next / len(g.source)
	}
	if carrying == 1 {
		return false
	} else {
		if !isOrdered(g.current, true, g.allowDuplicate) {
			return g.Next()
		}
		return carrying == 0
	}
}
