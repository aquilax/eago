package eago

import "testing"

func Compare(d1, d2 *DNA) int {
	// higher the fitness, the better
	switch {
	case d1.Fitness() > d2.Fitness():
		return 1
	case d1.Fitness() < d2.Fitness():
		return -1
	default:
		return 0
	}
}

func Eval(d *DNA) float64 {
	gene := d.Gene()
	score := 0.0
	for i := 0; i < 20; i++ {
		bit := string(gene[i])
		if bit == "1" {
			score++
		}
	}
	return score
}

func TestGA(t *testing.T) {
	// new config
	c := &Config{
		DNALen:       20,
		PopSize:      50,
		NumGen:       100,
		MutationRate: 0.3,
		Compare:      Compare,
		Evaluate:     Eval,
		Select:       TSelect,
	}

	ga := NewGA(c)
	ga.Run()
}
