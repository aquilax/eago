package eago

// score and fitness are directly related
// if g1.fitness < g2.fitness, return -1
// if g1.fitness = g2.fitness, return 0
// if g1.fitness > g2.fitness, return 1
func DirectCompare(d1, d2 *DNA) int {
	f1 := d1.Fitness()
	f2 := d2.Fitness()
	switch {
	case f1 < f2:
		return -1
	case f1 == f2:
		return 0
	default:
		return 1
	}
}

// score and fitness are inversely related
// if g1.fitness > g2.fitness, return -1
// if g1.fitness = g2.fitness, return 0
// if g1.fitness < g2.fitness, return 1
func InverseCompare(d1, d2 *DNA) int {
	f1 := d1.Fitness()
	f2 := d2.Fitness()
	switch {
	case f1 > f2:
		return -1
	case f1 == f2:
		return 0
	default:
		return 1
	}
}
