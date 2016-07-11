package eago

import "math/rand"

func TSelect(p []*DNA) *DNA {
	popSize := len(p)
	best := rand.Intn(popSize)
	for i := 0; i < popSize; i++ {

	}
	return p[best]
}
