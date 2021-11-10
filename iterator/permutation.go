package iterator

func hasSameValue(arr []int) bool {
	s := map[int]struct{}{}
	for _, v := range arr {
		if _, ok := s[v]; ok {
			return true
		}
		s[v] = struct{}{}
	}
	return false
}

type permutationGenerator struct {
	source  []interface{}
	nPick   int
	current []int
}

func NewIntPermutationGenerator(source []int, nPick int) *permutationGenerator {
	interfeced := make([]interface{}, len(source))
	for i, v := range source {
		interfeced[i] = v
	}
	return &permutationGenerator{
		source:  interfeced,
		nPick:   nPick,
		current: getInitialized(nPick, true),
	}
}

func NewStringPermutationGenerator(source []string, nPick int) *permutationGenerator {
	interfeced := make([]interface{}, len(source))
	for i, v := range source {
		interfeced[i] = v
	}
	return &permutationGenerator{
		source:  interfeced,
		nPick:   nPick,
		current: getInitialized(nPick, true),
	}
}

func (g *permutationGenerator) Next() bool {
	carrying := 1
	for i := len(g.current) - 1; i >= 0; i-- {
		next := g.current[i] + carrying
		g.current[i] = next % len(g.source)
		carrying = next / len(g.source)
	}
	if carrying == 1 {
		return false
	} else {
		if hasSameValue(g.current) {
			return g.Next()
		} else {
			return true
		}
	}
}
func (g *permutationGenerator) GetCurrent() []interface{} {
	rst := make([]interface{}, len(g.current))
	for i, v := range g.current {
		rst[i] = g.source[v]
	}
	return rst
}
