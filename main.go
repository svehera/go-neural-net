package main

import (
	"fmt"
)

func main() {

	net := newNetwork([]int{3, 5, 2, 1})
	//neoto := newNeuoron(3)

	//fmt.Printf("%v\n", neoto)
	fmt.Printf("%v\n", net.Inputs)
	fmt.Printf("%v\n", net.Hiddens)
	fmt.Printf("%v\n", net.Output)
	// Випадкове пcочаткове значення
	//rand.Seed(0)
	/*	fmt.Println(NandPerceptron([]float64{0, 0}))
		fmt.Println(NandPerceptron([]float64{0, 1}))
		fmt.Println(NandPerceptron([]float64{1, 0}))
		fmt.Println(NandPerceptron([]float64{1, 1}))
	*/
	/*fmt.Println()

	fmt.Println(AndPerceptron(0, 0))
	fmt.Println(AndPerceptron(1, 0))
	fmt.Println(AndPerceptron(0, 1))
	fmt.Println(AndPerceptron(1, 1))

	fmt.Println()

	fmt.Println(ORPerceptron(0, 0))
	fmt.Println(ORPerceptron(1, 0))
	fmt.Println(ORPerceptron(0, 1))
	fmt.Println(ORPerceptron(1, 1))

	fmt.Println()

	fmt.Println(NotPerceptron(0))
	fmt.Println(NotPerceptron(1))*/
	// Cтворення XOR патерну для тренування мережі
	/*	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}*/
}

func AndPerceptron(in1, in2 int) int {
	bias := -1
	weight1 := 1
	weight2 := 1

	if weighted := weight1*in1 + weight2*in2 + bias; weighted <= 0 {
		return 0
	}

	return 1
}

func ORPerceptron(in1, in2 int) int {
	bias := 0
	weight1 := 1
	weight2 := 1

	if weighted := weight1*in1 + weight2*in2 + bias; weighted <= 0 {
		return 0
	}

	return 1
}

func NotPerceptron(in1 int) int {
	weight1 := -1
	bias := 1
	if weighted := weight1*in1 + bias; weighted <= 0 {
		return 0
	}

	return 1
}

func NandPerceptron(inputs []float64) float64 {
	NandNeuron := Neuoron{
		Weights: []float64{-16, -16},
		Bias:    24,
	}
	return NandNeuron.CalculateOutput(inputs, SigmoidActivation)
}
