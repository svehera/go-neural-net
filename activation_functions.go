package main

import "math"

type ActivationFunc func(float64) float64

func SigmoidActivation(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}
