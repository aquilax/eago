package eago

// Compare function
// Arg1: gene1 to be compared
// Arg2: gene2 to compare with
// return gene1 < gene2:     -1
//        gene1 == gene2:     0
//        gene1 > gene2:      1
type CompareFunc func(*DNA, *DNA) int

// Evaluate function
// Arg: gene to be evaluated
// return fitness value in float64
type EvaluateFunc func(*DNA) float64

// Select function
// Arg1: Compare function
// Arg2: Population of genes
// return selected DNA
type SelectFunc func(CompareFunc, []*DNA) *DNA

// Crossover function
// Arg1: Parent 1 gene
// Arg2: Parent 2 gene
// return two children genes after crossover
type CrossoverFunc func(*DNA, *DNA) (*DNA, *DNA)

// configuration
type Config struct {
	GeneLen      int           // gene length
	PopSize      int           // population size
	NumGen       int           // number of generations
	MutationRate float64       // mutation rate
	Compare      CompareFunc   // compare function
	Evaluate     EvaluateFunc  // evaluate function
	Select       SelectFunc    // select function
	Crossover    CrossoverFunc // crossover function
}

// new configuration default to null
func NewConfig() *Config {
	return &Config{}
}

// set DNA length
func (c *Config) SetGeneLen(len int) {
	c.GeneLen = len
}

// set Population size
func (c *Config) SetPopSize(size int) {
	c.PopSize = size
}

// set number of generations
func (c *Config) SetNumGen(n int) {
	c.NumGen = n
}

// set mutation rate
func (c *Config) SetMutationRate(r float64) {
	c.MutationRate = r
}

// set comparison function
func (c *Config) SetCompare(fn CompareFunc) {
	c.Compare = fn
}

// set evaluation function
func (c *Config) SetEvaluate(fn EvaluateFunc) {
	c.Evaluate = fn
}

// set selection function
func (c *Config) SetSelect(fn SelectFunc) {
	c.Select = fn
}

// set crossover function
func (c *Config) SetCrossover(fn CrossoverFunc) {
	c.Crossover = fn
}
