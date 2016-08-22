package eago

// Evaluation function is a type of function that evaluates
// one or more DNAs' fitnesses and return their fitness value
// in float64. This step involves decoding DNAs (genotype)
// into a phenotype, then testing its performance.
type EvaluateFunc func(...*DNA) float64
