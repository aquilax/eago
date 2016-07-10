package eago

type Config struct {
	DNALen       int     // DNA length
	PopSize      int     // population size
	NumGen       int     // number of generations
	MutationRate float64 // mutation rate

	Compare   func(*DNA, *DNA) int          // compare two DNAs
	Evaluate  func(*DNA) float64            // evaluation function
	Select    func([]*DNA) *DNA             // select with replacement
	Crossover func(*DNA, *DNA) (*DNA, *DNA) // crossover
}

// new configuration default to null
func NewConfig() *Config {
	return &Config{}
}
