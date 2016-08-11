package eago

type EA interface {
	// assess fitness
	AssessFitness()
	// update best/population
	Update()
	// execute algorithm
	Run()
	// get best performing gene
	Best() string
}
