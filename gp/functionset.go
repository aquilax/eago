package gp

// Function set that will be used in GP.
type FunctionSet struct {
	fns []func(v ...interface{})
}

func NewFunctionSet() *FunctionSet {
	return &FunctionSet{
		fns: make([]func(v ...interface{})),
	}
}
