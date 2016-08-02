package eago

// Genetic Algorithm
type GA struct {
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}

func NewGA(conf *Config) *GA {
	return &GA{
		conf:       conf,                       // configuration
		best:       NewDNA(conf.GeneLen),       // random DNA as best
		population: make([]*DNA, conf.PopSize), // population
	}
}

// initialize random population
func (g *GA) InitPopulation() {
	len := g.conf.GeneLen
	for i, _ := range g.population {
		g.population[i] = NewDNA(len)
	}
}

// assess each DNA's fitness
func (g *GA) AssessFitness() {
	for i, _ := range g.population {
		g.population[i].Reset()
		g.population[i].Evaluate(g.conf.Evaluate)
	}
	g.population = g.quickSort(g.population)
}

// sort population by fitness (high - low)
func (g *GA) quickSort(p []*DNA) []*DNA {
	if len(p) <= 1 {
		return p
	}
	pivot := p[len(p)/2]
	low := []*DNA{}
	high := []*DNA{}
	equal := []*DNA{}
	for _, dna := range p {
		switch g.conf.Compare(dna, pivot) {
		case -1:
			low = append(low, dna)
		case 1:
			high = append(high, dna)
		default:
			equal = append(equal, dna)
		}
	}
	low = append(equal, g.quickSort(low)...)
	high = g.quickSort(high)
	return append(high, low...)
}

// update states
func (g *GA) Update() {
	// update the best
	localBest := g.population[0]
	if g.conf.Compare(localBest, g.best) == 1 {
		g.best.Copy(localBest)
	}
	// update the population
	for i := 0; i < g.conf.PopSize; i++ {
		p1 := g.conf.Select(g.conf.Compare, g.population)
		p2 := g.conf.Select(g.conf.Compare, g.population)
		c1, c2 := g.conf.Crossover(p1, p2)
		c1.Mutate(g.conf.MutationRate)
		c2.Mutate(g.conf.MutationRate)
		p1.Copy(c1)
		p2.Copy(c2)
	}
}

// run GA
func (g *GA) Run() {
	g.InitPopulation()
	for i := 0; i < g.conf.NumGen; i++ {
		g.AssessFitness()
		g.Update()
	}
}

// get best DNA
func (g *GA) Best() *DNA {
	return g.best
}
