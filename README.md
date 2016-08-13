# EAGO [![GoDoc](https://godoc.org/github.com/jinseokYeom/eago?status.svg)](https://godoc.org/github.com/jinseokYeom/eago)

![alt text](https://github.com/jinseokYeom/eago/blob/master/eago.png "EAGO")

## Overview
EAGO is a simple framework for evolutionary computation (EC) in Go. You can use
this framework to run any kind of EC experiment, simply by two steps: configure
and run. 

## Algorithms
Following are algorithms that are available in this framework. More will be added
with future updates.
* (mu, lambda) Evolution Strategy
* (mu + lambda) Evolution Strategy
* Genetic Algorithm

## Planned Algorithms
These are algorithms that are not implemented yet, but are planned to be added
in the future.
* Competitive/Cooperative Coevolution algorithms
* CMA-ES (Covariance Matrix Adaptation Evolution Strategy)
* NSGA-II (Non-dominated Sorting Genetic Algorithm 2)
* And more

## Installation
`go get github.com/jinseokYeom/eago`

## Example

```go
package main

import (
	"fmt"
    "github.com/jinseokYeom/eago"
)

// define an evaluation function
func Eval(d *eago.DNA) float64 {
	gene := d.Gene()
	score := 0.0
	for i := 0; i < 20; i++ {
		bit := string(gene[i])
		if bit == "1" {
			score++
		}
	}
	return score
}

func GeneticAlgorithm() {
    // Configure and create GA,
	ga := eago.NewGA(&eago.Config{
       	DrawGraph:    false,                // true if you want graph
        GeneLen:      20,                   // size of each DNA
		PopSize:      50,                   // population size (number of DNAs)
		NumGen:       100,                  // number of generations
		MutationRate: 0.3,                  // mutation rate
		Compare:      eago.DirectCompare(), // comparison method
		Evaluate:     Eval,                 // evaluation method
		Select:       eago.TSelect(),       // selection method
		Crossover:    eago.UCrossover(0.3), // crossover method
    })

    // ...and run!
    ga.Run()

    // Check the best DNA and its fitness.
	best := ga.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))

    // -------- RESULT --------
    // BEST: 111101111111001111011110111111
    // SCORE: 25
}
```

# EAGO/NE [![GoDoc](https://godoc.org/github.com/jinseokYeom/eago/ne?status.svg)](https://godoc.org/github.com/jinseokYeom/eago/ne)
EAGO/NE is an extension of EAGO for NeuroEvolution.
NeuroEvolution is a machine learning method that uses Evolution Algorithm
to train a neural network. For the sake of this extension, I have decided to
work on a different package, NEUGO (Neural Networks in Go), which will be an 
easy-to-use neural network package in Go.

## Planned Algorithms
* NE (Basic NeuroEvolution with GA)
* SANE (Symbiotic, Adaptive NeuroEvolution)
* ESP (Enforced Sub-Population)
* NEAT (NeuroEvolution of Augmenting Topologies)
* And more

# EAGO/GP [![GoDoc](https://godoc.org/github.com/jinseokYeom/eago/gp?status.svg)](https://godoc.org/github.com/jinseokYeom/eago/gp)
EAGO/GP is an extension of EAGO for Genetic Programming.

## Notes
1. Although I'm hoping to make this a helpful package for many people,
I started developing this package for the sake of learning about different
kinds of evolutionary algorithms in depth; I would appreciate any criticism
on this package! 
2. Simplicity and ease of use are the top priorities of this framework; 
some decisions were made during the development for such reason.
  * Genotype is only represented with binary string (at least for now).
  * Fitness values are always float64.
