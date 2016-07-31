package ne

type Config struct {
	NumInputs  int            // number of inputs
	NumOutputs int            // number of outputs
	NumLayers  int            // number of hidden layers
	NumNeurons int            // number of hidden neurons in a layer
	Activation ActivationFunc // activation function
}

// Create a new configuration set to null.
func NewConfig() *Config {
	return &Config{}
}

// Set number of inputs.
func (c *Config) SetNumInputs(n int) {
	c.NumInputs = n
}

// Set number of outputs.
func (c *Config) SetNumOutputs(n int) {
	c.NumOutputs = n
}
