package eago

import "fmt"

// Genetic Algorithm
type GA struct {
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}

func NewGA(conf *Config) *GA {
	return &GA{
		conf: conf,
		best: NewDNA(conf.GeneLen),
		// initialize population
		population: func() []*DNA {
			population := make([]*DNA, conf.PopSize)
			for i, _ := range population {
				population[i] = NewDNA(conf.GeneLen)
			}
			return population
		}(),
	}
}

// Initialize population.
func (g *GA) InitPopulation() {
	for i, _ := range g.population {
		g.population[i] = NewDNA(g.conf.GeneLen)
	}
}

// Assess each DNA's fitness.
func (g *GA) AssessFitness() {
	for i, _ := range g.population {
		g.population[i].Reset()
		g.population[i].Evaluate(g.conf.Evaluate)
	}
	g.population = quickSort(g.conf, g.population)
}

// Update the current generation.
func (g *GA) Update() {
	// update the best
	localBest := g.population[0]
	if g.conf.Compare(localBest, g.best) == 1 {
		g.best.Copy(localBest)
		fmt.Printf("BEST: %f\n", g.best.Fitness())
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

// Run GA.
func (g *GA) Run() {
	g.best.Evaluate(g.conf.Evaluate)
	for i := 0; i < g.conf.NumGen; i++ {
		g.AssessFitness()
		g.Update()
	}
}

// Get the best DNA.
func (g *GA) Best() *DNA {
	return g.best
}
