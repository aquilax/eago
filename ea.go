package eago

type EA interface {
	// initialize population
	InitPopulation()
	// assess fitness
	AssessFitness()
	// update best/population
	Update()
	// execute algorithm
	Run()
	// get best performing gene
	Best() string
}
