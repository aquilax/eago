package eago

import (
	"fmt"
	"log"
	"testing"
)

func TestMLESComma(t *testing.T) {
	mlesComma, err := NewMLESComma(5, 20, &Config{
		GeneLen:      30,
		PopSize:      100,
		NumGen:       100,
		MutationRate: 0.3,
		Compare:      InverseCompare(),
		Evaluate:     Eval,
		Crossover:    UCrossover(0.3),
	})
	if err != nil {
		log.Fatal(err)
	}
	mlesComma.Run()
	// get best dna
	best := mlesComma.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))
}
