package eago

// configuration
type Config struct {
	DrawGraph    bool          // whether to draw graph
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

// set whether to draw graph and output image file
func (c *Config) SetDrawGraph(b bool) {
	c.DrawGraph = b
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
