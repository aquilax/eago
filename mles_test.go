package eago

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestMLESComma(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// test (mu, lambda) ES
	mlesComma, err := NewMLESComma(5, 20, &Config{
		DrawGraph:    false,
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
	fmt.Printf("(Mu, Lambda) Evolution Strategy\n")
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))

	// test (mu + lambda) ES
	mlesPlus, err := NewMLESPlus(5, 20, &Config{
		DrawGraph:    false,
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
	mlesPlus.Run()
	// get best dna
	best = mlesPlus.Best()
	fmt.Printf("(Mu + Lambda) Evolution Strategy\n")
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))
}
