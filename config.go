package eago

type Config struct {
	DNALen  int64              // DNA length
	PopSize int64              // population size
	NumGen  int64              // number of generations
	Eval    func(*DNA) float64 // evaluation function
}

// new configuration default to null
func NewConfig() *Config {
	return &Config{
		DNALen:  0,   // default: 0
		PopSize: 0,   // default: 0
		NumGen:  0,   // default: 0
		Eval:    nil, // default: nil
	}
}
