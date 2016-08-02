# EAGO [![GoDoc](https://godoc.org/github.com/jinseokYeom/eago?status.svg)](https://godoc.org/github.com/jinseokYeom/eago)

![alt text](https://github.com/jinseokYeom/eago/blob/master/eago.png "EAGO")

## Overview
EAGO is a simple framework for evolutionary computation (EC) in Go. You can use
this framework to run any kind of EC experiment, simply by two steps: configure
and run. 

## Algorithms
Following are algorithms that are available in this framework. More will be added
with future updates.
* (mu, lambda) Evolution Strategy (in progress)
* (mu + lambda) Evolution Strategy (in progress)
* Genetic Algorithm
* Genetic Algorithm with Elitism (planned)
* Steady-State Genetic Algorithm (planned)
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
	// Configure first
	c := &eago.Config{
		GeneLen:      20,                   // size of each DNA
		PopSize:      50,                   // population size (number of DNAs)
		NumGen:       100,                  // number of generations
		MutationRate: 0.3,                  // mutation rate
		Compare:      eago.DirectCompare(), // comparison method
		Evaluate:     Eval,                 // evaluation method
		Select:       eago.TSelect(),       // selection method
		Crossover:    eago.UCrossover(0.3), // crossover method
	}

    // Create GA,
	ga := eago.NewGA(c)

    // ...and run!
    ga.Run()

    // Check the best DNA and its fitness.
	best := ga.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))
}
```

# EAGO/NE [![GoDoc](https://godoc.org/github.com/jinseokYeom/eago/ne?status.svg)](https://godoc.org/github.com/jinseokYeom/eago/ne)
EAGO/NE is an extention of EAGO for NeuroEvolution.
NeuroEvolution is a machine learning method that uses Evolution Algorithm
to train a neural network. ```github.com/jinseokYeom/eago/ne``` package is a modular
implementation that can be customized by the user's choice. 

## Algorithms (in progress)
* NE (Basic NeuroEvolution with GA)
* SANE (Symbiotic, Adaptive NeuroEvolution)
* ESP (Enforced Sub-Population)
* NEAT (NeuroEvolution of Augmenting Topologies)

## Note
Simplicity and ease of use are the top priorities of this framework; 
some decisions were made during the development for such reason.
* Genotype is only represented with binary string (at least for now).
* Fitness values are always float64.

## I NEED your support!
If you're reading this, you probably have at least some knowledge in Evolutionary
Computation and its application on optimization problems. I, myself, am still
a student who is also studying the subject matter as I work on this project, 
and there is no doubt that Evolutionary Computation is a very broad subject; 
hence, there is a lot to cover in this framework. This isn't something that can be
done by one person. More developers will be always appreciated!
