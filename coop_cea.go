package eago

// Cooperative Coevolution Algorithm
type CoopCEA struct {
	numSubp    int      // number of subpopulations
	conf       *Config  // configuration of each subpopulation
	best       []*DNA   // best resulting group of DNAs
	population [][]*DNA // group of sub-populations
}

func NewCoopCEA(numSubp int, conf *Config) *CoopCEA {
	return &CoopCEA{
		conf: conf,
		// initialize best group of DNAs
		best: func() []*DNA {
			best := make([]*DNA, numSubp)
			for i, _ := range best {
				best[i] = NewDNA(conf.GeneLen)
			}
			return best
		}(),
		// initialize population
		population: func() [][]*DNA {
			population := make([][]*DNA, numSubp)
			for i, _ := range population {
				population[i] = make([]*DNA, conf.PopSize)
				for j, _ := range population[i] {
					population[i][j] = NewDNA(conf.GeneLen)
				}
			}
			return population
		}(),
	}
}

// Initialize population.
func (c *CoopCEA) InitPopulation() {
	for i, _ := range c.population {
		for j, _ := range c.population[i] {
			c.population[i][j] = NewDNA(c.conf.GeneLen)
		}
	}
}

// Assess each DNA's fitness.
func (c *CoopCEA) AssessFitness() {
	selected := make([]*DNA, c.numSubp)
}
