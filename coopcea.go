package eago

// Cooperative Coevolution Algorithm
type CoopCEA struct {
	numSubp    int     // number of subpopulations
	conf       *Config // configuration of each subpopulation
	best       []*DNA  // best resulting group of DNAs
	population []*DNA  // group of sub-populations
}

func NewCoopCEA(numSubp int, conf *Config) *NewCoopCEA {
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
		population: func() []*DNA {
			totalSize := numSubp * conf.PopSize
			population := make([]*DNA, totalSize)
			for i, _ := range populations {
				population[i] = NewDNA(conf.GeneLen)
			}
			return population
		}(),
	}
}

// Initialize population.
func (c *CoopCEA) InitPopulation() {
	for i, _ := range c.population {
		c.population[i] = NewDNA(c.conf.GeneLen)
	}
}

// Assess each DNA's fitness.
func (c *CoopCEA) AssessFitness() {
	for i := 0; i < c.numSubp; i++ {

	}
}
