package eago

type Config struct {
	DNALen       int     // DNA length
	PopSize      int     // population size
	NumGen       int     // number of generations
	MutationRate float64 // mutation rate

	// ____Compare function____
	// Arg1: DNA1 to be compared
	// Arg2: DNA2 to compare with
	// return DNA1 < DNA2:     -1
	//        DNA1 == DNA2:     0
	//        DNA1 > DNA2:      1
	Compare func(*DNA, *DNA) int

	// ____Evaluate function____
	// Arg: DNA to be evaluated
	// return fitness value in float64
	Evaluate func(*DNA) float64

	// ____Select function____
	// Arg1: Compare function
	// Arg2: Population of DNAs
	// return selected DNA
	Select func(func(*DNA, *DNA) int, []*DNA) *DNA

	// ____Crossover function____
	// Arg1: Parent 1 DNA
	// Arg2: Parent 2 DNA
	// return two children DNA after crossover
	Crossover func(*DNA, *DNA) (*DNA, *DNA)
}

// new configuration default to null
func NewConfig() *Config {
	return &Config{}
}
