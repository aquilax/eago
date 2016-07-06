package eago

type EA interface {
	InitPopulation()
	FitnessAssess()
	Run()
	Update()
}
