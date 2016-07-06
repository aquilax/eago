package eago

import (
	"bytes"
	"math/rand"
)

type DNA struct {
	size int    // size of the chromosome
	gene string // binary-coded gene
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
