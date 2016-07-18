package eago

// Evaluation function is a type of function that evaluates
// a DNA's fitness and return its fitness value in float64.
// This step involves decoding the DNA (genotype) into a
// phenotype, then testing its performance.
type EvaluateFunc func(*DNA) float64
