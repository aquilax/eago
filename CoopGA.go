package eago

// Cooperative Coevolution Genetic Algorithm
type CoopGA struct {
	numSubp    int     // number of subpopulations
	conf       *Config // configuration of each subpopulation
	best       []*DNA  // best resulting group of DNAs
	population []*DNA  // group of sub-populations
}

func NewCoopGA(numSubp int, conf *Config) *NewCoopGA {
	return &CoopGA{
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
