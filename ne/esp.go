package ne

import "github.com/jinseokYeom/eago"

type Subpopulation []*DNA

// protected constructor for subpopulation
func newSubpopulation(num, len int) []*DNA {
	sp := make([]*DNA, num)
	for i := 0; i < num; i++ {
		sp[i] = NewDNA(len)
	}
	return sp
}

type ESP struct {
	nnConf *Config      // neural network configuration
	eaConf *eago.Config // evolutionary algorithms config
	//coop   *CoOpGA      // Cooperative coevolution GA
	nnet *NeuralNet // neural network
}

func NewESP(nn *Config, ea *eago.Config) *ESP {
	return &ESP{
		nnConf: nn,
		eaConf: ea,
		nnet:   NewNeuralNet(nn),
		/*
			        population: func() []*Subpopulation {
						// number of neuron + 1 (bias)
						numSubpopulation := conf.NumNeurons + 1
						pop := make([]*Subpopulation, numSubpopulation)
						for i := 0; i < numSubpopulation; i++ {
							pop[i] = newSubpopulation()
						}
						return pop
					}(),
		*/
	}
}
