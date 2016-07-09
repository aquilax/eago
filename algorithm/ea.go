package algorithm

type EA interface {
	InitPopulation()
	AssessFitness()
	Run() error
	Update()
}

// Genetic Algorithm
type GA struct {
	conf       *Config            // configuration
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
func (g *GA) InitPopulation() {
	len := g.conf.DNALen
	for i, _ := range g.population {
		g.population[i] = NewDNA(len)
	}
}

// assess each DNA's fitness
func (g *GA) AssessFitness() {
	for i, _ := range g.population {
		g.population[i].Evaluate(g.conf.Eval)
	}
}

// run GA
func (g *GA) Run() error {
	g.InitPopulation()
	g.AssessFitness()
	for i := 0; i < g.conf.NumGen; i++ {

	}
}
