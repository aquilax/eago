package eago

import "fmt"

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
func (m *MLESComma) InitPopulation() {
	len := m.conf.GeneLen
	for i, _ := range m.population {
		m.population[i] = NewDNA(len)
	}
}

// assess each DNA's fitness
func (m *MLESComma) AssessFitness() {
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
func (m *MLESComma) Update() {
	// assuming population is already sorted
	//selected := m.population[:m.mu]
	//ratio := m.lambda / m.mu
	//children := make([]*DNA, m.lambda)
	//for i := 0; i < ratio; i++ {
	//}
}

// run (mu, lambda) ES
func (m *MLESComma) Run() {
	m.InitPopulation()
}
