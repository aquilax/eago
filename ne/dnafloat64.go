package ne

import "math/rand"

// DNAFloat64 is a real-coded genotype with a slice of
// float64 as its gene.
type DNAFloat64 struct {
	size    int       // size of the chromosome
	gene    []float64 // real-coded gene
	fitness float64   // fitness score
}

// Create a new real-coded DNA given its size,
// mean value, and standard deviation for normally
// distributed random numbers.
func NewDNAFloat64(size int, m, sd float64) *DNAFloat64 {
	return &DNAFloat64{
		size: size,
		gene: func() []float64 {
			g := make([]float64, size)
			for i := 0; i < size; i++ {
				g[i] = rand.NormFloat64()*sd + m
			}
			return g
		}(),
		fitness: 0.0,
	}
}

// Get DNA size in int.
func (d *DNAFloat64) Size() int {
	return d.size
}

// Get DNA gene in string.
func (d *DNAFloat64) Gene() []float64 {
	return d.gene
}

// Get DNA's fitness value in float64.
func (d *DNAFloat64) Fitness() float64 {
	return d.fitness
}

// Reset fitness score with default score.
func (d *DNAFloat64) Reset() {
	d.fitness = 0.0
}

// Evaluate given a score.
func (d *DNAFloat64) Evaluate(eval EvaluateFunc) {
	d.fitness = score
}

// Copy other DNA's information.
func (d *DNAFloat64) Copy(d1 *DNAFloat64) {
	d.size = d1.size
	copy(d.gene, d1.gene)
	d.fitness = d1.fitness
}

// Gaussian convolution mutation given mutation rate
func (d *DNAFloat64) Mutate(r float64) {
	for i, _ := range d.gene {
		if rand.Float64() < r {
			// gaussian convolution
			conv := rand.NormFloat64()
			d.gene[i] += conv
		}
	}
}
