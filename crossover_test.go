package eago

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCrossover(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Printf("Seed: %d\n", seed)

	d1 := NewDNA(20)
	d2 := NewDNA(20)

	fmt.Printf("DNA 1: %s\n", d1.Gene())
	fmt.Printf("DNA 2: %s\n", d2.Gene())

	// One point crossover
	fmt.Printf("One point crossover\n")
	ch1, ch2 := Crossover1P(d1, d2)
	fmt.Printf("DNA 1: %s\n", ch1.Gene())
	fmt.Printf("DNA 2: %s\n", ch2.Gene())

	// Two point crossover
	fmt.Printf("Two point crossover\n")
	ch1, ch2 = Crossover2P(d1, d2)
	fmt.Printf("DNA 1: %s\n", ch1.Gene())
	fmt.Printf("DNA 2: %s\n", ch2.Gene())

	// Uniform crossover
	fmt.Printf("Uniform crossover\n")
	ch1, ch2 = UCrossover(0.2)(d1, d2)
	fmt.Printf("DNA 1: %s\n", ch1.Gene())
	fmt.Printf("DNA 2: %s\n", ch2.Gene())
}
