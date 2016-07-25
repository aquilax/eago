package ne

import "testing"

func TestNeuralNet(t *testing.T) {
	numInput := 3  // 3 inputs
	numOutput := 4 // 4 outputs

	nn := NewNeuralNet()
	inputs := make([]float64, numInput)
	nn.Update(inputs)
}
