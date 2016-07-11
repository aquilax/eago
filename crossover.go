package eago

import (
	"bytes"
	"math"
	"math/rand"
)

// One point crossover
func Crossover1P(p1, p2 *DNA) (*DNA, *DNA) {
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

// Two point crossover
func Crossover2P(p1, p2 *DNA) (*DNA, *DNA) {
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

// Uniform crossover given mutation rate
func UCrossover(r float64) func(*DNA, *DNA) (*DNA, *DNA) {
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
