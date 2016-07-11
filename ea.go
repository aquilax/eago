package eago

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

// Genetic Algorithm
type GA struct {
	conf       *Config            // configuration
	best       *DNA               // best performing gene
	population []*DNA             // population
	evalfunc   func(*DNA) float64 // evaluation function
}

func NewGA(conf *Config) *GA {
	return &GA{
		conf:       conf,
		population: make([]*DNA, conf.PopSize),
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
	low = append(g.quickSort(low), equal...)
	high = g.quickSort(high)
	return append(low, high...)
}

// update states
func (g *GA) update() {
	// update the best
	size := g.conf.PopSize
	g.best.Copy(g.population[size-1])
	// update the population
	for i := 0; i < g.conf.PopSize; i++ {
		p1 := g.conf.Select(g.population)
		p2 := g.conf.Select(g.population)
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
