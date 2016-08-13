package ne

import "github.com/jinseokYeom/neugo"

// NeuroEvolution is defined as an evaluation function
// that decode a DNA to a neural network (from neugo package)
// phenotype and evaluate the fitness.
func NE(e *neugo.Environment) EvaluateFunc {
	return func(*DNA) float64 {
		// decode DNA

		// creat neural network

		// evaluate neural network

		return 0.0
	}
}
