package eago

// (mu + lambda) Evolutionary Strategy
type MLESPlus struct {
	conf       *Config // configuration
	best       *DNA    // best performing gene
	population []*DNA  // population
}
