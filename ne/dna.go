package ne

import "math/rand"

// DNA in NE package is a real-coded genotype with a slice of
// float64 as its gene. This version of DNA is intended to be used
// specifically for NeuroEvolution.
type DNA struct {
	size    int       // size of the chromosome
	gene    []float64 // real-coded gene
	fitness float64   // fitness score
}

// Create a new real-coded DNA given its size.
func NewDNA(size int) *DNA {
	return &DNA{
		size: size,
		gene: func() []float64 {
			g := make([]float64, size)
			for i := 0; i < size; i++ {
				g[i] = rand.NormFloat64()
			}
			return g
		}(),
		fitness: 0.0,
	}
}

// Get DNA size in int.
func (d *DNA) Size() int {
	return d.size
}

// Get DNA gene in string.
func (d *DNA) Gene() []float64 {
	return d.gene
}

// Get DNA's fitness value in float64.
func (d *DNA) Fitness() float64 {
	return d.fitness
}

// Reset fitness score with default score.
func (d *DNA) Reset() {
	d.fitness = 0.0
}

// Evaluate given a score.
func (d *DNA) Evaluate(score float64) {
	d.fitness = score
}

// Copy other DNA's information.
func (d *DNA) Copy(d1 *DNA) {
	d.size = d1.size
	copy(d.gene, d1.gene)
	d.fitness = d1.fitness
}

// Gaussian convolution mutation given mutation rate
func (d *DNA) Mutate(r float64) {
	for i, _ := range d.gene {
		if rand.Float64() < r {
			// gaussian convolution
			conv := rand.NormFloat64()
			d.gene[i] += conv
		}
	}
}
