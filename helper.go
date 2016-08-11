package eago

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
