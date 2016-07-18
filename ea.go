package eago

import "fmt"

type EA interface {
	// initialize population
	initPopulation()
	// assess fitness
	assessFitness()
	// update best/population
	update()
	// execute algorithm
	Run()
}

// (mu, lambda) Evolutionary Strategy
type MLESComma struct {
	mu         int     // selected parents
	lambda     int     // generated children
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}

func NewMLESComma(mu, lambda int, conf *Config) *MLESComma {
	// lambda has to be divisable by mu
	if lambda%mu != 0 {
		err := "Lambda has to be divisable by mu: %d %% %d != 0"
		panic(fmt.Errorf(err, lambda, mu))
	}
	return &MLESComma{
		mu:         mu,                   // selected parents
		lambda:     lambda,               // generated children
		conf:       conf,                 // configuration
		best:       NewDNA(conf.GeneLen), // random DNA as best
		population: make([]*DNA, lambda), // population
	}
}

// initialize population
func (m *MLESComma) initPopulation() {
	len := m.conf.GeneLen
	for i, _ := range m.population {
		m.population[i] = NewDNA(len)
	}
}

// assess each DNA's fitness
func (m *MLESComma) assessFitness() {
	for i, _ := range m.population {
		m.population[i].Reset()
		m.population[i].Evaluate(m.conf.Evaluate)
	}
	m.population = m.quickSort(m.population)
}

// sort population by fitness (high - low)
func (m *MLESComma) quickSort(p []*DNA) []*DNA {
	if len(p) <= 1 {
		return p
	}
	pivot := p[len(p)/2]
	low := []*DNA{}
	high := []*DNA{}
	equal := []*DNA{}
	for _, dna := range p {
		switch m.conf.Compare(dna, pivot) {
		case -1:
			low = append(low, dna)
		case 1:
			high = append(high, dna)
		default:
			equal = append(equal, dna)
		}
	}
	low = append(equal, m.quickSort(low)...)
	high = m.quickSort(high)
	return append(high, low...)
}

// update
func (m *MLESComma) update() {
	// assuming population is already sorted
	selected := m.population[:m.mu]
	ratio := m.lambda / m.mu
	children := make([]*DNA, m.lambda)
	for i := 0; i < ratio; i++ {
	}
}

// run (mu, lambda) ES
func (m *MLESComma) Run() {
	m.initPopulation()
}

// (mu + lambda) Evolutionary Strategy
type MLESPlus struct {
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
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
		best:       NewDNA(conf.GeneLen),       // random DNA as best
		population: make([]*DNA, conf.PopSize), // population
	}
}

// initialize random population
func (g *GA) initPopulation() {
	len := g.conf.GeneLen
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
func (g *GA) Run() {
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
