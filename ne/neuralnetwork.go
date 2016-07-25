package ne

type NeuralNet struct {
	numInput  int // number of inputs
	numOutput int // number of outputs

}

func NewNeuralNet() *NeuralNet {
	return &NeuralNet{}
}
