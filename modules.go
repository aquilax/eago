package eago

// Initialize population and return population for a custom,
// modular implementation of evolutionary algorithm.
func InitPopulation(conf *Config) []*DNA {
	population := make([]*DNA, conf.PopSize)
	for i, _ := range population {
		population[i] = NewDNA(conf.GeneLen)
	}
	return population
}

// Assess fitness of each DNA in a population and return an updated
// population for a custom, modular implementation of evolutionary
// algorithm.
func AssessFitness(conf *Config, population []*DNA) []*DNA {
	for i, _ := range population {
		population[i].Reset()
		population[i].Evaluate(conf.Evaluate)
	}
	// return sorted population
	return quickSort(conf, population)
}

// quick sort a population based on their fitness
func quickSort(conf *Config, population []*DNA) []*DNA {
	popSize := len(population)
	if popSize <= 1 {
		return population
	}
	pivot := population[popSize/2]
	low := []*DNA{}
	high := []*DNA{}
	equal := []*DNA{}
	for _, dna := range population {
		switch conf.Compare(dna, pivot) {
		case -1:
			low = append(low, dna)
		case 1:
			high = append(high, dna)
		default:
			equal = append(equal, dna)
		}
	}
	low = append(equal, quickSort(conf, low)...)
	high = quickSort(conf, high)
	return append(high, low...)
}
