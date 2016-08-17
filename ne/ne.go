package ne

import (
	"math"

	"github.com/jinseokYeom/eago"
	"github.com/jinseokYeom/neugo"
)

// NeuroEvolution is defined as an evaluation function
// that decode a DNA to a neural network (from neugo package)
// phenotype and evaluate the fitness, given a configuration and
// an evaluation environment.
func NE(nn *neugo.NeuralNet, env neugo.Environment) eago.EvaluateFunc {
	// return the evaluation function
	return func(d *eago.DNA) float64 {
		// build the neural net with decoded DNA
		nn.Build(decode(d, nn.NumWeights()))
		return env(nn)
	}
}

// decode a DNA to get list of weights
func decode(d *eago.DNA, numWeights int) []float64 {
	prev := 0
	weights := make([]float64, numWeights)
	len := d.Size() / numWeights
	for i, _ := range weights {
		next := prev + len
		slice := d.Gene()[prev:next]
		weights[i] = func() float64 {
			weight := 0.0
			for i := 0; i < len; i++ {
				if string(slice[i]) == "1" {
					pow := float64(len - 1 - i)
					weight += math.Pow(2.0, pow)
				}
			}
			return weight
		}()
		prev = next
	}
	return weights
}

// SANE

// ESP

// NEAT
