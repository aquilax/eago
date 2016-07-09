package algorithm

// Gradient Ascent
// 1st arg: allowed number of iterations
// 2nd arg: random start vector
// 3rd arg: function to be optimized
// 4th arg: alpha (step size)
func GradientAscent(time int, vec []float64,
	fn func(vec []float64) float64, a float64) []float64 {
	for i := 0; i < time; i++ {
		for j, x := range vec {
			gradient := (fn(x+0.0001) - fn(x)) / 0.0001
			vec[j] += gradient * a
		}
	}
	return vec
}

// Gradient Descent
// 1st arg: allowed number of iterations
// 2nd arg: random start vector
// 3rd arg: function to be optimized
// 4th arg: alpha (step size)
func GradientDescent(time int, vec []float64,
	fn func(vec []float64) float64, a float64) []float64 {
	for i := 0; i < time; i++ {
		for j, x := range vec {
			gradient := (fn(x+0.0001) - fn(x)) / 0.0001
			vec[j] -= gradient * a
		}
	}
	return vec
}
