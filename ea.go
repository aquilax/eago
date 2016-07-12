package eago

type EA interface {
	// initialize population
	initPopulation()
	// assess fitness
	assessFitness()
	// update best/population
	update()
	// execute algorithm
	Go()
	// get best DNA
	Best() *DNA
}

// Genetic Algorithm
type GA struct {
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}

func NewGA(conf *Config) *GA {
	return &GA{
		conf:       conf,                       // configuration
		best:       NewDNA(conf.DNALen),        // random DNA as best
		population: make([]*DNA, conf.PopSize), // population
	}
}

// initialize random population
func (g *GA) initPopulation() {
	len := g.conf.DNALen
	for i, _ := range g.population {
		g.population[i] = NewDNA(len)
	}
}

// assess each DNA's fitness
func (g *GA) assessFitness() {
	for i, _ := range g.population {
		g.population[i].Reset()
		g.population[i].Evaluate(g.conf.Evaluate)
	}
	g.population = g.quickSort(g.population)
}

// sort population by fitness
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
func (g *GA) update() {
	// update the best
	best := g.population[0]
	g.best.Copy(best)
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
func (g *GA) Go() {
	g.initPopulation()
	for i := 0; i < g.conf.NumGen; i++ {
		g.assessFitness()
		g.update()
	}
}

// get best DNA
func (g *GA) Best() *DNA {
	return g.best
}
