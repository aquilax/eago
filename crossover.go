package eago

import (
	"bytes"
	"math"
	"math/rand"
)

// Crossover function is a type of function that takes two
// parent DNAs and perform crossover and return two children
// DNAs
type CrossoverFunc func(*DNA, *DNA) (*DNA, *DNA)

// One point crossover selects a random position in the two
// parent DNAs and swap parts after that selected position;
// then return two new DNAs created with swapped genes
func Crossover1P() CrossoverFunc {
	return func(p1, p2 *DNA) (*DNA, *DNA) {
		len1 := float64(p1.size)
		len2 := float64(p2.size)
		shorter := int(math.Min(len1, len2))
		xoverPoint := rand.Intn(shorter)

		var ch1, ch2 bytes.Buffer
		ch1.WriteString(p1.gene[:xoverPoint])
		ch1.WriteString(p2.gene[xoverPoint:])
		ch2.WriteString(p2.gene[:xoverPoint])
		ch2.WriteString(p1.gene[xoverPoint:])

		return &DNA{
				size:    p1.size,
				gene:    ch1.String(),
				fitness: 0.0,
			}, &DNA{
				size:    p2.size,
				gene:    ch2.String(),
				fitness: 0.0,
			}
	}
}

// Two point crossover selects two random positions in the
// two parent DNAs and swap parts in between the selected
// positions; then return two new DNAs created with swapped
// genes.
func Crossover2P() CrossoverFunc {
	return func(p1, p2 *DNA) (*DNA, *DNA) {
		len1 := float64(p1.size)
		len2 := float64(p2.size)
		shorter := int(math.Min(len1, len2))
		xoverPoint1 := rand.Intn(shorter)
		xoverPoint2 := rand.Intn(shorter)

		// if in wrong order, swap them
		if xoverPoint1 > xoverPoint2 {
			xoverPoint1, xoverPoint2 =
				xoverPoint2, xoverPoint1
		}

		var ch1, ch2 bytes.Buffer
		ch1.WriteString(p1.gene[:xoverPoint1])
		ch1.WriteString(p2.gene[xoverPoint1:xoverPoint2])
		ch1.WriteString(p1.gene[xoverPoint2:])

		ch2.WriteString(p2.gene[:xoverPoint1])
		ch2.WriteString(p1.gene[xoverPoint1:xoverPoint2])
		ch2.WriteString(p2.gene[xoverPoint2:])

		return &DNA{
				size:    p1.size,
				gene:    ch1.String(),
				fitness: 0.0,
			}, &DNA{
				size:    p2.size,
				gene:    ch2.String(),
				fitness: 0.0,
			}
	}
}

// Uniform crossover returns a crossover function given
// a swapping rate; as the function  iterates through the
// two parent  DNAs, if a randomly generated number is smaller
// than the swapping rate, it swaps the bits in that position;
// then returns two new DNAs created with swapped genes.
func UCrossover(r float64) CrossoverFunc {
	return func(p1, p2 *DNA) (*DNA, *DNA) {
		len1 := float64(p1.size)
		len2 := float64(p2.size)
		shorter := int(math.Min(len1, len2))

		var ch1, ch2 bytes.Buffer
		for i := 0; i < shorter; i++ {
			if rand.Float64() < r {
				ch1.WriteByte(p2.gene[i])
				ch2.WriteByte(p1.gene[i])
			} else {
				ch1.WriteByte(p1.gene[i])
				ch2.WriteByte(p2.gene[i])
			}
		}
		return &DNA{
				size:    p1.size,
				gene:    ch1.String(),
				fitness: 0.0,
			}, &DNA{
				size:    p2.size,
				gene:    ch2.String(),
				fitness: 0.0,
			}
	}
}
