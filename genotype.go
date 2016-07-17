package eago

import (
	"bytes"
	"math/rand"
)

// genotype interface
type Genotype interface {
	// get size
	Size() int
	// get gene
	Gene() string
	// get fitness
	Fitness() float64
	// copy information
	Copy(Genotype)
	// mutation
	Mutate(float64)
	// evaluate with evaluation function
	Evaluate(EvaluateFunc)
}

// binary-coded DNA
type DNA struct {
	size    int     // size of the chromosome
	gene    string  // binary-coded gene
	fitness float64 // fitness score
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
	}
}

// get size
func (d *DNA) Size() int {
	return d.size
}

// get gene
func (d *DNA) Gene() string {
	return d.gene
}

// get fitness
func (d *DNA) Fitness() float64 {
	return d.fitness
}

// reset fitness score with default score
func (d *DNA) Reset() {
	d.fitness = 0.0
}

// evaluate given an evaluation function
func (d *DNA) Evaluate(eval EvaluateFunc) {
	d.fitness = eval(d)
}

// copy other DNA's information
func (d *DNA) Copy(d1 *DNA) {
	d.size = d1.size
	d.gene = d1.gene
	d.fitness = d1.fitness
}

// bit flip mutation given mutation rate
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
