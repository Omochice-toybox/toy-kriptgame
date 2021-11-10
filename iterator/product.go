package iterator

type productGenerator struct {
	source  []interface{}
	nPick   int
	current []int
}

func NewStringProductGenerator(source []string, nPick int) *productGenerator {
	interfaced := make([]interface{}, len(source))
	for i, v := range source {
		interfaced[i] = v
	}
	current := make([]int, nPick)
	current[nPick-1]--
	return &productGenerator{
		source:  interfaced,
		nPick:   nPick,
		current: current,
	}
}

func NewIntProductGenerator(source []int, nPick int) *productGenerator {
	interfaced := make([]interface{}, len(source))
	for i, v := range source {
		interfaced[i] = v
	}
	current := make([]int, nPick)
	current[nPick-1]--
	return &productGenerator{
		source:  interfaced,
		nPick:   nPick,
		current: current,
	}
}

func (g *productGenerator) GetCurrent() []interface{} {
	rst := make([]interface{}, len(g.current))
	for i, v := range g.current {
		rst[i] = g.source[v]
	}
	return rst
}

func (g *productGenerator) Next() bool {
	carrying := 1
	for i := len(g.current) - 1; i >= 0; i-- {
		next := g.current[i] + carrying
		g.current[i] = next % len(g.source)
		carrying = next / len(g.source)
	}
	return carrying == 0
}
