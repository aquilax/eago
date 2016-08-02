package eago

import (
	"fmt"
	"testing"
)

func Eval(d *DNA) float64 {
	gene := d.Gene()
	size := d.Size()
	score := 0.0
	for i := 0; i < size; i++ {
		bit := string(gene[i])
		if bit == "1" {
			score += 1.0
		}
	}
	return score
}

func TestGA(t *testing.T) {
	ga := NewGA(&Config{
		GeneLen:      30,
		PopSize:      100,
		NumGen:       100,
		MutationRate: 0.3,
		Compare:      DirectCompare(),
		Evaluate:     Eval,
		Select:       TSelect(),
		Crossover:    UCrossover(0.3),
	})
	ga.Run()

	best := ga.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))
}
