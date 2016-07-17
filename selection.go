package eago

import "math/rand"

// Tournament Selection
func TSelect(c CompareFunc, p []*DNA) *DNA {
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
func FPSelect(c CompareFunc, p []*DNA) *DNA {
	popSize := len(p)
	best := p[rand.Intn(popSize)]
	bestScore := best.Fitness()
	for i := 0; i < popSize; i++ {
		if c(p[i], best) == 1 {
			best.Copy(p[i])
			bestScore = p[i].Fitness()
		}
	}
	// stochastic acceptance
	for {
		i := rand.Intn(popSize)
		r := p[i].Fitness() / bestScore
		if rand.Float64() < r {
			return p[i]
		}
	}
}
