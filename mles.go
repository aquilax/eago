package eago

import "errors"

var (
	// if lambda is not divisible by mu
	ErrMuLambda = errors.New("lambda has to be divisable by mu")
)

// (mu, lambda) Evolution Strategy
type MLESComma struct {
	mu, lambda int     // Mu and Lambda
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}

func NewMLESComma(mu, lambda int, conf *Config) (*MLESComma, error) {
	// lambda has to be divisable by mu
	if lambda%mu != 0 {
		return nil, ErrMuLambda
	}
	return &MLESComma{
		mu:     mu,
		lambda: lambda,
		conf:   conf,
		best:   NewDNA(conf.GeneLen),
		population: func() []*DNA {
			population := make([]*DNA, conf.PopSize)
			for i, _ := range population {
				population[i] = NewDNA(conf.GeneLen)
			}
			return population
		}(),
	}, nil
}

// Initialize the population with new random DNAs.
func (m *MLESComma) InitPopulation() {
	for i, _ := range m.population {
		m.population[i] = NewDNA(m.conf.GeneLen)
	}
}

// Assess each DNA's fitness.
func (m *MLESComma) AssessFitness() {
	for i, _ := range m.population {
		m.population[i].Reset()
		m.population[i].Evaluate(m.conf.Evaluate)
	}
	m.population = quickSort(m.conf, m.population)
}

// Update the current generation.
func (m *MLESComma) Update() {
	// update the best
	localBest := m.population[0]
	if m.conf.Compare(localBest, m.best) == 1 {
		m.best.Copy(localBest)
	}
	// assuming population is already sorted
	selected := m.population[:m.mu]
	ratio := m.lambda / m.mu
	children := make([]*DNA, 0, m.lambda)
	// O(ratio * mu) = O(lambda) -> Linear time
	for _, parent := range selected {
		for i := 0; i < ratio; i++ {
			child := NewDNA(m.conf.GeneLen)
			child.Copy(parent)
			child.Mutate(m.conf.MutationRate)
			children = append(children, child)
		}
	}
	// update the population
	m.population = children
}

// Run (mu, lambda) Evolution Strategy.
func (m *MLESComma) Run() {
	m.best.Evaluate(m.conf.Evaluate)
	for i := 0; i < m.conf.NumGen; i++ {
		m.AssessFitness()
		m.Update()
	}
}

// Get the best DNA.
func (m *MLESComma) Best() *DNA {
	return m.best
}

// (mu + lambda) Evolution Strategy
type MLESPlus struct {
	mu, lambda int     // Mu and Lambda
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}

func NewMLESPlus(mu, lambda int, conf *Config) *MLESPlus {
	return &MLESPlus{
		mu:     mu,
		lambda: lambda,
		conf:   conf,
		best:   NewDNA(conf.GeneLen),
		population: func() []*DNA {
			population := make([]*DNA, conf.PopSize)
			for i, _ := range population {
				population[i] = NewDNA(conf.GeneLen)
			}
			return population
		}(),
	}
}
