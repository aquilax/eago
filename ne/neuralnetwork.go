package ne

import (
	"log"

	"github.com/jinseokYeom/eago"
)

const (
	BIAS = -1.0 // bias for activation
)

type NeuralNet struct {
	conf    *Config   // neural net configuration
	weights []float64 // weights for activation
}

// Create a new neural network with 0 weights.
func NewNeuralNet(conf *Config) *NeuralNet {
	return &NeuralNet{
		conf: conf,
		weights: func(c *Config) []float64 {
			// number of weights including bias
			numWeights := (c.NumInputs+1)*c.NumNeurons +
				(c.NumLayers-1)*(c.NumNeurons+1)*c.NumNeurons +
				(c.NumNeurons+1)*c.NumOutputs
			weights := make([]float64, numWeights)
			return weights
		}(conf),
	}
}

// Get the neural network's weights.
func (n *NeuralNet) Weights() []float64 {
	return n.weights
}

// Get the number of neural network's weights.
func (n *NeuralNet) NumWeights() int {
	return len(n.weights)
}

// Build weights by decoding a DNA and
// generate neural network's weights.
func (n *NeuralNet) Build(d *eago.DNAFloat64) {
	if d.Size() != len(n.weights) {
		log.Fatal("Invalid DNA size")
	}
	n.weights = n.decode(d)
}
