package eago

import "math/rand"

// Selection function is a type of function that selects a
// parent DNA from a population given a comparison function.
type SelectFunc func(CompareFunc, []*DNA) *DNA

// TSelect() returns a selection function that performs
// Tournament Selection.
func TSelect() SelectFunc {
	return func(c CompareFunc, p []*DNA) *DNA {
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
}

// FPSelect() returns a selection function that performs
// Fitness-Proportionate Selection (not recommended).
func FPSelect() SelectFunc {
	return func(c CompareFunc, p []*DNA) *DNA {
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
}

// MLSelect() returns a selection function for MLES algorithms. It
// simply returns Lambda best DNAs from Mu DNAs.
