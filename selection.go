package eago

import "math/rand"

// Tournament Selection
func TSelect(c func(*DNA, *DNA) int, p []*DNA) *DNA {
	popSize := len(p)
	best := rand.Intn(popSize)
	for i := 0; i < popSize; i++ {
		next := rand.Intn(popSize)
		// if next DNA is superior
		if c(p[next], p[best]) == 1 {
			best = next
		}
	}
	return p[best]
}

// fitness-proportionate selection (not recommended)
func FPSelect(c func(*DNA, *DNA) int, p []*DNA) *DNA {
	popSize := len(p)
	var bestScore float64
	for i := 0; i < popSize; i++ {
		if c(p[i], p[best]) == 1 {
			bestScore = p[i].Fitness()
		}
	}
	// stochastic acceptance
	for {
		i := rand.Intn(popSize)
		r := p[i].Fitness() / best
		if rand.Float64() < r {
			return p[i]
		}
	}
}
