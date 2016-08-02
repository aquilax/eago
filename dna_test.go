package eago

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDNA(t *testing.T) {
	// rand seed set to 0
	rand.Seed(0)

	// create a new binary-coded DNA
	d := NewDNA(10)
	fmt.Printf("DNA size: %d\n", d.Size())
	fmt.Printf("DNA gene: %s\n", d.Gene())

	// rand seed set to 1
	rand.Seed(1)

	// create a new DNA
	d = NewDNA(10)
	fmt.Printf("DNA size: %d\n", d.Size())
	fmt.Printf("DNA gene: %s\n", d.Gene())

	// bit flip mutation
	mrate := 0.3
	d.Mutate(mrate)
	fmt.Printf("Mutation rate: %f\n", mrate)
	fmt.Printf("Mutated DNA gene: %s\n", d.gene)
}
