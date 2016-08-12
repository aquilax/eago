package eago

import (
	"bytes"
	"math/rand"
)

// DNA is a binary coded genotype. It's gene is encoded
// with a string of "0"s and "1"s.
type DNA struct {
	size      int     // size of the chromosome
	gene      string  // binary-coded gene
	fitness   float64 // fitness score
	evaluated bool    // true if it has been evaluated
}

func NewDNA(size int) *DNA {
	return &DNA{
		size: size,
		gene: func() string {
			var buf bytes.Buffer
			for i := 0; i < size; i++ {
				if rand.Float64() < 0.5 {
					buf.WriteString("1")
				} else {
					buf.WriteString("0")
				}
			}
			return buf.String()
		}(),
		fitness:   0.0,
		evaluated: false,
	}
}

// Get DNA size in int.
func (d *DNA) Size() int {
	return d.size
}

// Get DNA gene in string.
func (d *DNA) Gene() string {
	return d.gene
}

// Get DNA's fitness value in float64.
func (d *DNA) Fitness() float64 {
	return d.fitness
}

// Check if this DNA has been evaluated.
func (d *DNA) IsEvaluated() bool {
	return d.evaluated
}

// Reset fitness score with default score.
func (d *DNA) Reset() {
	d.fitness = 0.0
	d.evaluated = false
}

// Evaluate given an evaluation function.
func (d *DNA) Evaluate(eval EvaluateFunc) {
	d.fitness = eval(d)
}

// Copy other DNA's information.
func (d *DNA) Copy(d1 *DNA) {
	d.size = d1.size
	d.gene = d1.gene
	d.fitness = d1.fitness
	d.evaluated = d1.evaluated
}

// Bit flip mutation given mutation rate
func (d *DNA) Mutate(r float64) {
	var newGene bytes.Buffer
	for i := 0; i < d.size; i++ {
		bit := string(d.gene[i])
		if rand.Float64() > r {
			newGene.WriteString(bit)
		} else {
			if bit == "0" {
				newGene.WriteString("1")
			} else {
				newGene.WriteString("0")
			}
		}
	}
	// update gene
	d.gene = newGene.String()
}
