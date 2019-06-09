package main

import (
	"fmt"
	"math/rand"
)

type Layer []Neuoron

type Neuoron struct {
	Weights []float64
	Bias    float64
}

func (neuoron Neuoron) String() string {
	return fmt.Sprintf("Weights: %v\nBias: %f\n", neuoron.Weights, neuoron.Bias)
}

/*func (layer Layer) String() string {
	return fmt.Sprintf("New Layer: \n")
}*/

type Network struct {
	Inputs  []float64
	Hiddens []Layer
	Output  Layer
}

func newNeuoron(weightCount int) Neuoron {
	weights := make([]float64, weightCount)
	for i := range weights {
		weights[i] = rand.Float64()
	}
	return Neuoron{Weights: weights, Bias: rand.Float64()}
}

func newNetwork(input []float64, layers []int) Network {
	layersLen := len(layers)
	output := make([]Neuoron, layers[layersLen-1])
	hiddens := make([]Layer, layersLen-1)

	for i := range hiddens {

	}
	for i := 1; i < layersLen-1; i++ {
		neuorons := make([]Neuoron, layers[i])
		for j := 0; j < layers[i]; j++ {
			neuorons[j] = newNeuoron(layers[i-1])
		}
		hiddens[i-1] = neuorons
	}
	for i := range output {
		output[i] = newNeuoron(layers[layersLen-2])
	}

	return Network{Inputs: input, Hiddens: hiddens, Output: output}
}

func (neuron *Neuoron) CalculateWeights(inputs []float64) float64 {
	sum := 0.0
	for i := 0; i < len(neuron.Weights); i++ {
		sum += neuron.Weights[i] * inputs[i]
	}
	return sum + neuron.Bias
}

func (neuron *Neuoron) CalculateOutput(inputs []float64, activation ActivationFunc) float64 {
	return activation(neuron.CalculateWeights(inputs))
}
