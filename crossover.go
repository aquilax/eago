package eago

import (
	"bytes"
	"math"
	"math/rand"
)

// One point crossover
func Crossover1P(p1, p2 *DNA) {
	len1 := float64(p1.size)
	len2 := float64(p2.size)
	shorter := int(math.Min(len1, len2))
	xoverPoint := rand.Intn(shorter)

	var ch1, ch2 bytes.Buffer
	ch1.Write(p1.gene[:xoverPoint])
	ch1.Write(p2.gene[xoverPoint:])
	ch2.Write(p2.gene[:xoverPoint])
	ch2.Write(p1.gene[xoverPoint:])

	p1.gene = ch1.String()
	p2.gene = ch2.String()
}

// Two point crossover
func Crossover2P(p1, p2 *DNA) {
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
	ch1.Write(p1.gene[:xoverPoint1])
	ch1.Write(p2.gene[xoverPoint1:xoverPoint2])
	ch1.Write(p1.gene[xoverPoint2:])

	ch2.Write(p2.gene[:xoverPoint1])
	ch2.Write(p1.gene[xoverPoint1:xoverPoint2])
	ch2.Write(p2.gene[xoverPoint2:])

	p1.gene = ch1.String()
	p2.gene = ch2.String()
}

// Uniform crossover
func UCrossover(p1, p2 *DNA) {

}
