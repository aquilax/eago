package ne

import (
	"fmt"
	"log"
	"testing"

	"github.com/jinseokYeom/eago"
	"github.com/jinseokYeom/neugo"
)

func TestNE(t *testing.T) {
	// NE evaluation function
	xorTest, err := NE(&neugo.Config{
		NumInput:     2,               // number of inputs
		NumOutput:    1,               // number of outputs
		NumHidden:    3,               // number of neurons in hidden layer
		NumLayer:     1,               // number of hidden layers
		Bias:         -1.0,            // bias for activation
		WeightMean:   0.0,             // mean for weight
		WeightStdDev: 1.0,             // standard deviation for weight
		Activation:   neugo.Sigmoid(), // activation function
	}, neugo.XORTest())
	if err != nil {
		log.Fatal(err)
	}
	// neuroevolution
	ga := eago.NewGA(&eago.Config{
		DrawGraph:    false,                 // true if draw graph
		GeneLen:      1000,                  // size of each DNA
		PopSize:      100,                   // population size (number of DNAs)
		NumGen:       200,                   // number of generations
		MutationRate: 0.3,                   // mutation rate
		Compare:      eago.InverseCompare(), // comparison method
		Evaluate:     xorTest,               // evaluation method
		Select:       eago.TSelect(),        // selection method
		Crossover:    eago.UCrossover(0.3),  // crossover method
	})
	// run GA
	ga.Run()
	// check best
	best := ga.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %f\n", best.Fitness())
}
