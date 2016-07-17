package eago

// Compare function
// Arg1: DNA1 to be compared
// Arg2: DNA2 to compare with
// return DNA1 < DNA2:     -1
//        DNA1 == DNA2:     0
//        DNA1 > DNA2:      1
type CompareFunc func(*DNA, *DNA) int

// Evaluate function
// Arg: DNA to be evaluated
// return fitness value in float64
type EvaluateFunc func(*DNA) float64

// Select function
// Arg1: Compare function
// Arg2: Population of DNAs
// return selected DNA
type SelectFunc func(CompFunc, []*DNA) *DNA

// Crossover function
// Arg1: Parent 1 DNA
// Arg2: Parent 2 DNA
// return two children DNA after crossover
type CrossoverFunc func(*DNA, *DNA) (*DNA, *DNA)

type Config struct {
	DNALen       int           // DNA length
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
func (c *Config) DNALen(len int) {
	c.DNALen = len
}

// set Population size
func (c *Config) PopSize(size int) {
	c.PopSize = size
}

// set number of generations
func (c *Config) NumGen(n int) {
	c.NumGen = n
}

// set mutation rate
func (c *Config) MutationRate(r float64) {
	c.MutationRate = r
}

// set comparison function
func (c *Config) Compare(fn CompareFunc) {
	c.Compare = fn
}

// set evaluation function
func (c *Config) Evaluate(fn EvaluateFunc) {
	c.Evaluate = fn
}

// set selection function
func (c *Config) Select(fn SelectFunc) {
	c.Select = fn
}

// set crossover function
func (c *Config) Crossover(fn CrossoverFunc) {
	c.Crossover = fn
}
