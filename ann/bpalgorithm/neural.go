package main

type NeuralFactor struct {
	weight    float32
	lastDelta float32
	delta     float32
}

func (n *NeuralFactor) ApplyWeightChange(learningRate float32) {
	n.lastDelta = n.delta
	n.weight += n.delta * learningRate
}

func (n *NeuralFactor) ResetWeightChange() {
	n.lastDelta = 0
	n.delta = 0
}

type Neuron struct {
	output  float32
	err     float32
	lastErr float32
	bias    *NeuralFactor
}
