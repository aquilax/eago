package esp

import "github.com/jinseokYeom/eago"

type Subpopulation []*eago.DNA

// protected constructor for subpopulation
func newSubpopulation(num, len int) *Subpopulation {
	sp := make([]*eago.DNA, num)
	for i := 0; i < num; i++ {
		sp[i] = eago.NewDNA(len)
	}
	return sp
}

type ESP struct {
	conf       *Config          // configuration
	nnet       *NeuralNet       // neural network
	population []*Subpopulation // population of subpopulation
}

func NewESP(conf *Config) *ESP {
	return &ESP{
		conf: conf,
		nnet: NewNeuralNet(),
		population: func() []*Subpopulation {
			// number of neuron + 1 (bias)
			numSubpopulation := conf.NumNeuron + 1
			pop := make([]*Subpopulation, numSubpopulation)
			for i := 0; i < numSubpopulation; i++ {
				pop[i] = newSubpopulation()
			}
			return pop
		}(),
	}
}
