package eago

import "math/rand"

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
