package eago

// Compare function is a type of function that compares
// the first DNA to the second DNA, and returns -1 if
// the first DNA is less fit, 0 if it is as fit as the second,
// and 1 if it is more fit than the second.
type CompareFunc func(*DNA, *DNA) int

// Return a comparison function in which score and fitness
// are directly related.
func DirectCompare() CompareFunc {
	return func(d1, d2 *DNA) int {
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
}

// Return a comparison function in which score and fitness
// are inversely related.
func InverseCompare() CompareFunc {
	return func(d1, d2 *DNA) int {
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
}
